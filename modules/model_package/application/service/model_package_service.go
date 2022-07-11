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
	"context"
	"errors"
	"io"
	"path"

	appDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/application/dto"
	domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/domain/entity"
	domRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/domain/repository"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"github.com/minio/minio-go/v7"
	"github.com/rs/xid"

	//"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
	infStorageClient "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/storage/minio"
	infRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/infrastructure/repository"
)

// ModelPackageService type

type StorageClient interface {
	UploadFile(ioReader io.Reader, filePath string) error
	DeleteFile(filePath string) error
	GetFile(filePath string) (*minio.Object, error)
}

type ModelPackageService struct {
	BaseService
	repo          domRepo.IModelPackageRepo
	storageClient StorageClient
}

// NewModelPackageService new ModelPackageService
func NewModelPackageService(h *handler.Handler) (*ModelPackageService, error) {
	var err error

	svc := new(ModelPackageService)
	svc.handler = h
	// if err := svc.initBaseService(); err != nil {
	// 	return nil, err
	// }

	if svc.repo, err = infRepo.NewModelPackageRepo(h); err != nil {
		return nil, err
	}

	cfg, err := h.GetConfig()
	if err != nil {
		return nil, err
	}

	config := infStorageClient.Config{
		Endpoint:        cfg.Connectors.Storages.Minio.Endpoint,
		AccessKeyID:     cfg.Connectors.Storages.Minio.AccessKeyID,
		SecretAccessKey: cfg.Connectors.Storages.Minio.SecretAccessKey,
		UseSSL:          cfg.Connectors.Storages.Minio.UseSSL,
	}

	if svc.storageClient, err = infStorageClient.NewStorageClient(config, h, context.Background()); err != nil {
		return nil, err
	}

	return svc, nil
}

// Create
func (s *ModelPackageService) Create(req *appDTO.CreateModelPackageRequestDTO) (*appDTO.CreateModelPackageResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	// if err := req.Validate(); err != nil {
	// 	return nil, err
	// }

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

	domEntity.Validate(domAggregateModelPackage)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(domAggregateModelPackage)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.CreateModelPackageResponseDTO)
	resDTO.ModelPackageID = domAggregateModelPackage.ID

	return resDTO, nil
}

func (s *ModelPackageService) Delete(req *appDTO.DeleteModelPackageRequestDTO) (*appDTO.DeleteModelPackageResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }
	domAggregateModelPackage, err := s.repo.GetByID(req.ModelPackageID)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.DeleteModelPackageResponseDTO)

	check := domAggregateModelPackage.IsValidForDelete()
	if !check {
		return nil, errors.New("ModelPackage Has Deployed")
	}

	err = s.repo.Delete(req.ModelPackageID)
	if err != nil {
		return nil, err
	}

	resDTO.Message = "ModelPackage Delete Success"

	return resDTO, nil
}

func (s *ModelPackageService) Archive(req *appDTO.ArchiveModelPackageRequestDTO) (*appDTO.ArchiveModelPackageResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }
	domAggregateModelPackage, err := s.repo.GetByID(req.ModelPackageID)
	if err != nil {
		return nil, err
	}

	domAggregateModelPackage.SetArchived()

	err = s.repo.Save(domAggregateModelPackage)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.ArchiveModelPackageResponseDTO)
	resDTO.Message = "ModelPackage Archiving Success"

	return resDTO, nil
}

func (s *ModelPackageService) UpdateModelPackage(req *appDTO.UpdateModelPackageRequestDTO) (*appDTO.UpdateModelPackageResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	//toBe...
	//userID := "testID"

	//Find Domain Entity
	domAggregateModelPackage, err := s.repo.GetForUpdate(req.ModelPackageID)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.UpdateModelPackageResponseDTO)

	check := domAggregateModelPackage.IsValidForUpdate()
	if !check {
		return nil, errors.New("ModelPackage Has Archived")
	}

	if req.Name != "" {
		domAggregateModelPackage.SetName(req.Name)
	}
	if req.Description != "" {
		domAggregateModelPackage.SetDescription(req.Description)
	}
	if req.ModelDescription != "" {
		domAggregateModelPackage.SetModelDescription(req.ModelDescription)
	}
	if req.ModelFrameWork != "" {
		domAggregateModelPackage.SetModelFrameWork(req.ModelFrameWork)
	}
	if req.ModelFrameWorkVersion != "" {
		domAggregateModelPackage.SetModelFrameWorkVersion(req.ModelFrameWorkVersion)
	}
	if req.ModelName != "" {
		domAggregateModelPackage.SetModelName(req.ModelName)
	}
	if req.ModelVersion != "" {
		domAggregateModelPackage.SetModelVersion(req.ModelVersion)
	}
	if req.NegativeClassLabel != "" {
		domAggregateModelPackage.SetNegativeClassLabel(req.NegativeClassLabel)
	}
	if req.PositiveClassLabel != "" {
		domAggregateModelPackage.SetPositiveClassLabel(req.PositiveClassLabel)
	}
	if req.PredictionTargetName != "" {
		domAggregateModelPackage.SetPredictionTargetName(req.PredictionTargetName)
	}
	if req.PredictionThreshold != 0 {
		domAggregateModelPackage.SetPredictionThreshold(req.PredictionThreshold)
	}
	if req.TargetType != "" {
		domAggregateModelPackage.SetTargetType(req.TargetType)
	}

	domEntity.Validate(domAggregateModelPackage)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(domAggregateModelPackage)
	if err != nil {
		return nil, err
	}

	resDTO.Message = "ModelPackage Update Success"

	return resDTO, nil
}

