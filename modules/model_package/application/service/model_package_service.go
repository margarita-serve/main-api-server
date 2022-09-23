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
	"path/filepath"

	common "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/common"
	appDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/application/dto"
	domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/domain/entity"
	domRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/domain/repository"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rs/xid"

	//"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
	infStorageClient "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/storage/minio"
	infRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/infrastructure/repository"
	//appProjectDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/project/application/dto"
)

// ModelPackageService type

type StorageClient interface {
	UploadFile(ioReader interface{}, filePath string) error
	DeleteFile(filePath string) error
	GetFile(filePath string) (io.Reader, error)
}

type ModelPackageService struct {
	BaseService
	repo          domRepo.IModelPackageRepo
	storageClient StorageClient
	projectSvc    common.IProjectService
	publisher     common.EventPublisher
}

// type IProjectService interface {
// 	GetList(req *appProjectDTO.GetProjectListRequestDTO, i identity.Identity) (*appProjectDTO.GetProjectListResponseDTO, error)
// }

// NewModelPackageService new ModelPackageService
func NewModelPackageService(h *handler.Handler, projectSvc common.IProjectService) (*ModelPackageService, error) {
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

	svc.projectSvc = projectSvc

	return svc, nil
}

// Create
func (s *ModelPackageService) Create(req *appDTO.CreateModelPackageRequestDTO, i identity.Identity) (*appDTO.CreateModelPackageResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	// if err := req.Validate(); err != nil {
	// 	return nil, err
	// }

	//Project validation
	//projectReq := &appProjectDTO.GetProjectListRequestDTO{}
	projectRes, err := s.projectSvc.GetListInternal(i.Claims.Username)
	projectIdList := projectRes.Rows

	var listProjectId []string
	for _, rec := range projectIdList {
		listProjectId = append(listProjectId, rec.ProjectID)
	}

	var chkExist bool = false
	for _, oneOfProjectID := range listProjectId {
		if oneOfProjectID == req.ProjectID {
			chkExist = true
			break
		}
	}

	if !chkExist {
		return nil, errors.New("project not found")
	}
	//End Project validation

	guid := xid.New().String()

	//toBe...
	//OwnerID := "testID"

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
		i.Claims.Username,
	)
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
		return nil, errors.New("ModelPackage Has Deployed, Use Archive")
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

	if req.Name != nil {
		domAggregateModelPackage.SetName(*req.Name)
	}
	if req.Description != nil {
		domAggregateModelPackage.SetDescription(*req.Description)
	}
	if req.ModelDescription != nil {
		domAggregateModelPackage.SetModelDescription(*req.ModelDescription)
	}
	if req.ModelFrameWork != nil {
		domAggregateModelPackage.SetModelFrameWork(*req.ModelFrameWork)
	}
	if req.ModelFrameWorkVersion != nil {
		domAggregateModelPackage.SetModelFrameWorkVersion(*req.ModelFrameWorkVersion)
	}
	if req.ModelName != nil {
		domAggregateModelPackage.SetModelName(*req.ModelName)
	}
	if req.ModelVersion != nil {
		domAggregateModelPackage.SetModelVersion(*req.ModelVersion)
	}
	if req.NegativeClassLabel != nil {
		domAggregateModelPackage.SetNegativeClassLabel(*req.NegativeClassLabel)
	}
	if req.PositiveClassLabel != nil {
		domAggregateModelPackage.SetPositiveClassLabel(*req.PositiveClassLabel)
	}
	if req.PredictionTargetName != nil {
		domAggregateModelPackage.SetPredictionTargetName(*req.PredictionTargetName)
	}
	if req.PredictionThreshold != nil {
		domAggregateModelPackage.SetPredictionThreshold(*req.PredictionThreshold)
	}
	if req.TargetType != nil {
		domAggregateModelPackage.SetTargetType(*req.TargetType)
	}

	err = domEntity.Validate(domAggregateModelPackage)
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
	resDTO.ModelFileName = path.Base(res.ModelFilePath)
	resDTO.TrainingDatasetName = path.Base(res.TrainingDatasetPath)
	resDTO.HoldoutDatasetName = path.Base(res.HoldoutDatasetPath)

	return resDTO, nil
}

