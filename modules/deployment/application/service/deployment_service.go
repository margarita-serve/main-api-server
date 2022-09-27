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
	"fmt"

	common "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/common"
	appDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/application/dto"
	domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/entity"
	domRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/repository"
	domSchema "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/schema"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
	"github.com/rs/xid"

	domSvcInferenceSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/service/inference_service"
	domSvcInferenceSvcDto "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/service/inference_service/dto"
	infInfSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/infrastructure/inference_service/kserve"
	predictionSendSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/infrastructure/prediction_sender"
	infRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/infrastructure/repository"
	appResourceDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/resource/application/dto"
)

type IPredictionEnvService interface {
	GetByIDInternal(req *appResourceDTO.InternalGetPredictionEnvRequestDTO, i identity.Identity) (*appResourceDTO.InternalGetPredictionEnvResponseDTO, error)
}

// DeploymentService type
type DeploymentService struct {
	BaseService
	domInferenceSvc domSvcInferenceSvc.IInferenceServiceAdapter
	projectSvc      common.IProjectService
	modelPackageSvc common.IModelPackageService
	monitoringSvc   common.IMonitorService
	//predictionEnvSvc  IPredictionEnvService
	predictionSendSvc *predictionSendSvc.PredictionSender
	publisher         common.EventPublisher
	repo              domRepo.IDeploymentRepo
}

// NewDeploymentService new DeploymentService
func NewDeploymentService(h *handler.Handler, projectSvc common.IProjectService, modelPackageSvc common.IModelPackageService, monitorSvc common.IMonitorService, publisher common.EventPublisher) (*DeploymentService, error) {
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

	if svc.predictionSendSvc, err = predictionSendSvc.NewPredictionSendService(); err != nil {
		return nil, err
	}

	svc.modelPackageSvc = modelPackageSvc
	svc.monitoringSvc = monitorSvc
	svc.projectSvc = projectSvc
	//svc.predictionEnvSvc = predictionEnvSvc
	svc.publisher = publisher

	return svc, nil
}

// Create
func (s *DeploymentService) Create(req *appDTO.CreateDeploymentRequestDTO, i identity.Identity) (*appDTO.CreateDeploymentResponseDTO, error) {
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

	//Find ModelPackage
	resModelPackage, err := s.getModelPackageByID(req.ModelPackageID)
	if err != nil {
		return nil, err
	}

	//to be...
	if req.PredictionEnvID == "" {
		//Find  Project Default PredictionEnv
		req.PredictionEnvID = "dev01234567890123456"
	}

	// New deployment domain Instance
	domAggregateDeployment, err := domEntity.NewDeployment(
		guid,
		resModelPackage.ProjectID,
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
		i.Claims.Username,
	)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(domAggregateDeployment)
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
		ConnectionInfo:        resPredictionEnvInfo.ClusterInfo.InferenceSvcInfo.InfereceSvcAPISvrEndPoint,
		RequestCPU:            domAggregateDeployment.RequestCPU,
		RequestMEM:            domAggregateDeployment.RequestMEM,
		LimitCPU:              domAggregateDeployment.LimitCPU,
		LimitMEM:              domAggregateDeployment.LimitMEM,
	}

	if err := reqDomSvc.Validate(); err != nil {
		return nil, err
	}

	newModelHistoryID := xid.New().String()
	domAggregateDeployment.AddModelHistory(newModelHistoryID, resModelPackage.ModelName, resModelPackage.ModelVersion, resModelPackage.ModelPackageID)

	featureDriftTrackingBool := false
	accuracyAnalyzeBool := false
	if req.FeatureDriftTracking {
		featureDriftTrackingBool = req.FeatureDriftTracking
	}
	if req.AccuracyAnalyze {
		accuracyAnalyzeBool = req.AccuracyAnalyze
	}

	// WaitGroup 생성. 2개의 Go루틴을 기다림.
	// var wait sync.WaitGroup
	// var checkErrMsg error
	// wait.Add(2)

	// // ch생성
	// errs := make(chan error, 1)

	// go func() {
	// 	defer wait.Done() //끝나면 .Done() 호출
	// 	_, err = s.monitoringSvc.Create(reqMonitoring)
	// 	if err != nil {
	// 		errs <- err
	// 	}

	// }()

	// go func() {
	// 	defer wait.Done() //끝나면 .Done() 호출
	// 	err = domAggregateDeployment.RequestCreateInferenceService(s.domInferenceSvc, reqDomSvc)
	// 	if err != nil {
	// 		errs <- err
	// 	}
	// }()

	// wait.Wait() //Go루틴 모두 끝날 때까지 대기
	// close(errs)

	// checkErrMsg = <-errs

	// if checkErrMsg != nil {
	// 	reqDeleteInference := &appDTO.DeleteDeploymentRequestDTO{
	// 		//ProjectID:    req.ProjectID,
	// 		DeploymentID: guid,
	// 	}

	// 	_, err := s.Delete(reqDeleteInference, i)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("deployment create error: %s, %s", checkErrMsg, err)
	// 	}

	// 	return nil, fmt.Errorf("deployment create error: %s", checkErrMsg)
	// }

	err = domAggregateDeployment.RequestCreateInferenceService(s.domInferenceSvc, reqDomSvc)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(domAggregateDeployment)
	if err != nil {
		return nil, err
	}

	//Create Monitoring Service
	// reqMonitoring := &appMonitoringDTO.MonitorCreateRequestDTO{
	// 	DeploymentID:         domAggregateDeployment.ID,
	// 	ModelPackageID:       domAggregateDeployment.ModelPackageID,
	// 	FeatureDriftTracking: featureDriftTrackingBool,
	// 	AccuracyMonitoring:   accuracyAnalyzeBool,
	// 	AssociationID:        &req.AssociationID,
	// 	//toBe..
	// 	//AssociationIDInFeature:        &req.AssociationIDInFeature,
	// 	ModelHistoryID: newModelHistoryID,
	// }

	//s.addModelPackageDeployCount(req.ModelPackageID)
	s.publisher.Notify(common.NewEventDeploymentInferenceServiceCreated(domAggregateDeployment.ID, domAggregateDeployment.ModelPackageID, featureDriftTrackingBool, accuracyAnalyzeBool, req.AssociationID, req.AssociationIDInFeature, newModelHistoryID))

	//Find Domain Entity
	checkStatus, err := s.repo.GetByIDInternal(domAggregateDeployment.ID)
	if err != nil {
		return nil, err
	}

	//checkErrMsg = <-errs

	if checkStatus.ServiceStatus != "Ready" {
		reqDeleteInference := &appDTO.DeleteDeploymentRequestDTO{
			//ProjectID:    req.ProjectID,
			DeploymentID: guid,
		}

		_, _ = s.Delete(reqDeleteInference, i)
		return nil, fmt.Errorf("deployment create error: check monitoring system")
	}

	err = domAggregateDeployment.AddEventHistory("Create", "Deployment Created", i.Claims.Username)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(domAggregateDeployment)
	if err != nil {
		return nil, err
	}

	s.publisher.Notify(common.NewEventDeploymentCreated(domAggregateDeployment.ID, domAggregateDeployment.ModelPackageID, featureDriftTrackingBool, accuracyAnalyzeBool, req.AssociationID, req.AssociationIDInFeature, newModelHistoryID))

	// response dto
	resDTO := new(appDTO.CreateDeploymentResponseDTO)
	resDTO.DeploymentID = domAggregateDeployment.ID

	return resDTO, nil
}

