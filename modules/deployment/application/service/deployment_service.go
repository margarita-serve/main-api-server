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
	appDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/application/dto"
	domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/entity"
	domRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/repository"
	domSchema "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/schema"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"github.com/rs/xid"

	//"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
	domSvcInferenceSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/service/inference_service"
	domSvcInferenceSvcDto "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/service/inference_service/dto"
	infInfSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/infrastructure/inference_service/kserve"
	infRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/infrastructure/repository"
	appModelPackageDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/application/dto"
	appModelPackageSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/application/service"

	// appPredictionEnvDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/predictionEnv/application/dto"
	// appPredictionEnvSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/predictionEnv/application/service"
	// appMonitoringDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/application/dto"
	// appMonitoringSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/application/service"
	// appOrgDTO
	// appOrgSvc
	predictionSendSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/infrastructure/prediction_sender"
)

// DeploymentService type
type DeploymentService struct {
	BaseService
	domInferenceSvc domSvcInferenceSvc.IInferenceServiceAdapter
	modelPackageSvc *appModelPackageSvc.ModelPackageService
	// predictionEnvSvc   *appPredictionEnvSvc.PredictionEnvService
	// monitoringSvc          *appMonitoringSvc.MonitoringService
	predictionSendSvc *predictionSendSvc.PredictionSender
	repo              domRepo.IDeploymentRepo
}

// NewDeploymentService new DeploymentService
func NewDeploymentService(h *handler.Handler) (*DeploymentService, error) {
	var err error

	svc := new(DeploymentService)

	svc.handler = h
	if err := svc.initBaseService(); err != nil {
		return nil, err
	}

	if svc.repo, err = infRepo.NewDeploymentRepo(h); err != nil {
		return nil, err
	}

	if svc.domInferenceSvc, err = infInfSvc.NewInfSvcKserveAdapter(); err != nil {
		return nil, err
	}

	if svc.modelPackageSvc, err = appModelPackageSvc.NewModelPackageService(h); err != nil {
		return nil, err
	}

	if svc.predictionSendSvc, err = predictionSendSvc.NewPredictionSendService(); err != nil {
		return nil, err
	}

	return svc, nil
}

// Create
func (s *DeploymentService) Create(req *appDTO.CreateDeploymentRequestDTO) (*appDTO.CreateDeploymentResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	if err := req.Validate(); err != nil {
		return nil, err
	}

	guid := xid.New().String()

	//toBe...
	userID := "testID"

	// New deployment domain Instance
	domAggregateDeployment, err := domEntity.NewDeployment(
		guid,
		req.ProjectID,
		req.ModelPackageID,
		req.PredictionEnvID,
		req.Name,
		req.Description,
		req.Importance,
		"Normal",
		req.RequestCPU,
		req.RequestMEM,
		0,
		0,
		userID,
	)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(domAggregateDeployment)
	if err != nil {
		return nil, err
	}

	//Find ModelPackage
	resModelPackage, err := s.getModelPackageByID(req.ModelPackageID)
	if err != nil {
		return nil, err
	}

	//Find  PredictionEnv
	resPredictionEnvInfo, err := s.getPredictionEnvByID(req.PredictionEnvID)
	if err != nil {
		return nil, err
	}

	reqDomSvc := domSvcInferenceSvcDto.InferenceServiceCreateRequest{
		Namespace:             resPredictionEnvInfo.Namespace,
		DeploymentID:          guid,
		ModelFrameWork:        resModelPackage.ModelFrameWork,
		ModelFrameWorkVersion: resModelPackage.ModelFrameWorkVersion,
		ModelURL:              resModelPackage.ModelFilePath,
		ModelName:             resModelPackage.ModelName,
		ConnectionInfo:        resPredictionEnvInfo.ConnectionInfo,
		RequestCPU:            domAggregateDeployment.RequestCPU,
		RequestMEM:            domAggregateDeployment.RequestMEM,
		LimitCPU:              domAggregateDeployment.LimitCPU,
		LimitMEM:              domAggregateDeployment.LimitMEM,
	}

	if err := reqDomSvc.Validate(); err != nil {
		return nil, err
	}

	domAggregateDeployment.AddModelHistory(resModelPackage.ModelName, resModelPackage.ModelVersion)

	err = domAggregateDeployment.AddEventHistory("Create", "Deployment Created", userID)
	if err != nil {
		return nil, err
	}

	err = domAggregateDeployment.RequestCreateInferenceService(s.domInferenceSvc, reqDomSvc)
	if err != nil {
		errDelete := s.repo.Delete(domAggregateDeployment.ID)
		if errDelete != nil {
			return nil, errDelete
		}
		return nil, err
	}

	err = s.repo.Save(domAggregateDeployment)
	if err != nil {
		return nil, err
	}

	//Create Monitoring Service
	//createMonitoring
	//

	// response dto
	resDTO := new(appDTO.CreateDeploymentResponseDTO)
	resDTO.DeploymentID = domAggregateDeployment.ID

	return resDTO, nil
}