func (s *ModelPackageService) GetByID(req *appDTO.GetModelPackageRequestDTO) (*appDTO.GetModelPackageResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	res, err := s.repo.GetByID(req.ModelPackageID)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.GetModelPackageResponseDTO)
	resDTO.ModelPackageID = res.ID
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

func (s *ModelPackageService) GetList(req *appDTO.GetModelPackageListRequestDTO) (*appDTO.GetModelPackageListResponseDTO, error) {
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
	resDTO := new(appDTO.GetModelPackageListResponseDTO)
	resDTO.Limit = p.Limit
	resDTO.Page = p.Page
	resDTO.TotalRows = p.TotalRows
	resDTO.TotalPages = p.TotalPages

	var listModelPackage []*appDTO.GetModelPackageResponseDTO
	for _, rec := range resultList {
		tmp := new(appDTO.GetModelPackageResponseDTO)

		tmp.ModelPackageID = rec.ID
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
		tmp.ModelFileName = path.Base(rec.ModelFilePath)
		tmp.TrainingDatasetName = path.Base(rec.TrainingDatasetPath)
		tmp.HoldoutDatasetName = path.Base(rec.HoldoutDatasetPath)
		tmp.Archived = rec.Archived

		listModelPackage = append(listModelPackage, tmp)
	}

	resDTO.Rows = listModelPackage

	return resDTO, nil
}

func (s *ModelPackageService) UploadModel(req *appDTO.UploadModelRequestDTO) (*appDTO.UploadModelResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	//Find ModelPackage
	domAggregateModelPackage, err := s.repo.GetByID(req.ModelPackageID)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.UploadModelResponseDTO)

	check := domAggregateModelPackage.IsValidForUpdate()
	if !check {
		return nil, errors.New("ModelPackage Has Archived")
	}

	cfg, err := s.handler.GetConfig()
	if err != nil {
		return nil, err
	}

	if domAggregateModelPackage.ModelFilePath != "" {
		deleteFilePath := domAggregateModelPackage.ModelFilePath
		err = s.storageClient.DeleteFile(deleteFilePath)
		if err != nil {
			return nil, err
		}
	}

	uploadFilePath := cfg.DirLocations.ModelPackageFileRootPath + "/" + domAggregateModelPackage.ID + cfg.DirLocations.ModelPath + "/" + req.FileName

	err = s.storageClient.UploadFile(req.File, uploadFilePath)
	if err != nil {
		return nil, err
	}

	domAggregateModelPackage.SetModelPath(uploadFilePath)

	err = s.repo.Save(domAggregateModelPackage)
	if err != nil {
		return nil, err
	}

	resDTO.Message = "Upload Success"

	return resDTO, nil
}

func (s *ModelPackageService) UploadTrainingDataset(req *appDTO.UploadTrainingDatasetRequestDTO) (*appDTO.UploadTrainingDatasetResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	//Find ModelPackage
	domAggregateModelPackage, err := s.repo.GetByID(req.ModelPackageID)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.UploadTrainingDatasetResponseDTO)

	check := domAggregateModelPackage.IsValidForUpdate()
	if !check {
		return nil, errors.New("ModelPackage Has Archived")
	}

	cfg, err := s.handler.GetConfig()
	if err != nil {
		return nil, err
	}

	if domAggregateModelPackage.TrainingDatasetPath != "" {
		deleteFilePath := domAggregateModelPackage.TrainingDatasetPath
		err = s.storageClient.DeleteFile(deleteFilePath)
		if err != nil {
			return nil, err
		}
	}

	uploadFilePath := cfg.DirLocations.ModelPackageFileRootPath + "/" + domAggregateModelPackage.ID + cfg.DirLocations.TrainingDatasetPath + "/" + req.FileName

	err = s.storageClient.UploadFile(req.File, uploadFilePath)
	if err != nil {
		return nil, err
	}

	domAggregateModelPackage.SetTraningDatasetPath(uploadFilePath)

	err = s.repo.Save(domAggregateModelPackage)
	if err != nil {
		return nil, err
	}

	resDTO.Message = "Upload Success"

	return resDTO, nil
}