// ReplaceModel
func (s *DeploymentService) ReplaceModel(req *appDTO.ReplaceModelRequestDTO, i identity.Identity) (*appDTO.ReplaceModelResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	if err := req.Validate(); err != nil {
		return nil, err
	}

	listProjectId, err := s.checkProjectList(i)
	if err != nil {
		return nil, err
	}

	//Find Domain Entity
	domAggregateDeployment, err := s.repo.GetByID(req.DeploymentID, listProjectId)
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
		ConnectionInfo:        resPredictionEnvInfo.ClusterInfo.InferenceSvcInfo.InfereceSvcAPISvrEndPoint,
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

	newModelHistoryID := xid.New().String()
	domAggregateDeployment.AddModelHistory(newModelHistoryID, resModelPackage.ModelName, resModelPackage.ModelVersion, resModelPackage.ModelPackageID)

	// //Call Monitoring Service
	// //Send Replaced Model Info 수정
	// reqReplaceMonitoring := &appMonitoringDTO.MonitorReplaceModelRequestDTO{
	// 	DeploymentID:   req.DeploymentID,
	// 	ModelPackageID: req.ModelPackageID,
	// 	ModelHistoryID: newModelHistoryID,
	// }

	// // WaitGroup 생성. 2개의 Go루틴을 기다림.
	// var wait sync.WaitGroup
	// var checkErrMsg error
	// wait.Add(2)

	// // ch생성
	// errs := make(chan error, 1)

	// go func() {
	// 	defer wait.Done() //끝나면 .Done() 호출
	// 	_, err = s.monitoringSvc.MonitorReplaceModel(reqReplaceMonitoring)
	// 	if err != nil {
	// 		errs <- err
	// 	}

	// }()

	// go func() {
	// 	defer wait.Done() //끝나면 .Done() 호출
	// 	err = domAggregateDeployment.RequestReplaceModelInferenceService(s.domInferenceSvc, reqDomSvc)
	// 	if err != nil {
	// 		errs <- err
	// 	}
	// }()

	// wait.Wait() //Go루틴 모두 끝날 때까지 대기
	// close(errs)

	// checkErrMsg = <-errs

	// if checkErrMsg != nil {
	// 	domAggregateDeployment.SetServiceStatusReady()

	// 	err = s.repo.Save(domAggregateDeployment)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	return nil, fmt.Errorf("deployment replace error: %s", checkErrMsg)
	// }

	err = domAggregateDeployment.RequestReplaceModelInferenceService(s.domInferenceSvc, reqDomSvc)
	if err != nil {
		return nil, err
	}

	err = domAggregateDeployment.AddEventHistory("ReplaceModel", reqDomSvc.ModelName+" Reason: "+req.Reason, i.Claims.Username)
	if err != nil {
		return nil, err
	}

	domAggregateDeployment.ChangeModelPackage(req.ModelPackageID)

	err = s.repo.Save(domAggregateDeployment)
	if err != nil {
		return nil, err
	}

	//s.addModelPackageDeployCount(req.ModelPackageID)
	s.publisher.Notify(common.NewEventDeploymentModelReplaced(domAggregateDeployment.ID, domAggregateDeployment.ModelPackageID, newModelHistoryID))

	// response dto
	resDTO := new(appDTO.ReplaceModelResponseDTO)
	resDTO.Message = "Replace Model Success"

	return resDTO, nil
}

