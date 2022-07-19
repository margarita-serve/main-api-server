package service

import (
	"context"
	"io"

	infStorageClient "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/storage/minio"

	// "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/application/dto"
	appModelPackageDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/application/dto"
	// appModelPackageSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/application/service"
	appDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring_mockup/application/dto"
	domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring_mockup/domain/entity"
	domRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring_mockup/domain/repository"
	domSvcMonitorSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring_mockup/domain/service"
	domSvcMonitorSvcAccuracyDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring_mockup/domain/service/accuracy/dto"
	domSvcMonitorSvcDriftDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring_mockup/domain/service/data_drift/dto"
	infAccuracySvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring_mockup/infrastructure/monitor_service/accuracy"
	infDriftSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring_mockup/infrastructure/monitor_service/data_drift"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"

	//domSvcMonitorSvcAccuracyDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring_mockup/domain/service"
	infRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring_mockup/infrastructure/repository"
	"github.com/minio/minio-go/v7"
)

/*

application service 영역.
외부와 domain을 이어줌.
domain을 호출하여 사용.
들어가야할 내용은 단순 실행뿐.

*/

type IModelPackageService interface {
	GetByIDInternal(req *appModelPackageDTO.InternalGetModelPackageRequestDTO) (*appModelPackageDTO.InternalGetModelPackageResponseDTO, error)
}

type StorageClient interface {
	UploadFile(ioReader io.Reader, filePath string) error
	DeleteFile(filePath string) error
	GetFile(filePath string) (*minio.Object, error)
}
type MonitorService struct {
	BaseService
	domMonitorDriftSvc    domSvcMonitorSvc.IExternalDriftMonitorAdapter
	domMonitorAccuracySvc domSvcMonitorSvc.IExternalAccuracyMonitorAdapter
	modelPackageSvc       IModelPackageService
	repo                  domRepo.IMonitorRepo
	storageClient         StorageClient
}

