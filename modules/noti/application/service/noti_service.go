package service

/*응용서비스 영역
- 리포지터리에서 애그리거트 조회/저장/생성/결과리턴/도메인기능실행
- 트랜잭션 처리
- 이벤트처리
- 접근제어
- 복잡하다면 도메인 로직이 포함되지않았나 의심
- 표현영역과 도메인 영역을 연결하는 매개체 열할
- 한 응용 서비스 클래스에서 한개내지 2~3개의 기능 구현 중복되는 로직은 별도 클래스로 작성해서 포함
*/

import (
	"errors"

	common "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/common"
	appDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/application/dto"

	//domRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/resource/domain/repository"

	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
)

// type IAuthService interface {
// 	GetUserByNameInternal(userName string) (*common.InternalGetUserByNameResponse, error)
// }

type IWebHookService interface {
	SendWebHook(req *appDTO.SendWebHookRequestDTO, i identity.Identity) error
}

type NotiService struct {
	BaseService
	//repo           domRepo.IPredictionEnvRepo
	EmailSvc      common.IEmailService
	DeploymentSvc common.IDeploymentService
	ProjectSvc    common.IProjectService
	AuthSvc       common.IAuthService
	WebHookSvc    IWebHookService
}

type TemplateData struct {
	deploymentName string
	//deploymentURL  string
	additionalData string
}

// NewNotiService new NotiService
func NewNotiService(h *handler.Handler, emailSvc common.IEmailService, deploymentSvc common.IDeploymentService, projectSvc common.IProjectService, authSvc common.IAuthService, webHookService IWebHookService) (*NotiService, error) {
	// var err error

	svc := new(NotiService)
	svc.handler = h
	if err := svc.initBaseService(); err != nil {
		return nil, err
	}

	// if svc.repo, err = infRepo.NewPredictionEnvRepo(h); err != nil {
	// 	return nil, err
	// }

	svc.EmailSvc = emailSvc
	svc.DeploymentSvc = deploymentSvc
	svc.ProjectSvc = projectSvc
	svc.AuthSvc = authSvc
	svc.WebHookSvc = webHookService

	return svc, nil
}

// SendNoti
func (s *NotiService) SendNoti(deploymentID string, notiCategory string, additionalData string) error {
	//authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	// if err := req.Validate(); err != nil {
	// 	return nil, err
	// }

	///////////////////////////////////////////////////////
	// to be need injection identity from monitoring module
	///////////////////////////////////////////////////////
	i := s.systemIdentity

	///////////////////////
	//Get UserIno Sequence
	///////////////////////
	var templateData TemplateData

	//Get Deployment Info
	resDeploy, err := s.DeploymentSvc.GetByIDInternal(deploymentID)
	if err != nil {
		return err
	}

	resPrj, err := s.ProjectSvc.GetByIDInternal(resDeploy.ProjectID)
	if err != nil {
		return err
	}

	//Get UserEmail
	resAuth, err := s.AuthSvc.GetUserByNameInternal(resPrj.CreatedBy)
	if err != nil {
		return err
	}

	emailAddress := resAuth.Email

	nickName := resAuth.NickName
	if resAuth.NickName == "" {
		nickName = "user"
	}

	/////////////////////////////////////
	// End of Get UserIno Sequence
	/////////////////////////////////////

	templateData.deploymentName = resDeploy.Name
	templateData.additionalData = additionalData

	//Send Email
	err = s.sendNotiEmail(notiCategory, emailAddress, nickName, templateData, i)
	if err != nil {
		return err
	}

	//Triggering WebHook
	reqWebHookEvent := &appDTO.SendWebHookRequestDTO{
		DeploymentID:  deploymentID,
		TriggerSource: notiCategory,
		TriggerStatus: additionalData,
	}

	err = s.WebHookSvc.SendWebHook(reqWebHookEvent, i)
	if err != nil {
		return err
	}

	// response dto
	// resDTO := new(appDTO.NotiResponseDTO)
	// resDTO.Message = "Success"

	return nil
}

func (s *NotiService) sendNotiEmail(notiCategory string, emailAddress string, nickName string, templateData TemplateData, i identity.Identity) error {
	cfg, err := s.handler.GetConfig()
	if err != nil {
		return err
	}

	var templateCode string
	switch notiCategory {
	case "Datadrift":
		templateCode = "datadrift-alert-html"
	case "Accuracy":
		templateCode = "accuracy-alert-html"
	case "Service":
		templateCode = "service-alert-html"
	default:
		return errors.New("noticategory not found")
	}

	fromEmail := cfg.SMTPServers.DefaultSMTP.SenderEmail
	fromName := cfg.SMTPServers.DefaultSMTP.SenderName

	// url := fmt.Sprintf(cfg.IAM.Registration.ActivationURL, i.RequestInfo.Host)
	// activationURL := fmt.Sprintf("%s/%s/%s", url, resReg.ActivationCode, "html")

	// send activate registration via email (sub)domain [email module]
	reqEmail := new(common.SendEmailReqDTO)
	reqEmail.TemplateCode = templateCode
	reqEmail.From = &common.MailAddressDTO{Email: fromEmail, Name: fromName}
	reqEmail.To = &common.MailAddressDTO{Email: emailAddress, Name: nickName}
	reqEmail.TemplateData = map[string]interface{}{
		"Body.deploymentName": templateData.deploymentName,
		//"Body.deploymentURL":    templateData.deploymentURL,
		"Body.additionalData": templateData.additionalData,
	}
	reqEmail.ProcessingType = "ASYNC"

	if _, err := s.EmailSvc.SendInternal(reqEmail, i); err != nil {
		return err
	}

	return nil
}

func (s *NotiService) Update(event common.Event) {
	switch actualEvent := event.(type) {
	case common.MonitoringAccuracyStatusChangedToFailing:
		//
		s.SendNoti(actualEvent.DeploymentID(), "Accuracy", "Failing")
	case common.MonitoringAccuracyStatusChangedToAtrisk:
		//
		s.SendNoti(actualEvent.DeploymentID(), "Accuracy", "Atrisk")
	case common.MonitoringDataDriftStatusChangedToFailing:
		//
		s.SendNoti(actualEvent.DeploymentID(), "Datadrift", "Failing")
	case common.MonitoringDataDriftStatusChangedToAtrisk:
		//
		s.SendNoti(actualEvent.DeploymentID(), "Datadrift", "Atrisk")
	case common.MonitoringServiceHealthStatusChangedToFailing:
		//
		s.SendNoti(actualEvent.DeploymentID(), "Service", "Failing")
	case common.MonitoringServiceHealthStatusChangedToAtrisk:
		//
		s.SendNoti(actualEvent.DeploymentID(), "Service", "Atrisk")
	default:
		return

	}
}