// UpdateResources
func (s *DeploymentService) updateResources(domAggregateDeployment *domEntity.Deployment, i identity.Identity) error {
	//Find ModelPackage
	resModelPackage, err := s.getModelPackageByID(domAggregateDeployment.ModelPackageID)
	if err != nil {
		return err
	}

	//Find  PredictionEnv
	resPredictionEnvInfo, err := s.getPredictionEnvByID(domAggregateDeployment.PredictionEnvID)
	if err != nil {
		return err
	}

	reqDomSvc := domSvcInferenceSvcDto.InferenceServiceReplaceModelRequest{
		Namespace:             resPredictionEnvInfo.Namespace,
		DeploymentID:          domAggregateDeployment.ID,
		ModelFrameWork:        resModelPackage.ModelFrameWork,
		ModelFrameWorkVersion: resModelPackage.ModelFrameWorkVersion,
		ModelURL:              resModelPackage.ModelFilePath,
		ModelName:             resModelPackage.ModelName,
		ConnectionInfo:        resPredictionEnvInfo.ClusterInfo.InferenceSvcInfo.InfereceSvcAPISvrEndPoint,
		RequestCPU:            domAggregateDeployment.RequestCPU,
		RequestMEM:            domAggregateDeployment.RequestMEM,
		LimitCPU:              domAggregateDeployment.LimitCPU,
		LimitMEM:              domAggregateDeployment.LimitMEM,
	}

	// domAggregateDeployment.SetServiceStatusReplacingModel()

	err = domAggregateDeployment.RequestReplaceModelInferenceService(s.domInferenceSvc, reqDomSvc)
	if err != nil {
		return err
	}

	return nil
}