// ReplaceModel
func (s *DeploymentService) ReplaceModel(req *appDTO.ReplaceModelRequestDTO) (*appDTO.ReplaceModelResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	//toBe...
	userID := "testID"

	if err := req.Validate(); err != nil {
		return nil, err
	}

	//Find Domain Entity
	domAggregateDeployment, err := s.repo.GetByID(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	//Find ModelPackage
	resModelPackage, err := s.getModelPackageByID(req.ModelPackageID)
	if err != nil {
		return nil, err
	}

	//Find  PredictionEnv
	resPredictionEnvInfo, err := s.getPredictionEnvByID(domAggregateDeployment.PredictionEnvID)
	if err != nil {
		return nil, err
	}

	reqDomSvc := domSvcInferenceSvcDto.InferenceServiceReplaceModelRequest{
		Namespace:             resPredictionEnvInfo.Namespace,
		DeploymentID:          domAggregateDeployment.ID,
		ModelFrameWork:        resModelPackage.ModelFrameWork,
		ModelFrameWorkVersion: resModelPackage.ModelFrameWorkVersion,
		ModelURL:              resModelPackage.ModelFilePath,
		ModelName:             resModelPackage.ModelName,
		ConnectionInfo:        resPredictionEnvInfo.ConnectionInfo,
		RequestCPU:            domAggregateDeployment.RequestCPU,
		RequestMEM:            domAggregateDeployment.RequestMEM,
		LimitCPU:              domAggregateDeployment.LimitCPU,
		LimitMEM:              domAggregateDeployment.LimitMEM,
	}

	// if err := reqDomSvc.Validate(); err != nil {
	// 	return nil, err
	// }
	domAggregateDeployment.SetServiceStatusReplacingModel()

	err = s.repo.Save(domAggregateDeployment)
	if err != nil {
		return nil, err
	}

	domAggregateDeployment.AddModelHistory(resModelPackage.ModelName, resModelPackage.ModelVersion)

	err = domAggregateDeployment.AddEventHistory("ReplaceModel", reqDomSvc.ModelName+" Reason: "+req.Reason, userID)
	if err != nil {
		return nil, err
	}

	err = domAggregateDeployment.RequestReplaceModelInferenceService(s.domInferenceSvc, reqDomSvc)
	if err != nil {
		return nil, err
	}

	domAggregateDeployment.ChangeModelPackage(req.ModelPackageID)

	err = s.repo.Save(domAggregateDeployment)
	if err != nil {
		return nil, err
	}

	//Call Monitoring Service
	//Send Replaced Model Info
	//

	// response dto
	resDTO := new(appDTO.ReplaceModelResponseDTO)
	resDTO.Message = "Replace Model Success"

	return resDTO, nil
}

// ReplaceModel
func (s *DeploymentService) UpdateDeployment(req *appDTO.UpdateDeploymentRequestDTO) (*appDTO.UpdateDeploymentResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	//toBe...
	userID := "testID"

	if err := req.Validate(); err != nil {
		return nil, err
	}

	//Find Domain Entity
	domAggregateDeployment, err := s.repo.GetForUpdate(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		domAggregateDeployment.UpdateDeploymentName(req.Name)
	}
	if req.Description != "" {
		domAggregateDeployment.UpdateDeploymentDescription(req.Description)
	}
	if req.Importance != "" {
		domAggregateDeployment.UpdateDeploymentImportance(req.Importance)
	}

	err = domAggregateDeployment.AddEventHistory("Update", "Deployment is Updated", userID)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(domAggregateDeployment)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.UpdateDeploymentResponseDTO)
	resDTO.Message = "Update Deployment Success"

	return resDTO, nil
}