func (s *ModelPackageService) GetList(req *appDTO.GetModelPackageListRequestDTO, i identity.Identity) (*appDTO.GetModelPackageListResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }
	//projectReq := &appProjectDTO.GetProjectListRequestDTO{}
	projectRes, err := s.projectSvc.GetListInternal(i.Claims.Username)
	projectIdList := projectRes.Rows

	var listProjectId []string
	for _, rec := range projectIdList {
		listProjectId = append(listProjectId, rec.ProjectID)
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

func (s *ModelPackageService) GetListByProject(req *appDTO.GetModelPackageListByProjectRequestDTO) (*appDTO.GetModelPackageListByProjectResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	resultList, err := s.repo.GetListByProject(req.ProjectID)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.GetModelPackageListByProjectResponseDTO)

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

	check := domAggregateModelPackage.IsValidForUploadFile()
	if !check {
		return nil, errors.New("model package has deployed")
	}

	cfg, err := s.handler.GetConfig()
	if err != nil {
		return nil, err
	}

	//모델 프레임워크별 지원 확장자 검증
	err = validation.Validate(filepath.Ext(req.FileName), validation.When(domAggregateModelPackage.ModelFrameWork == "TensorFlow", validation.In(".zip", ".tar").Error("must be a extension in (.zip, .tar)")),
		validation.When(domAggregateModelPackage.ModelFrameWork == "SkLearn", validation.In(".pkl", ".joblib", ".pickle").Error("must be a extension in (.pkl, .joblib, .pickle)")),
		validation.When(domAggregateModelPackage.ModelFrameWork == "PyTorch", validation.In(".pt").Error("must be a extension in (.pt)")),
		validation.When(domAggregateModelPackage.ModelFrameWork == "XGBoost", validation.In(".bst").Error("must be a extension in (.bst)")),
		validation.When(domAggregateModelPackage.ModelFrameWork == "LightGBM", validation.In(".bst").Error("must be a extension in (.bst)")))
	if err != nil {
		return nil, err
	}

	//특정 모델 프레임워크의 경우 서빙모듈(kserve)에서 로드할때 model.pkl or model.extenstion 으로 고정해서 로드하기 때문에 파일명을 바꿔준다
	if domAggregateModelPackage.ModelFrameWork == "SkLearn" || domAggregateModelPackage.ModelFrameWork == "XGBoost" || domAggregateModelPackage.ModelFrameWork == "LightGBM" {
		req.FileName = "model" + filepath.Ext(req.FileName)
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

	check := domAggregateModelPackage.IsValidForUploadFile()
	if !check {
		return nil, errors.New("model package has deployed")
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

	check := domAggregateModelPackage.IsValidForUploadFile()
	if !check {
		return nil, errors.New("model package has deployed")
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

func (s *ModelPackageService) GetByIDInternal(modelPackageID string) (*common.InternalGetModelPackageResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	res, err := s.repo.GetByID(modelPackageID)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(common.InternalGetModelPackageResponseDTO)
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

func (s *ModelPackageService) GetModelFile(modelPackageID string) (io.Reader, string, error) {
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

func (s *ModelPackageService) GetTrainingDatasetFile(modelPackageID string) (io.Reader, string, error) {
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

func (s *ModelPackageService) GetHoldoutDatasetFile(modelPackageID string) (io.Reader, string, error) {
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

func (s *ModelPackageService) AddDeployCount(modelPackageID string) error {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }
	domAggregateModelPackage, err := s.repo.GetForUpdate(modelPackageID)
	if err != nil {
		return err
	}

	domAggregateModelPackage.AddDeployCount()

	err = s.repo.Save(domAggregateModelPackage)
	if err != nil {
		return err
	}

	return nil
}

func (s *ModelPackageService) Update(event common.Event) {
	switch actualEvent := event.(type) {
	case common.DeploymentCreated:
		//
		s.AddDeployCount(actualEvent.ModelPackageID())
	case common.DeploymentModelReplaced:
		//
		s.AddDeployCount(actualEvent.ModelPackageID())
	default:
		return

	}
}