// ReplaceModel
func (s *DeploymentService) UpdateDeployment(req *appDTO.UpdateDeploymentRequestDTO, i identity.Identity) (*appDTO.UpdateDeploymentResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	if err := req.Validate(); err != nil {
		return nil, err
	}

	//Find Domain Entity
	domAggregateDeployment, err := s.repo.GetForUpdate(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	resModelPackage, err := s.getModelPackageByID(domAggregateDeployment.ModelPackageID)
	if err != nil {
		return nil, err
	}

	var updateInfoString string = ""

	if req.Name != nil {
		updateInfoString += fmt.Sprintf("%s: %s -> %s  ", "Name", domAggregateDeployment.Name, *req.Name)
		domAggregateDeployment.UpdateDeploymentName(*req.Name)
	}
	if req.Description != nil {
		updateInfoString += fmt.Sprintf("%s: %s -> %s  ", "Description", domAggregateDeployment.Description, *req.Description)
		domAggregateDeployment.UpdateDeploymentDescription(*req.Description)
	}
	if req.Importance != nil {
		updateInfoString += fmt.Sprintf("%s: %s ->% s  ", "Importance", domAggregateDeployment.Importance, *req.Importance)
		domAggregateDeployment.UpdateDeploymentImportance(*req.Importance)
	}
	// if req.AssociationID != nil {
	// 	reqUpdateAssociationID := new(appMonitoringDTO.UpdateAssociationIDRequestDTO)
	// 	reqUpdateAssociationID.DeploymentID = req.DeploymentID
	// 	reqUpdateAssociationID.AssociationID = req.AssociationID
	// 	//toBe..
	// 	//reqUpdateAssociationID.AssociationIDInFeature = req.AssociationIDInFeature

	// 	_, err = s.monitoringSvc.UpdateAssociationID(reqUpdateAssociationID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	// if req.FeatureDriftTracking != nil {
	// 	if *req.FeatureDriftTracking {
	// 		println("FeatureDriftTracking true")

	// 		reqDriftActive := new(appMonitoringDTO.MonitorDriftActiveRequestDTO)
	// 		reqDriftActive.DeploymentID = req.DeploymentID
	// 		reqDriftActive.ModelPackageID = resModelPackage.ModelPackageID
	// 		reqDriftActive.CurrentModelID = currentModelID

	// 		_, err = s.monitoringSvc.SetDriftMonitorActive(reqDriftActive)
	// 		if err != nil {
	// 			return nil, err
	// 		}

	// 	} else {
	// 		reqDriftInActive := new(appMonitoringDTO.MonitorDriftInActiveRequestDTO)
	// 		reqDriftInActive.DeploymentID = req.DeploymentID

	// 		_, err = s.monitoringSvc.SetDriftMonitorInActive(reqDriftInActive)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 	}
	// }
	// if req.AccuracyAnalyze != nil {
	// 	if *req.AccuracyAnalyze {
	// 		println("AccuracyAnalyze true")
	// 		reqAccuracyActive := new(appMonitoringDTO.MonitorAccuracyActiveRequestDTO)
	// 		reqAccuracyActive.DeploymentID = req.DeploymentID
	// 		reqAccuracyActive.ModelPackageID = resModelPackage.ModelPackageID
	// 		reqAccuracyActive.AssociationID = req.AssociationID
	// 		reqAccuracyActive.CurrentModelID = currentModelID
	// 		//toBe..
	// 		//reqAccuracyActive.AssociationIDInFeature = req.AssociationIDInFeature

	// 		_, err = s.monitoringSvc.SetAccuracyMonitorActive(reqAccuracyActive)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 	} else {
	// 		reqAccuracyInActive := new(appMonitoringDTO.MonitorAccuracyInActiveRequestDTO)
	// 		reqAccuracyInActive.DeploymentID = req.DeploymentID

	// 		_, err = s.monitoringSvc.SetAccuracyMonitorInActive(reqAccuracyInActive)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 	}
	// }

	if (req.RequestCPU != nil) || (req.RequestMEM != nil) {
		if req.RequestCPU != nil {
			updateInfoString += fmt.Sprintf("%s: %f->%f\n", "RequestCPU", domAggregateDeployment.RequestCPU, *req.RequestCPU)
			domAggregateDeployment.UpdateDeploymentRequestCPU(*req.RequestCPU)
		}
		if req.RequestMEM != nil {
			updateInfoString += fmt.Sprintf("%s: %f->%f\n", "RequestMEM", domAggregateDeployment.RequestMEM, *req.RequestMEM)
			domAggregateDeployment.UpdateDeploymentRequestMEM(*req.RequestMEM)
		}

		err := s.updateResources(domAggregateDeployment, i)
		if err != nil {
			return nil, err
		}
	}

	err = domAggregateDeployment.AddEventHistory("Update", fmt.Sprintf("%s (%s)", "Deployment is Updated", updateInfoString), i.Claims.Username)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(domAggregateDeployment)
	if err != nil {
		return nil, err
	}

	var currentModelID string
	for _, history := range domAggregateDeployment.ModelHistory {
		if history.ApplyHistoryTag == "Current" {
			currentModelID = history.ID
			break
		}
	}

	// Send Event
	//if req.AssociationID != nil {
	//	s.publisher.Notify(common.NewEventDeploymentAssociationIDUpdated(domAggregateDeployment.ID, resModelPackage.ModelPackageID, currentModelID, *req.AssociationID))
	//}

	if req.FeatureDriftTracking != nil {
		if *req.FeatureDriftTracking {
			s.publisher.Notify(common.NewEventDeploymentFeatureDriftTrackingEnabled(domAggregateDeployment.ID, resModelPackage.ModelPackageID, currentModelID))
		} else {
			s.publisher.Notify(common.NewEventDeploymentFeatureDriftTrackingDisabled(domAggregateDeployment.ID, resModelPackage.ModelPackageID, currentModelID))
		}
	}

	if req.AccuracyAnalyze != nil {
		if *req.AccuracyAnalyze {
			if req.AssociationID == nil {
				s.publisher.Notify(common.NewEventDeploymentAccuracyAnalyzeEnabled(domAggregateDeployment.ID, resModelPackage.ModelPackageID, currentModelID, "None", false))
			} else {
				s.publisher.Notify(common.NewEventDeploymentAccuracyAnalyzeEnabled(domAggregateDeployment.ID, resModelPackage.ModelPackageID, currentModelID, *req.AssociationID, *req.AssociationIDInFeature))
			}
		} else {
			s.publisher.Notify(common.NewEventDeploymentAccuracyAnalyzeDisabled(domAggregateDeployment.ID, resModelPackage.ModelPackageID, currentModelID))
		}
	}

	// response dto
	resDTO := new(appDTO.UpdateDeploymentResponseDTO)
	resDTO.Message = "Update Deployment Success"

	return resDTO, nil
}

func (s *DeploymentService) Delete(req *appDTO.DeleteDeploymentRequestDTO, i identity.Identity) (*appDTO.DeleteDeploymentResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	listProjectId, err := s.checkProjectList(i)
	if err != nil {
		return nil, err
	}

	//Find Domain Entity
	domAggregateDeployment, err := s.repo.GetByID(req.DeploymentID, listProjectId)
	if err != nil {
		return nil, err
	}

	//Find  PredictionEnv
	predictionEnvInfo, err := s.getPredictionEnvByID(domAggregateDeployment.PredictionEnvID)
	if err != nil {
		return nil, err
	}

	reqDomSvc := domSvcInferenceSvcDto.InferenceServiceDeleteRequest{
		ConnectionInfo: predictionEnvInfo.ClusterInfo.InferenceSvcInfo.InfereceSvcAPISvrEndPoint,
		Namespace:      predictionEnvInfo.Namespace,
		DeploymentID:   domAggregateDeployment.ID,
	}

	err = domAggregateDeployment.AddEventHistory("Delete", "Deployment is Deleted", i.Claims.Username)
	if err != nil {
		return nil, err
	}

	err = domAggregateDeployment.RequestDeleteInferenceService(s.domInferenceSvc, reqDomSvc)
	if err != nil {
		return nil, err
	}

	// reqDeleteMonitoring := &appMonitoringDTO.MonitorDeleteRequestDTO{
	// 	DeploymentID: domAggregateDeployment.ID,
	// }

	// _, err = s.monitoringSvc.Delete(reqDeleteMonitoring)
	// if err != nil {
	// 	//return nil, fmt.Errorf("monitoring delete error: %s", err)
	// 	fmt.Errorf("monitoring delete error: %s", err)
	// }

	err = s.repo.Delete(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	s.publisher.Notify(common.NewEventDeploymentDeleted(domAggregateDeployment.ID))

	resDTO := new(appDTO.DeleteDeploymentResponseDTO)
	resDTO.Message = "Deployment Delete Success"

	return resDTO, nil
}

func (s *DeploymentService) GetByID(req *appDTO.GetDeploymentRequestDTO, i identity.Identity) (*appDTO.GetDeploymentResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	listProjectId, err := s.checkProjectList(i)
	if err != nil {
		return nil, err
	}

	res, err := s.repo.GetByID(req.DeploymentID, listProjectId)
	if err != nil {
		return nil, err
	}

	//Find Monitor
	resMonitor, _ := s.getMonitorByID(res.ID)

	//Find ModelPackage
	resModelPackage, _ := s.getModelPackageByID(res.ModelPackageID)
	// if err != nil {
	// 	return nil, err
	// }

	//Find PredictionEnv
	resPredictionEnv, _ := s.getPredictionEnvByID(res.PredictionEnvID)
	// if err != nil {
	// 	return nil, err
	// }

	//Find ProjectInfo
	resProject, _ := s.getProjectByID(res.ProjectID)
	// if err != nil {
	// 	return nil, err
	// }

	// response dto
	resDTO := new(appDTO.GetDeploymentResponseDTO)
	resDTO.ID = res.ID
	resDTO.ProjectID = res.ProjectID
	resDTO.ProjectName = resProject.Name
	resDTO.ModelPackageID = res.ModelPackageID
	resDTO.PredictionEnvID = res.PredictionEnvID
	resDTO.PredictionEnvName = resPredictionEnv.Name
	resDTO.Name = res.Name
	resDTO.Description = res.Description
	resDTO.Importance = res.Importance
	resDTO.RequestCPU = res.RequestCPU
	resDTO.RequestMEM = res.RequestMEM
	resDTO.ActiveStatus = res.ActiveStatus
	resDTO.ServiceStatus = res.ServiceStatus
	resDTO.ChangeRequested = res.ChangeRequested

	if resModelPackage != nil {
		resDTO.ModelPackageName = resModelPackage.Name
	}

	if resMonitor != nil {
		resDTO.DriftStatus = resMonitor.DriftStatus
		resDTO.AccuracyStatus = resMonitor.AccuracyStatus
		resDTO.ServiceHealthStatus = resMonitor.ServiceHealthStatus
		resDTO.FeatureDriftTracking = resMonitor.FeatureDriftTracking
		resDTO.AccuracyAnalyze = resMonitor.AccuracyMonitoring
		resDTO.AssociationID = resMonitor.AssociationID
	}
	//toBe..
	//resDTO.AssociationIDInFeature = resMonitor.Monitor.AssociationIDInFeature

	return resDTO, nil
}

func (s *DeploymentService) GetByIDInternal(deploymentID string) (*common.InternalGetByIDInternalResponse, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	res, err := s.repo.GetByIDInternal(deploymentID)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(common.InternalGetByIDInternalResponse)
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

func (s *DeploymentService) GetList(req *appDTO.GetDeploymentListRequestDTO, i identity.Identity) (*appDTO.GetDeploymentListResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	listProjectId, err := s.checkProjectList(i)
	if err != nil {
		return nil, err
	}

	reqP := infRepo.Pagination{
		Limit: req.Limit,
		Page:  req.Page,
		Sort:  req.Sort,
	}

	resultList, pagination, err := s.repo.GetList(req.Name, reqP, listProjectId)
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
		//Find Monitor
		resMonitor, _ := s.getMonitorByID(rec.ID)

		//Find ModelPackage
		resModelPackage, _ := s.getModelPackageByID(rec.ModelPackageID)
		// if err != nil {
		// 	return nil, err
		// }

		//Find PredictionEnv
		resPredictionEnv, _ := s.getPredictionEnvByID(rec.PredictionEnvID)
		// if err != nil {
		// 	return nil, err
		// }

		//Find ProjectInfo
		resProject, _ := s.getProjectByID(rec.ProjectID)
		// if err != nil {
		// 	return nil, err
		// }

		tmp := new(appDTO.DeploymentList)

		tmp.ID = rec.ID
		tmp.ProjectID = rec.ProjectID
		tmp.ProjectName = resProject.Name
		tmp.ModelPackageID = rec.ModelPackageID
		tmp.PredictionEnvID = rec.PredictionEnvID
		tmp.PredictionEnvName = resPredictionEnv.Name
		tmp.Name = rec.Name
		tmp.Importance = rec.Importance
		tmp.ActiveStatus = rec.ActiveStatus
		tmp.ServiceStatus = rec.ServiceStatus

		if resModelPackage != nil {
			tmp.ModelPackageName = resModelPackage.Name
			tmp.ModelFrameWork = resModelPackage.ModelFrameWork
		}

		if resMonitor != nil {
			tmp.DriftStatus = resMonitor.DriftStatus
			tmp.AccuracyStatus = resMonitor.AccuracyStatus
			tmp.ServiceHealthStatus = resMonitor.ServiceHealthStatus
		}

		listDeployment = append(listDeployment, tmp)
	}

	resDTO.Rows = listDeployment

	return resDTO, nil
}

func (s *DeploymentService) SetActive(req *appDTO.ActiveDeploymentRequestDTO, i identity.Identity) (*appDTO.ActiveDeploymentResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

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
		ConnectionInfo:        resPredictionEnvInfo.ClusterInfo.InferenceSvcInfo.InfereceSvcAPISvrEndPoint,
		RequestCPU:            domAggregateDeployment.RequestCPU,
		LimitCPU:              domAggregateDeployment.LimitCPU,
		RequestMEM:            domAggregateDeployment.RequestMEM,
		LimitMEM:              domAggregateDeployment.LimitMEM,
	}

	err = domAggregateDeployment.AddEventHistory("Active", "Deployment is Activated", i.Claims.Username)
	if err != nil {
		return nil, err
	}

	err = domAggregateDeployment.SetDeploymentActive(s.domInferenceSvc, reqDom)
	if err != nil {
		return nil, err
	}

	// //monitor active
	// reqMonitor := &appMonitoringDTO.MonitorGetByIDRequestDTO{
	// 	ID: req.DeploymentID,
	// }

	// resMonitor, err := s.monitoringSvc.GetByID(reqMonitor)

	// reqServiceHealth := &appMonitoringDTO.MonitorServiceHealthActiveRequestDTO{
	// 	DeploymentID:   req.DeploymentID,
	// 	CurrentModelID: "",
	// }
	// _, err = s.monitoringSvc.SetServiceHealthMonitorActive(reqServiceHealth)
	// if err != nil {
	// 	return nil, err
	// }
	// if resMonitor.Monitor.FeatureDriftTracking == true {
	// 	reqDrift := &appMonitoringDTO.MonitorDriftActiveRequestDTO{
	// 		DeploymentID:   req.DeploymentID,
	// 		ModelPackageID: "",
	// 		CurrentModelID: "",
	// 	}
	// 	_, err = s.monitoringSvc.SetDriftMonitorActive(reqDrift)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }
	// if resMonitor.Monitor.AccuracyMonitoring == true {
	// 	reqAccuracy := &appMonitoringDTO.MonitorAccuracyActiveRequestDTO{
	// 		DeploymentID:   req.DeploymentID,
	// 		ModelPackageID: "",
	// 		AssociationID:  nil,
	// 		CurrentModelID: "",
	// 	}
	// 	_, err = s.monitoringSvc.SetAccuracyMonitorActive(reqAccuracy)
	// }
	// if err != nil {
	// 	return nil, err
	// }

	err = s.repo.Save(domAggregateDeployment)
	if err != nil {
		return nil, err
	}

	s.publisher.Notify(common.NewEventDeploymentActived(domAggregateDeployment.ID))

	// response dto
	resDTO := new(appDTO.ActiveDeploymentResponseDTO)
	resDTO.Message = "Deployment Activation Success"

	return resDTO, nil
}

func (s *DeploymentService) SetInActive(req *appDTO.InActiveDeploymentRequestDTO, i identity.Identity) (*appDTO.InActiveDeploymentResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

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
		ConnectionInfo:        resPredictionEnvInfo.ClusterInfo.InferenceSvcInfo.InfereceSvcAPISvrEndPoint,
	}

	err = domAggregateDeployment.AddEventHistory("InActive", "Deployment is InActivated", i.Claims.Username)
	if err != nil {
		return nil, err
	}

	err = domAggregateDeployment.SetDeploymentInActive(s.domInferenceSvc, reqDom)
	if err != nil {
		return nil, err
	}

	// //monitor inactive
	// reqMonitor := &appMonitoringDTO.MonitorGetByIDRequestDTO{
	// 	ID: req.DeploymentID,
	// }

	// resMonitor, err := s.monitoringSvc.GetByID(reqMonitor)
	// reqServiceHealth := &appMonitoringDTO.MonitorServiceHealthInActiveRequestDTO{
	// 	DeploymentID: req.DeploymentID,
	// }
	// _, err = s.monitoringSvc.SetServiceHealthMonitorInActive(reqServiceHealth)
	// if err != nil {
	// 	return nil, err
	// }
	// if resMonitor.Monitor.FeatureDriftTracking == true {
	// 	reqDrift := &appMonitoringDTO.MonitorDriftInActiveRequestDTO{
	// 		DeploymentID: req.DeploymentID,
	// 	}
	// 	_, err = s.monitoringSvc.SetDriftMonitorInActive(reqDrift)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }
	// if resMonitor.Monitor.AccuracyMonitoring == true {
	// 	reqAccuracy := &appMonitoringDTO.MonitorAccuracyInActiveRequestDTO{
	// 		DeploymentID: req.DeploymentID,
	// 	}
	// 	_, err = s.monitoringSvc.SetAccuracyMonitorInActive(reqAccuracy)
	// }
	// if err != nil {
	// 	return nil, err
	// }

	err = s.repo.Save(domAggregateDeployment)
	if err != nil {
		return nil, err
	}

	s.publisher.Notify(common.NewEventDeploymentInActived(domAggregateDeployment.ID))

	// response dto
	resDTO := new(appDTO.InActiveDeploymentResponseDTO)
	resDTO.Message = "Deployment InActivation Success"

	return resDTO, nil
}