func NewMonitorService(h *handler.Handler, modelPackageSvc IModelPackageService) (*MonitorService, error) {
	var err error

	svc := new(MonitorService)

	svc.handler = h
	// base service init
	if err := svc.initBaseService(); err != nil {
		return nil, err
	}
	// monitor repo 생성
	if svc.repo, err = infRepo.NewMonitorRepo(h); err != nil {
		return nil, err
	}

	if svc.domMonitorDriftSvc, err = infDriftSvc.NewDataDriftAdapter(); err != nil {
		return nil, err
	}

	if svc.domMonitorAccuracySvc, err = infAccuracySvc.NewAccuracyAdapter(); err != nil {
		return nil, err
	}

	svc.modelPackageSvc = modelPackageSvc

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

func (s *MonitorService) Create(req *appDTO.MonitorCreateRequestDTO) (*appDTO.MonitorCreateResponseDTO, error) {
	// validate check

	domAggregateMonitor, err := domEntity.NewMonitor(
		req.DeploymentID,
		req.ModelPackageID,
	)

	if err != nil {
		return nil, err
	}

	resModelPackage, err := s.getModelPackageByID(req.ModelPackageID)

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	if req.FeatureDriftTracking == true {
		domAggregateMonitor.SetDataDriftSetting(
			req.DataDriftSetting.MonitorRange,
			req.DataDriftSetting.DriftMetricType,
			req.DataDriftSetting.DriftThreshold,
			req.DataDriftSetting.ImportanceThreshold,
			req.DataDriftSetting.LowImportanceAtRiskCount,
			req.DataDriftSetting.LowImportanceFailingCount,
			req.DataDriftSetting.HighImportanceAtRiskCount,
			req.DataDriftSetting.HighImportanceFailingCount,
		)
		reqDomDriftSvc := domSvcMonitorSvcDriftDTO.DataDriftCreateRequest{
			InferenceName:              req.DeploymentID,
			ModelHistoryID:             req.ModelHistoryID,
			TargetLabel:                resModelPackage.PredictionTargetName,
			ModelType:                  resModelPackage.TargetType,
			Framework:                  resModelPackage.ModelFrameWork,
			TrainDatasetPath:           resModelPackage.TrainingDatasetPath,
			ModelPath:                  resModelPackage.ModelFilePath,
			DriftThreshold:             domAggregateMonitor.DataDriftSetting.DriftThreshold,
			ImportanceThreshold:        domAggregateMonitor.DataDriftSetting.ImportanceThreshold,
			MonitorRange:               domAggregateMonitor.DataDriftSetting.MonitorRange,
			LowImportanceAtRiskCount:   domAggregateMonitor.DataDriftSetting.LowImportanceAtRiskCount,
			LowImportanceFailingCount:  domAggregateMonitor.DataDriftSetting.LowImportanceFailingCount,
			HighImportanceAtRiskCount:  domAggregateMonitor.DataDriftSetting.HighImportanceAtRiskCount,
			HighImportanceFailingCount: domAggregateMonitor.DataDriftSetting.HighImportanceFailingCount,
		}

		err = domAggregateMonitor.SetFeatureDriftTrackingOn(s.domMonitorDriftSvc, reqDomDriftSvc)
		if err != nil {
			return nil, err
		}
		err = s.repo.Save(domAggregateMonitor)
		if err != nil {
			return nil, err
		}
	}
	if req.AccuracyMonitoring == true {
		domAggregateMonitor.SetAccuracySetting(
			req.AccuracySetting.MetricType,
			req.AccuracySetting.Measurement,
			req.AccuracySetting.AtRiskValue,
			req.AccuracySetting.FailingValue,
		)
		reqDomAccuracySvc := domSvcMonitorSvcAccuracyDTO.AccuracyCreateRequest{
			InferenceName:    req.DeploymentID,
			ModelHistoryID:   req.ModelHistoryID,
			DatasetPath:      resModelPackage.HoldoutDatasetPath,
			ModelPath:        resModelPackage.ModelFilePath,
			TargetLabel:      resModelPackage.PredictionTargetName,
			AssociationID:    req.AssociationID,
			ModelType:        resModelPackage.TargetType,
			Framework:        resModelPackage.ModelFrameWork,
			DriftMetrics:     domAggregateMonitor.AccuracySetting.MetricType,
			DriftMeasurement: domAggregateMonitor.AccuracySetting.Measurement,
			AtriskValue:      domAggregateMonitor.AccuracySetting.AtRiskValue,
			FailingValue:     domAggregateMonitor.AccuracySetting.FailingValue,
			PositiveClass:    resModelPackage.PositiveClassLabel,
			NegativeClass:    resModelPackage.NegativeClassLabel,
			BinaryThreshold:  resModelPackage.PredictionThreshold,
		}

		err = domAggregateMonitor.SetAccuracyMonitoringOn(s.domMonitorAccuracySvc, reqDomAccuracySvc)
		if err != nil {
			return nil, err
		}
		err = s.repo.Save(domAggregateMonitor)
		if err != nil {
			return nil, err
		}
	}
	resDTO := new(appDTO.MonitorCreateResponseDTO)
	resDTO.DeploymentID = domAggregateMonitor.ID

	return resDTO, nil
}

func (s *MonitorService) Delete(req *appDTO.MonitorDeleteRequestDTO) (*appDTO.MonitorDeleteResponseDTO, error) {
	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	// Drift OFF
	reqDomDriftSvc := domSvcMonitorSvcDriftDTO.DataDriftDeleteRequest{
		InferenceName: req.DeploymentID,
	}
	err = domAggregateMonitor.SetFeatureDriftTrackingOff(s.domMonitorDriftSvc, reqDomDriftSvc)
	if err != nil {
		return nil, err
	}

	// Accuracy OFF
	reqDomAccuracySvc := domSvcMonitorSvcAccuracyDTO.AccuracyDeleteRequest{
		InferenceName: req.DeploymentID,
	}
	err = domAggregateMonitor.SetAccuracyMonitoringOff(s.domMonitorAccuracySvc, reqDomAccuracySvc)
	if err != nil {
		return nil, err
	}

	err = s.repo.Delete(domAggregateMonitor.ID)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.MonitorDeleteResponseDTO)
	resDTO.Message = "DataDrift Monitor Inactive Success"

	return resDTO, nil
}