func (s *DeploymentService) Delete(req *appDTO.DeleteDeploymentRequestDTO) (*appDTO.DeleteDeploymentResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	//toBe...
	userID := "testID"

	//Find Domain Entity
	domAggregateDeployment, err := s.repo.GetByID(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	//Find  PredictionEnv
	predictionEnvInfo, err := s.getPredictionEnvByID(domAggregateDeployment.PredictionEnvID)
	if err != nil {
		return nil, err
	}

	reqDomSvc := domSvcInferenceSvcDto.InferenceServiceDeleteRequest{
		ConnectionInfo: predictionEnvInfo.ConnectionInfo,
		Namespace:      predictionEnvInfo.Namespace,
		DeploymentID:   domAggregateDeployment.ID,
	}

	err = domAggregateDeployment.AddEventHistory("Delete", "Deployment is Deleted", userID)
	if err != nil {
		return nil, err
	}

	err = domAggregateDeployment.RequestDeleteInferenceService(s.domInferenceSvc, reqDomSvc)
	if err != nil {
		return nil, err
	}

	err = s.repo.Delete(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.DeleteDeploymentResponseDTO)
	resDTO.Message = "Deployment Delete Success"

	return resDTO, nil
}

func (s *DeploymentService) GetByID(req *appDTO.GetDeploymentRequestDTO) (*appDTO.GetDeploymentResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	res, err := s.repo.GetByID(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.GetDeploymentResponseDTO)
	resDTO.ID = res.ID
	resDTO.ProjectID = res.ProjectID
	resDTO.ModelPackageID = res.ModelPackageID
	resDTO.PredictionEnvID = res.PredictionEnvID
	resDTO.Name = res.Name
	resDTO.Description = res.Description
	resDTO.Importance = res.Importance
	resDTO.RequestCPU = res.RequestCPU
	resDTO.RequestMEM = res.RequestMEM
	resDTO.ActiveStatus = res.ActiveStatus
	resDTO.ServiceStatus = res.ServiceStatus
	resDTO.ChangeRequested = res.ChangeRequested

	return resDTO, nil
}

// func (s *DeploymentService) GetByName(req *appDTO.GetDeploymentByNametRequestDTO) (*appDTO.GetDeploymentByNameResponseDTO, error) {
// 	// //authorization
// 	// if i.CanAccessCurrentRequest() == false {
// 	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
// 	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
// 	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
// 	// }

// 	res, err := s.repo.GetByName(req.Name)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// response dto

// 	var listDeployment []*appDTO.Deployment
// 	for _, rec := range res {
// 		tmp := new(appDTO.Deployment)

// 		tmp.ID = rec.ID
// 		tmp.ProjectID = rec.ProjectID
// 		tmp.ModelPackageID = rec.ModelPackageID
// 		tmp.PredictionEnvID = rec.PredictionEnvID
// 		tmp.Name = rec.Name
// 		tmp.Description = rec.Description
// 		tmp.Importance = rec.Importance
// 		tmp.RequestCPU = rec.RequestCPU
// 		tmp.RequestMEM = rec.RequestMEM
// 		tmp.ActiveStatus = rec.ActiveStatus
// 		tmp.ServiceStatus = rec.ServiceStatus
// 		tmp.ChangeRequested = rec.ChangeRequested

// 		listDeployment = append(listDeployment, tmp)
// 	}

// 	resDTO := new(appDTO.GetDeploymentByNameResponseDTO)
// 	resDTO.Deployments = listDeployment

// 	return resDTO, nil
// }

func (s *DeploymentService) GetList(req *appDTO.GetDeploymentListRequestDTO) (*appDTO.GetDeploymentListResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	reqP := infRepo.Pagination{
		Limit: req.Limit,
		Page:  req.Page,
		Sort:  req.Sort,
	}

	resultList, pagination, err := s.repo.GetList(req.Name, reqP)
	if err != nil {
		return nil, err
	}

	//interface type을 concrete type으로 변환
	//domain layer에서 pagination type을 모르기 때문에 interface type으로 정의 후 변환한다
	p := pagination.(infRepo.Pagination)

	// response dto
	resDTO := new(appDTO.GetDeploymentListResponseDTO)
	resDTO.Limit = p.Limit
	resDTO.Page = p.Page
	resDTO.TotalRows = p.TotalRows
	resDTO.TotalPages = p.TotalPages

	var listDeployment []*appDTO.DeploymentList
	for _, rec := range resultList {
		tmp := new(appDTO.DeploymentList)

		tmp.ID = rec.ID
		tmp.ProjectID = rec.ProjectID
		tmp.ModelPackageID = rec.ModelPackageID
		tmp.PredictionEnvID = rec.PredictionEnvID
		tmp.Name = rec.Name
		tmp.Importance = rec.Importance

		//Find ModelPackage
		resModelPackage, err := s.getModelPackageByID(rec.ModelPackageID)
		if err != nil {
			return nil, err
		}
		tmp.ModelFrameWork = resModelPackage.ModelFrameWork

		listDeployment = append(listDeployment, tmp)
	}

	resDTO.Rows = listDeployment

	return resDTO, nil
}

func (s *DeploymentService) SetActive(req *appDTO.ActiveDeploymentRequestDTO) (*appDTO.ActiveDeploymentResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	//toBe...
	userID := "testID"

	domAggregateDeployment, err := s.repo.GetForUpdate(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	//Find ModelPackage
	resModelPackage, err := s.getModelPackageByID(domAggregateDeployment.ModelPackageID)
	if err != nil {
		return nil, err
	}

	//Find  PredictionEnv
	resPredictionEnvInfo, err := s.getPredictionEnvByID(domAggregateDeployment.PredictionEnvID)
	if err != nil {
		return nil, err
	}
	reqDom := domSvcInferenceSvcDto.InferenceServiceActiveRequest{
		Namespace:             resPredictionEnvInfo.Namespace,
		DeploymentID:          domAggregateDeployment.ID,
		ModelFrameWork:        resModelPackage.ModelFrameWork,
		ModelFrameWorkVersion: resModelPackage.ModelFrameWorkVersion,
		ModelURL:              resModelPackage.ModelFilePath,
		ConnectionInfo:        resPredictionEnvInfo.ConnectionInfo,
	}

	err = domAggregateDeployment.AddEventHistory("Active", "Deployment is Activated", userID)
	if err != nil {
		return nil, err
	}

	err = domAggregateDeployment.SetDeploymentActive(s.domInferenceSvc, reqDom)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(domAggregateDeployment)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.ActiveDeploymentResponseDTO)
	resDTO.Message = "Deployment Activation Success"

	return resDTO, nil
}

func (s *DeploymentService) SetInActive(req *appDTO.InActiveDeploymentRequestDTO) (*appDTO.InActiveDeploymentResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	//toBe...
	userID := "testID"

	domAggregateDeployment, err := s.repo.GetForUpdate(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	//Find ModelPackage
	resModelPackage, err := s.getModelPackageByID(domAggregateDeployment.ModelPackageID)
	if err != nil {
		return nil, err
	}

	//Find  PredictionEnv
	resPredictionEnvInfo, err := s.getPredictionEnvByID(domAggregateDeployment.PredictionEnvID)
	if err != nil {
		return nil, err
	}
	reqDom := domSvcInferenceSvcDto.InferenceServiceInActiveRequest{
		Namespace:             resPredictionEnvInfo.Namespace,
		DeploymentID:          domAggregateDeployment.ID,
		ModelFrameWork:        resModelPackage.ModelFrameWork,
		ModelFrameWorkVersion: resModelPackage.ModelFrameWorkVersion,
		ModelURL:              resModelPackage.ModelFilePath,
		ConnectionInfo:        resPredictionEnvInfo.ConnectionInfo,
	}

	err = domAggregateDeployment.AddEventHistory("InActive", "Deployment is InActivated", userID)
	if err != nil {
		return nil, err
	}

	err = domAggregateDeployment.SetDeploymentInActive(s.domInferenceSvc, reqDom)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(domAggregateDeployment)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.InActiveDeploymentResponseDTO)
	resDTO.Message = "Deployment InActivation Success"

	return resDTO, nil
}

func (s *DeploymentService) SendPrediction(req *appDTO.SendPredictionRequestDTO) (*appDTO.SendPredictionResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	//Find Domain Entity
	domAggregateDeployment, err := s.repo.GetByID(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	//Find  PredictionEnv
	predictionEnvInfo, err := s.getPredictionEnvByID(domAggregateDeployment.PredictionEnvID)
	if err != nil {
		return nil, err
	}

	host := domAggregateDeployment.ID + "." + predictionEnvInfo.Namespace + "." + predictionEnvInfo.InfereceSvcHostName
	URL := "http://" + predictionEnvInfo.InferenceSvcIngressHost + ":" + predictionEnvInfo.InferenceSvcIngressPort + "/v1/models/" + domAggregateDeployment.ID + ":predict"

	sendResult, err := s.predictionSendSvc.SendPrediction(URL, host, []byte(req.JsonData))
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.SendPredictionResponseDTO)
	resDTO.Message = "Send Success"
	resDTO.PredictionResult = string(sendResult)

	return resDTO, nil
}

func (s *DeploymentService) getModelPackageByID(modelPackageID string) (*appModelPackageDTO.InternalGetModelPackageResponseDTO, error) {
	reqModelPackage := &appModelPackageDTO.InternalGetModelPackageRequestDTO{
		ModelPackageID: modelPackageID,
	}

	resModelPackage, err := s.modelPackageSvc.GetByIDInternal(reqModelPackage)
	if err != nil {
		return nil, err
	}

	return resModelPackage, err
}

func (s *DeploymentService) getPredictionEnvByID(predictionEnvID string) (*domSchema.PredictionEnv, error) {
	//dev mode only
	predictionEnvInfo := &domSchema.PredictionEnv{
		Namespace: "koreserve",
		//ConnectionInfo: "http://192.168.88.161:30070"
		ConnectionInfo:          "http://localhost:5000",
		InfereceSvcHostName:     "kserve.acornsoft.io",
		InferenceSvcIngressHost: "192.168.88.161",
		InferenceSvcIngressPort: "31000",
	}

	return predictionEnvInfo, nil
}

func (s *DeploymentService) GetGovernanceHistory(req *appDTO.GetGovernanceHistoryRequestDTO) (*appDTO.GetGovernanceHistoryResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	//Find Domain Entity
	domAggregateDeployment, err := s.repo.GetGovernanceHistory(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.GetGovernanceHistoryResponseDTO)
	resDTO.EventHistory = domAggregateDeployment.EventHistory

	return resDTO, nil
}

func (s *DeploymentService) GetModelHistory(req *appDTO.GetModelHistoryRequestDTO) (*appDTO.GetModelHistoryResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	//Find Domain Entity
	domAggregateDeployment, err := s.repo.GetByID(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.GetModelHistoryResponseDTO)
	resDTO.ModelHistory = domAggregateDeployment.ModelHistory
	// for _, eh := range domAggregateDeployment.EventHistory {
	// 	mh := appDTO.ModelHistory{}

	// 	mh.DeploymentID = eh.DeploymentID
	// 	mh.EventDate = eh.EventDate
	// 	mh.EventType = eh.EventType
	// 	mh.ID = eh.ID
	// 	mh.LogMessage = eh.LogMessage
	// 	mh.UserID = eh.UserID

	// 	resDTO.ModelHistory = append(resDTO.ModelHistory, mh)
	// }

	return resDTO, nil
}

// func (s *DeploymentService) checkApprovalProcess(req string) (*domSchema.PredictionEnv, error) {
// 	//Check Organization ApprovalPolicy
// 	resApprovalPolicy, err := s.getApprovalPolicy(deploymentID string, orgID string, policyTriggerType string, importance string, userID string)
// 	if err != nil {
// 		return nil, err
// 	}
// }