func (s *DeploymentService) SendPrediction(req *appDTO.SendPredictionRequestDTO, i identity.Identity) (*appDTO.SendPredictionResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	listProjectId, err := s.checkProjectList(i)
	if err != nil {
		return nil, err
	}

	//Find Domain Entity
	domAggregateDeployment, err := s.repo.GetByID(req.DeploymentID, listProjectId)
	if err != nil {
		return nil, err
	}

	if domAggregateDeployment.ActiveStatus == "InActive" {
		return nil, fmt.Errorf("deployment is inactive")
	}

	//Find  PredictionEnv
	predictionEnvInfo, err := s.getPredictionEnvByID(domAggregateDeployment.PredictionEnvID)
	if err != nil {
		return nil, err
	}

	host := domAggregateDeployment.ID + "." + predictionEnvInfo.Namespace + "." + predictionEnvInfo.ClusterInfo.InferenceSvcInfo.InfereceSvcHostName
	URL := predictionEnvInfo.ClusterInfo.InferenceSvcInfo.InferenceSvcIngressEndPoint + "/v1/models/" + domAggregateDeployment.ID + ":predict"

	sendResult, err := s.predictionSendSvc.SendPrediction(URL, host, []byte(req.JsonData))
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.SendPredictionResponseDTO)
	resDTO.Message = "Send Success"
	resDTO.PredictionResult = string(sendResult)

	return resDTO, nil
}