func (s *MonitorService) PatchDriftMonitorSetting(req *appDTO.MonitorDriftPatchRequestDTO) (*appDTO.MonitorDriftPatchResponseDTO, error) {
	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)
	if err != nil {
		return nil, err
	}
	resDTO := new(appDTO.MonitorDriftPatchResponseDTO)
	if domAggregateMonitor.FeatureDriftTracking == false {
		resDTO.DeploymentID = req.DeploymentID
		resDTO.Message = "DataDrift Tracking is False"
		return resDTO, nil
	}
	reqDomDriftSvc := domSvcMonitorSvcDriftDTO.DataDriftPatchRequest{
		InferenceName:              req.DeploymentID,
		DriftThreshold:             req.DataDriftSetting.DriftThreshold,
		ImportanceThreshold:        req.DataDriftSetting.ImportanceThreshold,
		MonitorRange:               req.DataDriftSetting.MonitorRange,
		LowImportanceAtRiskCount:   req.DataDriftSetting.LowImportanceAtRiskCount,
		LowImportanceFailingCount:  req.DataDriftSetting.LowImportanceFailingCount,
		HighImportanceAtRiskCount:  req.DataDriftSetting.HighImportanceAtRiskCount,
		HighImportanceFailingCount: req.DataDriftSetting.HighImportanceFailingCount,
	}

	err = domAggregateMonitor.PatchDataDriftSetting(s.domMonitorDriftSvc, reqDomDriftSvc)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(domAggregateMonitor)
	if err != nil {
		return nil, err
	}

	resDTO.DeploymentID = req.DeploymentID
	resDTO.Message = "DataDrift Monitor Patch Success"

	return resDTO, nil
}

func (s *MonitorService) SetDriftMonitorActive(req *appDTO.MonitorDriftActiveRequestDTO) (*appDTO.MonitorDriftActiveResponseDTO, error) {
	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)

	// modelHistoryDTO := new(appDeploymentDTO.GetModelHistoryRequestDTO)
	// modelHistoryDTO.DeploymentID = domAggregateMonitor.ID

	// modelHistoryRes, err := s.deploymentSvc.GetModelHistory(modelHistoryDTO)
	// modelHistory := modelHistoryRes.ModelHistory.([]entity.ModelHistory)

	// var modelHistoryID string
	// for _, v := range modelHistory {
	// 	if v.ApplyHistoryTag == "current" {
	// 		modelHistoryID = v.Version
	// 	}
	// }

	resModelPackage, err := s.getModelPackageByID(req.ModelPackageID)
	if err != nil {
		return nil, err
	}

	reqDomDriftSvc := domSvcMonitorSvcDriftDTO.DataDriftCreateRequest{
		InferenceName:              req.DeploymentID,
		ModelHistoryID:             req.CurrentModelID,
		TargetLabel:                resModelPackage.PredictionTargetName,
		ModelType:                  resModelPackage.TargetType,
		Framework:                  resModelPackage.ModelFrameWork,
		DriftThreshold:             req.DataDriftSetting.DriftThreshold,
		ImportanceThreshold:        req.DataDriftSetting.ImportanceThreshold,
		MonitorRange:               req.DataDriftSetting.MonitorRange,
		LowImportanceAtRiskCount:   req.DataDriftSetting.LowImportanceAtRiskCount,
		LowImportanceFailingCount:  req.DataDriftSetting.LowImportanceFailingCount,
		HighImportanceAtRiskCount:  req.DataDriftSetting.HighImportanceAtRiskCount,
		HighImportanceFailingCount: req.DataDriftSetting.HighImportanceFailingCount,
	}

	err = domAggregateMonitor.SetFeatureDriftTrackingOn(s.domMonitorDriftSvc, reqDomDriftSvc)
	if err != nil {
		return nil, err
	}
	err = s.repo.Save(domAggregateMonitor)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.MonitorDriftActiveResponseDTO)
	resDTO.DeploymentID = domAggregateMonitor.ID

	return resDTO, nil
}

