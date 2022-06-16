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
	appDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/application/dto"
	domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/domain/entity"
	domRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/domain/repository"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"github.com/rs/xid"

	//"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
	infRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/infrastructure/repository"
)

// ModelPackageService type
type ModelPackageService struct {
	BaseService
	repo domRepo.IModelPackageRepo
}

// NewModelPackageService new ModelPackageService
func NewModelPackageService(h *handler.Handler) (*ModelPackageService, error) {
	var err error

	svc := new(ModelPackageService)
	//svc.handler = h
	//if err := svc.initBaseService(); err != nil {
	//	return nil, err
	//}

	if svc.repo, err = infRepo.NewModelPackageRepo(h); err != nil {
		return nil, err
	}

	return svc, nil
}

// Create
func (s *ModelPackageService) Create(req *appDTO.ModelPackageCreateRequestDTO) (*appDTO.ModelPackageCreateResponseDTO, error) {
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
	domAggregateModelPackage, err := domEntity.NewModelPackage(
		guid,
		req.ProjectID,
		req.Name,
		req.Description,
		req.ModelName,
		req.ModelVersion,
		req.ModelDescription,
		req.TargetType,
		req.PredictionTargetName,
		req.ModelFrameWork,
		req.ModelFrameWorkVersion,
		req.PredictionThreshold,
		req.PositiveClassLabel,
		req.NegativeClassLabel,
		OwnerID,
	)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(domAggregateModelPackage)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.ModelPackageCreateResponseDTO)
	resDTO.ModelPackageID = domAggregateModelPackage.ID

	return resDTO, nil
}

func (s *ModelPackageService) Delete(req *appDTO.ModelPackageDeleteRequestDTO) (*appDTO.ModelPackageDeleteResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	err := s.repo.Delete(req.ModelPackageID)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.ModelPackageDeleteResponseDTO)
	resDTO.Message = "ModelPackage Delete Success"

	return resDTO, nil
}

func (s *ModelPackageService) GetByID(req *appDTO.ModelPackageGetRequestDTO) (*appDTO.ModelPackageGetResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	res, err := s.repo.Get(req.ModelPackageID)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.ModelPackageGetResponseDTO)
	resDTO.ID = res.ID
	resDTO.ProjectID = res.ProjectID
	resDTO.Name = res.Name
	resDTO.Description = res.Description
	resDTO.ModelName = res.ModelName
	resDTO.ModelVersion = res.ModelVersion
	resDTO.ModelDescription = res.ModelDescription
	resDTO.TargetType = res.TargetType
	resDTO.PredictionTargetName = res.PredictionTargetName
	resDTO.ModelFrameWork = res.ModelFrameWork
	resDTO.ModelFrameWorkVersion = res.ModelFrameWorkVersion
	resDTO.PredictionThreshold = res.PredictionThreshold
	resDTO.PositiveClassLabel = res.PositiveClassLabel
	resDTO.NegativeClassLabel = res.NegativeClassLabel

	return resDTO, nil
}

func (s *ModelPackageService) GetByIDInternal(req *appDTO.ModelPackageGetInternalRequestDTO) (*appDTO.ModelPackageGetInternalResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	res, err := s.repo.Get(req.ModelPackageID)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.ModelPackageGetInternalResponseDTO)
	resDTO.ID = res.ID
	resDTO.ProjectID = res.ProjectID
	resDTO.Name = res.Name
	resDTO.Description = res.Description
	resDTO.ModelName = res.ModelName
	resDTO.ModelVersion = res.ModelVersion
	resDTO.ModelDescription = res.ModelDescription
	resDTO.TargetType = res.TargetType
	resDTO.PredictionTargetName = res.PredictionTargetName
	resDTO.ModelFrameWork = res.ModelFrameWork
	resDTO.ModelFrameWorkVersion = res.ModelFrameWorkVersion
	resDTO.PredictionThreshold = res.PredictionThreshold
	resDTO.PositiveClassLabel = res.PositiveClassLabel
	resDTO.NegativeClassLabel = res.NegativeClassLabel
	resDTO.ModelFilePath = res.ModelFilePath
	resDTO.TrainingDatasetPath = res.TrainingDatasetPath
	resDTO.HoldoutDatasetPath = res.HoldoutDatasetPath

	return resDTO, nil
}

func (s *ModelPackageService) GetByName(req *appDTO.ModelPackageGetByNametRequestDTO) (*appDTO.ModelPackageGetByNameResponseDTO, error) {
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

	var listModelPackage []*appDTO.ModelPackage
	for _, rec := range res {
		tmp := new(appDTO.ModelPackage)

		tmp.ID = rec.ID
		tmp.ProjectID = rec.ProjectID
		tmp.Name = rec.Name
		tmp.Description = rec.Description
		tmp.ModelName = rec.ModelName
		tmp.ModelVersion = rec.ModelVersion
		tmp.ModelDescription = rec.ModelDescription
		tmp.TargetType = rec.TargetType
		tmp.PredictionTargetName = rec.PredictionTargetName
		tmp.ModelFrameWork = rec.ModelFrameWork
		tmp.ModelFrameWorkVersion = rec.ModelFrameWorkVersion
		tmp.PredictionThreshold = rec.PredictionThreshold
		tmp.PositiveClassLabel = rec.PositiveClassLabel
		tmp.NegativeClassLabel = rec.NegativeClassLabel

		listModelPackage = append(listModelPackage, tmp)
	}

	resDTO := new(appDTO.ModelPackageGetByNameResponseDTO)
	resDTO.ModelPackages = listModelPackage

	return resDTO, nil
}

func (s *ModelPackageService) GetAll(req *appDTO.ModelPackageGetByNametRequestDTO) (*appDTO.ModelPackageGetByNameResponseDTO, error) {
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

	var listModelPackage []*appDTO.ModelPackage
	for _, rec := range res {
		tmp := new(appDTO.ModelPackage)

		tmp.ID = rec.ID
		tmp.ProjectID = rec.ProjectID
		tmp.Name = rec.Name
		tmp.Description = rec.Description
		tmp.ModelName = rec.ModelName
		tmp.ModelVersion = rec.ModelVersion
		tmp.ModelDescription = rec.ModelDescription
		tmp.TargetType = rec.TargetType
		tmp.PredictionTargetName = rec.PredictionTargetName
		tmp.ModelFrameWork = rec.ModelFrameWork
		tmp.ModelFrameWorkVersion = rec.ModelFrameWorkVersion
		tmp.PredictionThreshold = rec.PredictionThreshold
		tmp.PositiveClassLabel = rec.PositiveClassLabel
		tmp.NegativeClassLabel = rec.NegativeClassLabel

		listModelPackage = append(listModelPackage, tmp)
	}

	resDTO := new(appDTO.ModelPackageGetByNameResponseDTO)
	resDTO.ModelPackages = listModelPackage

	return resDTO, nil
}