func (s *DeploymentService) getModelPackageByID(modelPackageID string) (*common.InternalGetModelPackageResponseDTO, error) {
	resModelPackage, err := s.modelPackageSvc.GetByIDInternal(modelPackageID)
	if err != nil {
		return nil, err
	}

	return resModelPackage, err
}

func (s *DeploymentService) getPredictionEnvByID(predictionEnvID string) (*domSchema.PredictionEnv, error) {
	// req := &appResourceDTO.InternalGetPredictionEnvRequestDTO{
	// 	PredictionEnvID: predictionEnvID,
	// }

	// res, err := s.predictionEnvSvc.GetByIDInternal(req, i)
	// if err != nil {
	// 	return nil, err
	// }

	// //dev mode only
	// predictionEnvInfo := &domSchema.PredictionEnv{
	// 	Namespace:                   res.Namespace,
	// 	InfereceSvcAPISvrEndPoint:   res.ClusterInfo.InferenceSvcInfo.InfereceSvcAPISvrEndPoint,
	// 	InfereceSvcHostName:         res.ClusterInfo.InferenceSvcInfo.InfereceSvcHostName,
	// 	InferenceSvcIngressEndPoint: res.ClusterInfo.InferenceSvcInfo.InferenceSvcIngressEndPoint,
	// }

	//dev mode only
	predictionEnvInfo := &domSchema.PredictionEnv{
		Name:            "Default",
		PredictionEnvID: "test",
		Namespace:       "koreserve",
		ClusterInfo: domSchema.ClusterInfo{
			InferenceSvcInfo: domSchema.InferenceSvcInfo{
				InfereceSvcAPISvrEndPoint:   "http://192.168.88.161:30070",
				InfereceSvcHostName:         "kserve.acornsoft.io",
				InferenceSvcIngressEndPoint: "http://192.168.88.161:31000",
			},
		},
	}

	return predictionEnvInfo, nil
}

