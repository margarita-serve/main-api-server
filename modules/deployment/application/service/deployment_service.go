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
)

// DeploymentService type
type DeploymentService struct {
	BaseService
	domInferenceSvc     domSvcInferenceSvc.IInferenceServiceAdapter
	modelPackageService *appModelPackageSvc.ModelPackageService
	// predictionEnvService   *appPredictionEnvSvc.PredictionEnvService
	// monitoringSvc          *appMonitoringSvc.MonitoringService
	repo domRepo.IDeploymentRepo
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

	if svc.modelPackageService, err = appModelPackageSvc.NewModelPackageService(h); err != nil {
		return nil, err
	}

	return svc, nil
}

// Create
func (s *DeploymentService) Create(req *appDTO.DeploymentCreateRequestDTO) (*appDTO.DeploymentCreateResponseDTO, error) {
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
	OwnerID := "testID"

	// New deployment domain Instance
	domAggregateDeployment, err := domEntity.NewDeployment(
		guid,
		req.ProjectID,
		req.ModelPackageID,
		req.PredictionEnvID,
		req.Name,
		req.Description,
		req.Importance,
		req.RequestCPU,
		req.RequestMEM,
		req.RequestGPU,
		OwnerID,
	)
	if err != nil {
		return nil, err
	}

	//search  ModelPackage
	reqModelPackage := &appModelPackageDTO.ModelPackageGetInternalRequestDTO{
		ModelPackageID: req.ModelPackageID,
	}

	resModelPackage, err := s.modelPackageService.GetByIDInternal(reqModelPackage)
	if err != nil {
		return nil, err
	}
	resModelPackage.ModelFilePath = "s3://testmodel/mpg2"

	//search  PredictionEnv
	predictionEnvInfo := domSchema.PredictionEnv{
		Namespace:      "koreserve",
		ConnectionInfo: "http://192.168.88.161:30070",
	}

	reqDomSvc := domSvcInferenceSvcDto.InferenceServiceCreateRequest{
		Namespace:      predictionEnvInfo.Namespace,
		Inferencename:  guid,
		ModelFrameWork: resModelPackage.ModelFrameWork,
		ModelURL:       resModelPackage.ModelFilePath,
		ModelName:      resModelPackage.ModelName,
		ConnectionInfo: predictionEnvInfo.ConnectionInfo,
	}

	if err := reqDomSvc.Validate(); err != nil {
		return nil, err
	}

	err = s.repo.Save(domAggregateDeployment)
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
	resDTO := new(appDTO.DeploymentCreateResponseDTO)
	resDTO.DeploymentID = domAggregateDeployment.ID

	return resDTO, nil
}