func (s *ModelPackageService) UploadHoldoutDataset(req *appDTO.UploadHoldoutDatasetRequestDTO) (*appDTO.UploadHoldoutDatasetResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	//Find ModelPackage
	domAggregateModelPackage, err := s.repo.GetByID(req.ModelPackageID)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.UploadHoldoutDatasetResponseDTO)

	check := domAggregateModelPackage.IsValidForUpdate()
	if !check {
		return nil, errors.New("ModelPackage Has Archived")
	}

	cfg, err := s.handler.GetConfig()
	if err != nil {
		return nil, err
	}

	if domAggregateModelPackage.HoldoutDatasetPath != "" {
		deleteFilePath := domAggregateModelPackage.HoldoutDatasetPath
		err = s.storageClient.DeleteFile(deleteFilePath)
		if err != nil {
			return nil, err
		}
	}

	uploadFilePath := cfg.DirLocations.ModelPackageFileRootPath + "/" + domAggregateModelPackage.ID + cfg.DirLocations.HoldoutDatasetPath + "/" + req.FileName

	err = s.storageClient.UploadFile(req.File, uploadFilePath)
	if err != nil {
		return nil, err
	}

	domAggregateModelPackage.SetHoldoutDatasetPath(uploadFilePath)

	err = s.repo.Save(domAggregateModelPackage)
	if err != nil {
		return nil, err
	}

	resDTO.Message = "Upload Success"

	return resDTO, nil
}

func (s *ModelPackageService) GetByIDInternal(req *appDTO.InternalGetModelPackageRequestDTO) (*appDTO.InternalGetModelPackageResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	res, err := s.repo.GetByID(req.ModelPackageID)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.InternalGetModelPackageResponseDTO)
	resDTO.ModelPackageID = res.ID
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
	resDTO.Archived = res.Archived

	return resDTO, nil
}

func (s *ModelPackageService) GetModelFile(modelPackageID string) (*minio.Object, string, error) {
	res, err := s.repo.GetByID(modelPackageID)
	if err != nil {
		return nil, "", err
	}

	fileReader, err := s.storageClient.GetFile(res.ModelFilePath)
	if err != nil {
		return nil, "", err
	}

	fileName := path.Base(res.ModelFilePath)

	return fileReader, fileName, nil
}

func (s *ModelPackageService) GetTrainingDatasetFile(modelPackageID string) (*minio.Object, string, error) {
	res, err := s.repo.GetByID(modelPackageID)
	if err != nil {
		return nil, "", err
	}

	fileReader, err := s.storageClient.GetFile(res.TrainingDatasetPath)
	if err != nil {
		return nil, "", err
	}

	fileName := path.Base(res.TrainingDatasetPath)

	return fileReader, fileName, nil
}

func (s *ModelPackageService) GetHoldoutDatasetFile(modelPackageID string) (*minio.Object, string, error) {
	res, err := s.repo.GetByID(modelPackageID)
	if err != nil {
		return nil, "", err
	}

	fileReader, err := s.storageClient.GetFile(res.HoldoutDatasetPath)
	if err != nil {
		return nil, "", err
	}

	fileName := path.Base(res.HoldoutDatasetPath)

	return fileReader, fileName, nil
}

// func (s *ModelPackageService) GetTrainingDatasetFileStorageURL(modelPackageID string) (string, error) {
// 	res, err := s.repo.GetByID(modelPackageID)
// 	if err != nil {
// 		return "", err
// 	}

// 	return s.storageClient.GetEndPoint() + "/" + res.TrainingDatasetPath, err
// }

// func (s *ModelPackageService) GetHoldoutDatasetFileStorageURL(modelPackageID string) (string, error) {
// 	res, err := s.repo.GetByID(modelPackageID)
// 	if err != nil {
// 		return "", err
// 	}

// 	return s.storageClient.GetEndPoint() + "/" + res.HoldoutDatasetPath, err
// }