func (s *DeploymentService) GetGovernanceHistory(req *appDTO.GetGovernanceHistoryRequestDTO, i identity.Identity) (*appDTO.GetGovernanceHistoryResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	// listProjectId, err := s.checkProjectList(i)
	// if err != nil {
	// 	return nil, err
	// }

	//Find Domain Entity
	EventHistory, err := s.repo.GetGovernance(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.GetGovernanceHistoryResponseDTO)
	resDTO.EventHistory = EventHistory

	return resDTO, nil
}

func (s *DeploymentService) GetModelHistory(req *appDTO.GetModelHistoryRequestDTO, i identity.Identity) (*appDTO.GetModelHistoryResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	listProjectId, err := s.checkProjectList(i)
	if err != nil {
		return nil, err
	}

	//Find Domain Entity
	domAggregateDeployment, err := s.repo.GetByID(req.DeploymentID, listProjectId)
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

func (s *DeploymentService) checkProjectList(i identity.Identity) ([]string, error) {
	var listProjectId []string

	projectRes, err := s.projectSvc.GetListInternal(i.Claims.Username)
	if err != nil {
		return listProjectId, err
	}

	projectIdList := projectRes.Rows

	for _, rec := range projectIdList {
		listProjectId = append(listProjectId, rec.ProjectID)
	}
	return listProjectId, nil
}

func (s *DeploymentService) getProjectByID(projectID string) (*common.GetProjectInternalResponseDTO, error) {
	resProject, err := s.projectSvc.GetByIDInternal(projectID)
	if err != nil {
		return nil, err
	}

	return resProject, err
}

func (s *DeploymentService) getMonitorByID(monitorID string) (*common.MonitorGetByIDInternalResponseDTO, error) {
	// reqMonitor := &appMonitoringDTO.MonitorGetByIDRequestDTO{
	// 	ID: monitorID,
	// }

	resMonitor, err := s.monitoringSvc.GetByIDInternal(monitorID)
	if err != nil {
		return nil, err
	}

	return resMonitor, err
}

func (s *DeploymentService) Update(event common.Event) {
	switch actualEvent := event.(type) {
	case common.MonitoringCreated:
		//
		fmt.Sprintf("MonitoringCreated %s", actualEvent.DeploymentID())

		//Find Domain Entity
		domAggregateDeployment, _ := s.repo.GetByIDInternal(actualEvent.DeploymentID())
		domAggregateDeployment.SetServiceStatusReady()
		err := s.repo.Save(domAggregateDeployment)
		fmt.Printf("err: %v\n", err)
	case common.MonitoringCreateFailed:
		//
		fmt.Sprintf("MonitoringCreateFailed %s", actualEvent.DeploymentID())

		domAggregateDeployment, _ := s.repo.GetByIDInternal(actualEvent.DeploymentID())
		domAggregateDeployment.SetServiceStatusError()
		err := s.repo.Save(domAggregateDeployment)
		fmt.Printf("err: %v\n", err)
	case common.MonitoringAccuracyStatusChangedToFailing:
		//
		s.AddGovernanceHistory(actualEvent.DeploymentID(), "AccuracyAlert", "failing")
	case common.MonitoringAccuracyStatusChangedToAtrisk:
		//
		s.AddGovernanceHistory(actualEvent.DeploymentID(), "AccuracyAlert", "atrisk")
	case common.MonitoringDataDriftStatusChangedToFailing:
		//
		s.AddGovernanceHistory(actualEvent.DeploymentID(), "DataDriftAlert", "failing")
	case common.MonitoringDataDriftStatusChangedToAtrisk:
		//
		s.AddGovernanceHistory(actualEvent.DeploymentID(), "DataDriftAlert", "atrisk")
	case common.MonitoringServiceHealthStatusChangedToFailing:
		//
		s.AddGovernanceHistory(actualEvent.DeploymentID(), "ServiceAlert", "failing")
	case common.MonitoringServiceHealthStatusChangedToAtrisk:
		//
		s.AddGovernanceHistory(actualEvent.DeploymentID(), "ServiceAlert", "atrisk")
	default:
		return

	}
}

func (s *DeploymentService) AddGovernanceHistory(deploymentID string, eventType string, logMessage string) error {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	res, err := s.repo.GetByIDInternal(deploymentID)
	if err != nil {
		return err
	}

	err = res.AddEventHistory(eventType, logMessage, "system")
	if err != nil {
		return err
	}

	err = s.repo.Save(res)
	if err != nil {
		return err
	}

	return nil
}
