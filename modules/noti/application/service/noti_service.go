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

	appDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/application/dto"

	//domRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/resource/domain/repository"
	appAuthDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/auths/application/dto"
	appDeploymentDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/application/dto"
	appEmailDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/application/dto"
	appProjectDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/project/application/dto"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
)

type IEmailService interface {
	Send(req *appEmailDTO.SendEmailReqDTO, i identity.Identity) (*appEmailDTO.SendEmailResDTO, error)
}
type IDeploymentService interface {
	GetByIDInternal(req *appDeploymentDTO.GetDeploymentRequestDTO, i identity.Identity) (*appDeploymentDTO.GetDeploymentResponseDTO, error)
}
type IProjectService interface {
	GetByIDInternal(req *appProjectDTO.GetProjectRequestDTO) (*appProjectDTO.GetProjectResponseDTO, error)
}

type IAuthService interface {
	GetUserByName(req *appAuthDTO.GetUserByNameReqDTO, i identity.Identity) (*appAuthDTO.GetUserByNameResDTO, error)
}

type IGovernanceHistoryService interface {
	AddGovernanceHistory(req *appDeploymentDTO.AddGovernanceHistoryRequestDTO, i identity.Identity) error
}

type NotiService struct {
	BaseService
	//repo           domRepo.IPredictionEnvRepo
	EmailSvc             IEmailService
	DeploymentSvc        IDeploymentService
	ProjectSvc           IProjectService
	AuthSvc              IAuthService
	GovernanceHistorySvc IGovernanceHistoryService
}

type TemplateData struct {
	deploymentName string
	//deploymentURL  string
	additionalData string
}

// NewNotiService new NotiService
func NewNotiService(h *handler.Handler, emailSvc IEmailService, deploymentSvc IDeploymentService, projectSvc IProjectService, authSvc IAuthService, governanceHistorySvc IGovernanceHistoryService) (*NotiService, error) {
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
	svc.GovernanceHistorySvc = governanceHistorySvc
	return svc, nil
}

// SendNoti
func (s *NotiService) SendNoti(req *appDTO.NotiRequestDTO, i identity.Identity) error {
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
	i = s.systemIdentity

	///////////////////////
	//Get UserIno Sequence
	///////////////////////
	var templateData TemplateData

	//Get Deployment Info
	reqDeploy := &appDeploymentDTO.GetDeploymentRequestDTO{
		DeploymentID: req.DeploymentID,
	}
	resDeploy, err := s.DeploymentSvc.GetByIDInternal(reqDeploy, i)
	if err != nil {
		return err
	}

	//Get Project UserInfo
	reqPrj := &appProjectDTO.GetProjectRequestDTO{
		ProjectID: resDeploy.ProjectID,
	}

	resPrj, err := s.ProjectSvc.GetByIDInternal(reqPrj)
	if err != nil {
		return err
	}

	//Get UserEmail
	reqAuth := &appAuthDTO.GetUserByNameReqDTO{
		UserName: resPrj.CreatedBy,
	}
	resAuth, err := s.AuthSvc.GetUserByName(reqAuth, i)
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
	templateData.additionalData = req.AdditionalData

	//Send Email
	err = s.sendNotiEmail(req.NotiCategory, emailAddress, nickName, templateData, i)
	if err != nil {
		return err
	}

	//Log Deployment Log
	err = s.addGovernanceLog(req.NotiCategory, req.DeploymentID, req.AdditionalData, i)
	if err != nil {
		return err
	}

	//Triggering WebHook

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
	// "Datadrift", "Accuracy", "Service"
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
	reqEmail := new(appEmailDTO.SendEmailReqDTO)
	reqEmail.TemplateCode = templateCode
	reqEmail.From = &appEmailDTO.MailAddressDTO{Email: fromEmail, Name: fromName}
	reqEmail.To = &appEmailDTO.MailAddressDTO{Email: emailAddress, Name: nickName}
	reqEmail.TemplateData = map[string]interface{}{
		"Body.deploymentName": templateData.deploymentName,
		//"Body.deploymentURL":    templateData.deploymentURL,
		"Body.additionalData": templateData.additionalData,
	}
	reqEmail.ProcessingType = "ASYNC"
	print("test 8")
	if _, err := s.EmailSvc.Send(reqEmail, i); err != nil {
		return err
	}
	print("test 9")
	return nil
}

func (s *NotiService) addGovernanceLog(notiCategory string, deploymentID string, logMessage string, i identity.Identity) error {
	// "DataDriftAlert", "AccuracyAlert", "ServiceAlert"
	var eventType string
	switch notiCategory {
	case "Datadrift":
		eventType = "DataDriftAlert"
	case "Accuracy":
		eventType = "AccuracyAlert"
	case "Service":
		eventType = "ServiceAlert"
	default:
		return errors.New("eventType not found")
	}

	govReq := &appDeploymentDTO.AddGovernanceHistoryRequestDTO{
		DeploymentID: deploymentID,
		EventType:    eventType,
		LogMessage:   logMessage,
	}

	err := s.GovernanceHistorySvc.AddGovernanceHistory(govReq, i)
	if err != nil {
		return err
	}

	return nil
}