func (s *MonitorService) SetDriftMonitorInActive(req *appDTO.MonitorDriftInActiveRequestDTO) (*appDTO.MonitorDriftInActiveResponseDTO, error) {
	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)
	if err != nil {
		return nil, err
	}
	reqDomDriftSvc := domSvcMonitorSvcDriftDTO.DataDriftDeleteRequest{
		InferenceName: req.DeploymentID,
	}
	err = domAggregateMonitor.SetFeatureDriftTrackingOff(s.domMonitorDriftSvc, reqDomDriftSvc)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(domAggregateMonitor)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.MonitorDriftInActiveResponseDTO)
	resDTO.Message = "DataDrift Monitor Inactive Success"

	return resDTO, nil
}

func (s *MonitorService) GetFeatureDetail(req *appDTO.FeatureDriftGetRequestDTO) (*appDTO.FeatureDriftGetResponseDTO, error) {
	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)
	if err != nil {
		return nil, err
	}
	reqDomDriftSvc := domSvcMonitorSvcDriftDTO.DataDriftGetRequest{
		InferenceName:  req.DeploymentID,
		ModelHistoryID: req.ModelHistoryID,
		StartTime:      req.StartTime,
		EndTime:        req.EndTime,
	}
	res, err := domAggregateMonitor.GetFeatureDetail(s.domMonitorDriftSvc, reqDomDriftSvc)
	if err != nil {
		return nil, err
	}
	resDTO := new(appDTO.FeatureDriftGetResponseDTO)
	resDTO.Message = res.Message
	resDTO.StartTime = res.StartTime
	resDTO.EndTime = res.EndTime
	resDTO.Data = res.Data
	resDTO.PredictionCount = res.PredictionCount

	return resDTO, nil
}

func (s *MonitorService) GetFeatureDrift(req *appDTO.FeatureDriftGetRequestDTO) (*appDTO.FeatureDriftGetResponseDTO, error) {
	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)
	if err != nil {
		return nil, err
	}
	reqDomDriftSvc := domSvcMonitorSvcDriftDTO.DataDriftGetRequest{
		InferenceName:  req.DeploymentID,
		ModelHistoryID: req.ModelHistoryID,
		StartTime:      req.StartTime,
		EndTime:        req.EndTime,
	}
	res, err := domAggregateMonitor.GetFeatureDrift(s.domMonitorDriftSvc, reqDomDriftSvc)
	if err != nil {
		return nil, err
	}
	resDTO := new(appDTO.FeatureDriftGetResponseDTO)
	resDTO.Message = res.Message
	resDTO.StartTime = res.StartTime
	resDTO.EndTime = res.EndTime
	resDTO.Data = res.Data
	resDTO.PredictionCount = res.PredictionCount

	return resDTO, nil
}

func (s *MonitorService) PatchAccuracyMonitorSetting(req *appDTO.MonitorAccuracyPatchRequestDTO) (*appDTO.MonitorAccuracyPatchResponseDTO, error) {
	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)
	if err != nil {
		return nil, err
	}
	resDTO := new(appDTO.MonitorAccuracyPatchResponseDTO)
	if domAggregateMonitor.AccuracyMonitoring == false {
		resDTO.DeploymentID = req.DeploymentID
		resDTO.Message = "Accuracy Monitor is False"
		return resDTO, nil
	}

	reqDomAccuracySvc := domSvcMonitorSvcAccuracyDTO.AccuracyPatchRequest{
		InferenceName:    req.DeploymentID,
		DriftMetrics:     req.AccuracySetting.MetricType,
		DriftMeasurement: req.AccuracySetting.Measurement,
		AtriskValue:      req.AccuracySetting.AtRiskValue,
		FailingValue:     req.AccuracySetting.FailingValue,
	}

	err = domAggregateMonitor.PatchAccuracySetting(s.domMonitorAccuracySvc, reqDomAccuracySvc)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(domAggregateMonitor)
	if err != nil {
		return nil, err
	}

	resDTO.DeploymentID = req.DeploymentID
	resDTO.Message = "Accuracy Monitor Patch Success"

	return resDTO, nil
}