func (s *DeploymentService) Delete(req *appDTO.DeploymentDeleteRequestDTO) (*appDTO.DeploymentDeleteResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	domAggregateDeployment, err := s.repo.Get(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	//search  PredictionEnv
	predictionEnvInfo := domSchema.PredictionEnv{
		Namespace:      "koreserve",
		ConnectionInfo: "http://192.168.88.161:30070",
	}

	reqDomSvc := domSvcInferenceSvcDto.InferenceServiceDeleteRequest{
		ConnectionInfo: predictionEnvInfo.ConnectionInfo,
		Namespace:      predictionEnvInfo.Namespace,
		Inferencename:  domAggregateDeployment.ID,
	}

	err = domAggregateDeployment.RequestDeleteInferenceService(s.domInferenceSvc, reqDomSvc)
	if err != nil {
		return nil, err
	}

	err = s.repo.Delete(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.DeploymentDeleteResponseDTO)
	resDTO.Message = "Deployment Delete Success"

	return resDTO, nil
}

func (s *DeploymentService) GetByID(req *appDTO.DeploymentGetRequestDTO) (*appDTO.DeploymentGetResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	res, err := s.repo.Get(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.DeploymentGetResponseDTO)
	resDTO.ID = res.ID
	resDTO.ProjectID = res.ProjectID
	resDTO.ModelPackageID = res.ModelPackageID
	resDTO.PredictionEnvID = res.PredictionEnvID
	resDTO.Name = res.Name
	resDTO.Description = res.Description
	resDTO.Importance = res.Importance
	resDTO.RequestCPU = res.RequestCPU
	resDTO.RequestMEM = res.RequestMEM
	resDTO.RequestGPU = res.RequestGPU
	resDTO.ActiveStatus = res.ActiveStatus
	resDTO.ServiceStatus = res.ServiceStatus
	resDTO.ChangeRequested = res.ChangeRequested

	return resDTO, nil
}

func (s *DeploymentService) GetByName(req *appDTO.DeploymentGetByNametRequestDTO) (*appDTO.DeploymentGetByNameResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	res, err := s.repo.ByName(req.Name)
	if err != nil {
		return nil, err
	}

	// response dto

	var listDeployment []*appDTO.Deployment
	for _, rec := range res {
		tmp := new(appDTO.Deployment)

		tmp.ID = rec.ID
		tmp.ProjectID = rec.ProjectID
		tmp.ModelPackageID = rec.ModelPackageID
		tmp.PredictionEnvID = rec.PredictionEnvID
		tmp.Name = rec.Name
		tmp.Description = rec.Description
		tmp.Importance = rec.Importance
		tmp.RequestCPU = rec.RequestCPU
		tmp.RequestMEM = rec.RequestMEM
		tmp.RequestGPU = rec.RequestGPU
		tmp.ActiveStatus = rec.ActiveStatus
		tmp.ServiceStatus = rec.ServiceStatus
		tmp.ChangeRequested = rec.ChangeRequested

		listDeployment = append(listDeployment, tmp)
	}

	resDTO := new(appDTO.DeploymentGetByNameResponseDTO)
	resDTO.Deployments = listDeployment

	return resDTO, nil
}

func (s *DeploymentService) GetAll(req *appDTO.DeploymentGetByNametRequestDTO) (*appDTO.DeploymentGetByNameResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }
	print(req.Name)
	res, err := s.repo.ByName(req.Name)
	if err != nil {
		return nil, err
	}

	// response dto

	var listDeployment []*appDTO.Deployment
	for _, rec := range res {
		tmp := new(appDTO.Deployment)

		tmp.ID = rec.ID
		tmp.ProjectID = rec.ProjectID
		tmp.ModelPackageID = rec.ModelPackageID
		tmp.PredictionEnvID = rec.PredictionEnvID
		tmp.Name = rec.Name
		tmp.Description = rec.Description
		tmp.Importance = rec.Importance
		tmp.RequestCPU = rec.RequestCPU
		tmp.RequestMEM = rec.RequestMEM
		tmp.RequestGPU = rec.RequestGPU
		tmp.ActiveStatus = rec.ActiveStatus
		tmp.ServiceStatus = rec.ServiceStatus
		tmp.ChangeRequested = rec.ChangeRequested

		listDeployment = append(listDeployment, tmp)
	}

	resDTO := new(appDTO.DeploymentGetByNameResponseDTO)
	resDTO.Deployments = listDeployment

	return resDTO, nil
}

func (s *DeploymentService) SetActive(req *appDTO.DeploymentActiveRequestDTO) (*appDTO.DeploymentActiveResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	domAggregateDeployment, err := s.repo.Get(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	//search  ModelPackage
	modelPackageInfo := domSchema.ModelPackage{
		ModelURL:       "s3://testmodel/mpg2",
		ModelFrameWork: "tensorflow",
	}
	//search  PredictionEnv
	predictionEnvInfo := domSchema.PredictionEnv{
		Namespace:      "koreserve",
		ConnectionInfo: "http://192.168.88.161:30070",
	}

	reqDom := domSvcInferenceSvcDto.InferenceServiceCreateRequest{
		Namespace:      predictionEnvInfo.Namespace,
		Inferencename:  domAggregateDeployment.ID,
		ModelFrameWork: modelPackageInfo.ModelFrameWork,
		ModelURL:       modelPackageInfo.ModelURL,
		ConnectionInfo: predictionEnvInfo.ConnectionInfo,
	}

	// if err := reqDom.Validate(); err != nil {
	// 	return nil, err
	// }

	err = domAggregateDeployment.SetDeploymentActive(s.domInferenceSvc, reqDom)
	if err != nil {
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
	resDTO := new(appDTO.DeploymentActiveResponseDTO)
	resDTO.Message = "Deployment Delete Success"

	return resDTO, nil
}

// // Active
// func (s *DeploymentService) Active(req *appDTO.DeploymentActiveRequestDTO) (*appDTO.DeploymentActiveResponseDTO, error) {
// 	// //authorization
// 	// if i.CanAccessCurrentRequest() == false {
// 	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
// 	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
// 	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
// 	// }

// 	if err := req.Validate(); err != nil {
// 		return nil, err
// 	}

// 	res, err := s.repo.Get(req.DeploymentID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// response dto
// 	resDTO := new(appDTO.DeploymentActiveResponseDTO)
// 	//resDTO.Message = res.Message
// 	resDTO.Message = res.Message

// 	return resDTO, nil
// }