func (s *MonitorService) SetAccuracyMonitorActive(req *appDTO.MonitorAccuracyActiveRequestDTO) (*appDTO.MonitorAccuracyActiveResponseDTO, error) {
	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)

	// modelHistoryDTO := new(appDeploymentDTO.GetModelHistoryRequestDTO)
	// modelHistoryDTO.DeploymentID = domAggregateMonitor.ID

	// modelHistoryRes, err := s.deploymentSvc.GetModelHistory(modelHistoryDTO)
	// modelHistory := modelHistoryRes.ModelHistory.([]entity.ModelHistory)

	// var modelHistoryID string
	// for _, v := range modelHistory {
	// 	if v.ApplyHistoryTag == "current" {
	// 		modelHistoryID = v.Version
	// 	}
	// }

	resModelPackage, err := s.getModelPackageByID(req.ModelPackageID)
	if err != nil {
		return nil, err
	}

	reqDomAccuracySvc := domSvcMonitorSvcAccuracyDTO.AccuracyCreateRequest{
		InferenceName:    req.DeploymentID,
		ModelHistoryID:   req.CurrentModelID,
		DatasetPath:      resModelPackage.HoldoutDatasetPath,
		ModelPath:        resModelPackage.ModelFilePath,
		TargetLabel:      resModelPackage.PredictionTargetName,
		AssociationID:    req.AssociationID,
		ModelType:        resModelPackage.TargetType,
		Framework:        resModelPackage.ModelFrameWork,
		DriftMetrics:     req.AccuracySetting.MetricType,
		DriftMeasurement: req.AccuracySetting.Measurement,
		AtriskValue:      req.AccuracySetting.AtRiskValue,
		FailingValue:     req.AccuracySetting.FailingValue,
		PositiveClass:    resModelPackage.PositiveClassLabel,
		NegativeClass:    resModelPackage.NegativeClassLabel,
		BinaryThreshold:  resModelPackage.PredictionThreshold,
	}

	err = domAggregateMonitor.SetAccuracyMonitoringOn(s.domMonitorAccuracySvc, reqDomAccuracySvc)
	if err != nil {
		return nil, err
	}
	err = s.repo.Save(domAggregateMonitor)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.MonitorAccuracyActiveResponseDTO)
	resDTO.DeploymentID = domAggregateMonitor.ID

	return resDTO, nil
}

func (s *MonitorService) SetAccuracyMonitorInActive(req *appDTO.MonitorAccuracyInActiveRequestDTO) (*appDTO.MonitorAccuracyInActiveResponseDTO, error) {
	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)
	if err != nil {
		return nil, err
	}
	reqDomAccuracySvc := domSvcMonitorSvcAccuracyDTO.AccuracyDeleteRequest{
		InferenceName: req.DeploymentID,
	}
	err = domAggregateMonitor.SetAccuracyMonitoringOff(s.domMonitorAccuracySvc, reqDomAccuracySvc)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(domAggregateMonitor)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.MonitorAccuracyInActiveResponseDTO)
	resDTO.Message = "Accuracy Monitor Inactive Success"

	return resDTO, nil
}

func (s *MonitorService) GetAccuracy(req *appDTO.AccuracyGetRequestDTO) (*appDTO.AccuracyGetResponseDTO, error) {
	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)
	if err != nil {
		return nil, err
	}
	reqDomAccuracySvc := domSvcMonitorSvcAccuracyDTO.AccuracyGetRequest{
		InferenceName:  req.DeploymentID,
		ModelHistoryID: req.ModelHistoryID,
		DataType:       req.Type,
		StartTime:      req.StartTime,
		EndTime:        req.EndTime,
	}

	res, err := domAggregateMonitor.GetAccuracy(s.domMonitorAccuracySvc, reqDomAccuracySvc)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.AccuracyGetResponseDTO)
	resDTO.Message = res.Message
	resDTO.Data = res.Data
	resDTO.StartTime = res.StartTime
	resDTO.EndTime = res.EndTIme

	return resDTO, nil
}

func (s *MonitorService) GetByID(req *appDTO.MonitorGetByIDRequestDTO) (*appDTO.MonitorGetByIDResponseDTO, error) {
	res, err := s.repo.Get(req.ID)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.MonitorGetByIDResponseDTO)
	resDTO.Monitor = res

	return resDTO, nil
}

func (s *MonitorService) getModelPackageByID(modelPackageID string) (*appModelPackageDTO.InternalGetModelPackageResponseDTO, error) {
	reqModelPackage := &appModelPackageDTO.InternalGetModelPackageRequestDTO{
		ModelPackageID: modelPackageID,
	}

	resModelPackage, err := s.modelPackageSvc.GetByIDInternal(reqModelPackage)
	if err != nil {
		return nil, err
	}

	return resModelPackage, err
}

func (s *MonitorService) GetMonitorSetting(req *appDTO.MonitorGetSettingRequestDTO) (*appDTO.MonitorGetSettingResponseDTO, error) {
	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)
	if err != nil {
		return nil, err
	}
	resDTO := new(appDTO.MonitorGetSettingResponseDTO)
	resDTO.AccuracySetting = appDTO.AccuracySetting(domAggregateMonitor.AccuracySetting)
	resDTO.DataDriftSetting = appDTO.DataDriftSetting(domAggregateMonitor.DataDriftSetting)

	return resDTO, nil
}

func (s *MonitorService) UploadActual(req *appDTO.UploadActualRequestDTO) (*appDTO.UploadActualResponseDTO, error) {
	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	cfg, err := s.handler.GetConfig()
	if err != nil {
		return nil, err
	}

	uploadFilePath := cfg.DirLocations.MonitoringFileRootPath + "/" + domAggregateMonitor.ID + cfg.DirLocations.ActualDataPath + "/" + req.FileName

	err = s.storageClient.UploadFile(req.File, uploadFilePath)

	if err != nil {
		return nil, err
	}
	reqDomActualSvc := domSvcMonitorSvcAccuracyDTO.AccuracyPostActualRequest{
		InferenceName:  req.DeploymentID,
		DatasetPath:    uploadFilePath,
		ActualResponse: req.ActualResponse,
	}

	res, err := domAggregateMonitor.PostActual(s.domMonitorAccuracySvc, reqDomActualSvc)
	if err != nil {
		err = s.storageClient.DeleteFile(uploadFilePath)
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	resDTO := new(appDTO.UploadActualResponseDTO)
	resDTO.Message = res.Message
	resDTO.DeploymentID = res.InferenceName

	return resDTO, nil
}

func (s *MonitorService) GetFeatureDriftGraph(req *appDTO.FeatureDriftGetRequestDTO) (*appDTO.FeatureDriftGetResponseDTO, error) {
	return nil, nil
}

func (s *MonitorService) GetFeatureDetailGraph(req *appDTO.FeatureDriftGetRequestDTO) (*appDTO.FeatureDriftGetResponseDTO, error) {
	return nil, nil
}

func (s *MonitorService) GetMonitorStatusList(req *appDTO.MonitorGetStatusListRequestDTO) (*appDTO.MonitorGetStatusListResponseDTO, error) {
	return nil, nil
}
