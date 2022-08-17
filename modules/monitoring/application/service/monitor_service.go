package service

import (
	"context"
	"io"
	"sync"

	infStorageClient "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/storage/minio"
	appModelPackageDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/application/dto"
	appDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/application/dto"
	domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/entity"
	domRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/repository"
	domSvcMonitorSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service"
	domSvcMonitorSvcAccuracyDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service/accuracy/dto"
	domSvcMonitorSvcDriftDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service/data_drift/dto"
	infAccuracySvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/infrastructure/monitor_service/accuracy"
	infDriftSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/infrastructure/monitor_service/data_drift"
	infGraphSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/infrastructure/monitor_service/graph"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"

	//domSvcMonitorSvcAccuracyDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service"
	infRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/infrastructure/repository"
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
	UploadFile(ioReader interface{}, filePath string) error
	DeleteFile(filePath string) error
	GetFile(filePath string) (io.Reader, error)
}

type MonitorService struct {
	BaseService
	domMonitorDriftSvc    domSvcMonitorSvc.IExternalDriftMonitorAdapter
	domMonitorAccuracySvc domSvcMonitorSvc.IExternalAccuracyMonitorAdapter
	domMonitorGraphSvc    domSvcMonitorSvc.IExternalGraphMonitorAdapter
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

	if svc.domMonitorDriftSvc, err = infDriftSvc.NewDataDriftAdapter(h); err != nil {
		return nil, err
	}

	if svc.domMonitorAccuracySvc, err = infAccuracySvc.NewAccuracyAdapter(h); err != nil {
		return nil, err
	}

	if svc.domMonitorGraphSvc, err = infGraphSvc.NewGraphAdapter(h); err != nil {
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

	// drift accuracy monitor 생성 request channel 로 변경하여 비동기 처리로 수정해야함함
	if err := req.Validate(); err != nil {
		return nil, err
	}

	domAggregateMonitor, err := domEntity.NewMonitor(
		req.DeploymentID,
		req.ModelPackageID,
	)
	resModelPackage, err := s.getModelPackageByID(req.ModelPackageID)
	if err != nil {
		return nil, err
	}

	// feature drift 여부에 상관없이 셋팅값은 초기셋팅으로 설정 이후 패치만 받음
	domAggregateMonitor.SetDataDriftSetting(
		"",
		"",
		0,
		0,
		1,
		0,
		0,
		1,
	)

	// accuracy 여부에 상관없이 셋팅값은 초기셋팅으로 설정 이후 패치만 받음
	domAggregateMonitor.SetAccuracySetting(
		"",
		"",
		5,
		10,
		resModelPackage.TargetType,
	)

	err = s.repo.Save(domAggregateMonitor)
	if err != nil {
		return nil, err
	}

	var wait sync.WaitGroup
	wait.Add(2)

	errs := make(chan error, 2)

	if req.FeatureDriftTracking == true {
		reqDomDriftSvc := domSvcMonitorSvcDriftDTO.DataDriftCreateRequest{
			InferenceName:              req.DeploymentID,
			ModelHistoryID:             req.ModelHistoryID,
			TargetLabel:                resModelPackage.PredictionTargetName,
			ModelType:                  resModelPackage.TargetType,
			Framework:                  resModelPackage.ModelFrameWork,
			TrainDatasetPath:           resModelPackage.TrainingDatasetPath,
			ModelPath:                  resModelPackage.ModelFilePath,
			DriftThreshold:             domAggregateMonitor.DriftThreshold,
			ImportanceThreshold:        domAggregateMonitor.ImportanceThreshold,
			MonitorRange:               domAggregateMonitor.MonitorRange,
			LowImportanceAtRiskCount:   domAggregateMonitor.LowImportanceAtRiskCount,
			LowImportanceFailingCount:  domAggregateMonitor.LowImportanceFailingCount,
			HighImportanceAtRiskCount:  domAggregateMonitor.HighImportanceAtRiskCount,
			HighImportanceFailingCount: domAggregateMonitor.HighImportanceFailingCount,
		}

		go func() {
			defer wait.Done()
			err = domAggregateMonitor.SetFeatureDriftTrackingOn(s.domMonitorDriftSvc, reqDomDriftSvc)
			if err != nil {
				errs <- err
			} else {
				domAggregateMonitor.SetDriftCreatedTrue()
			}
		}()
	} else {
		wait.Done()
	}

	if req.AccuracyMonitoring == true {
		reqDomAccuracySvc := domSvcMonitorSvcAccuracyDTO.AccuracyCreateRequest{
			InferenceName:    req.DeploymentID,
			ModelHistoryID:   req.ModelHistoryID,
			DatasetPath:      resModelPackage.HoldoutDatasetPath,
			ModelPath:        resModelPackage.ModelFilePath,
			TargetLabel:      resModelPackage.PredictionTargetName,
			ModelType:        resModelPackage.TargetType,
			Framework:        resModelPackage.ModelFrameWork,
			DriftMetrics:     domAggregateMonitor.MetricType,
			DriftMeasurement: domAggregateMonitor.Measurement,
			AtriskValue:      domAggregateMonitor.AtRiskValue,
			FailingValue:     domAggregateMonitor.FailingValue,
			PositiveClass:    resModelPackage.PositiveClassLabel,
			NegativeClass:    resModelPackage.NegativeClassLabel,
			BinaryThreshold:  resModelPackage.PredictionThreshold,
		}
		if req.AssociationID != nil {
			reqDomAccuracySvc.AssociationID = *req.AssociationID
		}
		go func() {
			defer wait.Done()
			err = domAggregateMonitor.SetAccuracyMonitoringOn(s.domMonitorAccuracySvc, reqDomAccuracySvc)
			if err != nil {
				errs <- err
			} else {
				domAggregateMonitor.SetAccuracyCreatedTrue()
			}
		}()
	} else {
		wait.Done()
	}

	wait.Wait()
	close(errs)

	var checkErrMsg error = <-errs
	if checkErrMsg != nil {
		// 에러시 만들어진 모니터 삭제 및 취소
		if domAggregateMonitor.DriftCreated == true {
			reqDomDriftSvc := domSvcMonitorSvcDriftDTO.DataDriftDeleteRequest{
				InferenceName: req.DeploymentID,
			}
			err = domAggregateMonitor.SetFeatureDriftTrackingOff(s.domMonitorDriftSvc, reqDomDriftSvc)
			if err != nil {
				return nil, err
			}
		}
		if domAggregateMonitor.AccuracyCreated == true {
			reqDomAccuracySvc := domSvcMonitorSvcAccuracyDTO.AccuracyDeleteRequest{
				InferenceName: req.DeploymentID,
			}
			err = domAggregateMonitor.SetAccuracyMonitoringOff(s.domMonitorAccuracySvc, reqDomAccuracySvc)
			if err != nil {
				return nil, err
			}
		}
		err = s.repo.Delete(domAggregateMonitor.ID)
		return nil, checkErrMsg
	}

	err = s.repo.Save(domAggregateMonitor)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.MonitorCreateResponseDTO)
	resDTO.DeploymentID = domAggregateMonitor.ID

	return resDTO, nil
}

func (s *MonitorService) MonitorReplaceModel(req *appDTO.MonitorReplaceModelRequestDTO) (*appDTO.MonitorReplaceModelResponseDTO, error) {

	if err := req.Validate(); err != nil {
		return nil, err
	}

	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	resModelPackage, err := s.getModelPackageByID(req.ModelPackageID)

	if err != nil {
		return nil, err
	}

	if domAggregateMonitor.FeatureDriftTracking == true {
		domAggregateMonitor.SetDriftCreatedFalse()
		reqDomDriftSvc := domSvcMonitorSvcDriftDTO.DataDriftCreateRequest{
			InferenceName:              req.DeploymentID,
			ModelHistoryID:             req.ModelHistoryID,
			TargetLabel:                resModelPackage.PredictionTargetName,
			ModelType:                  resModelPackage.TargetType,
			Framework:                  resModelPackage.ModelFrameWork,
			TrainDatasetPath:           resModelPackage.TrainingDatasetPath,
			ModelPath:                  resModelPackage.ModelFilePath,
			DriftThreshold:             domAggregateMonitor.DriftThreshold,
			ImportanceThreshold:        domAggregateMonitor.ImportanceThreshold,
			MonitorRange:               domAggregateMonitor.MonitorRange,
			LowImportanceAtRiskCount:   domAggregateMonitor.LowImportanceAtRiskCount,
			LowImportanceFailingCount:  domAggregateMonitor.LowImportanceFailingCount,
			HighImportanceAtRiskCount:  domAggregateMonitor.HighImportanceAtRiskCount,
			HighImportanceFailingCount: domAggregateMonitor.HighImportanceFailingCount,
		}

		err = domAggregateMonitor.SetFeatureDriftTrackingOn(s.domMonitorDriftSvc, reqDomDriftSvc)
		if err != nil {
			return nil, err
		}

		domAggregateMonitor.SetDriftCreatedTrue()

		err = s.repo.Save(domAggregateMonitor)
		if err != nil {
			return nil, err
		}
	}

	if domAggregateMonitor.AccuracyMonitoring == true {
		domAggregateMonitor.SetAccuracyCreatedFalse()
		reqDomAccuracySvc := domSvcMonitorSvcAccuracyDTO.AccuracyCreateRequest{
			InferenceName:    req.DeploymentID,
			ModelHistoryID:   req.ModelHistoryID,
			DatasetPath:      resModelPackage.HoldoutDatasetPath,
			ModelPath:        resModelPackage.ModelFilePath,
			TargetLabel:      resModelPackage.PredictionTargetName,
			AssociationID:    domAggregateMonitor.AssociationID,
			ModelType:        resModelPackage.TargetType,
			Framework:        resModelPackage.ModelFrameWork,
			DriftMetrics:     domAggregateMonitor.MetricType,
			DriftMeasurement: domAggregateMonitor.Measurement,
			AtriskValue:      domAggregateMonitor.AtRiskValue,
			FailingValue:     domAggregateMonitor.FailingValue,
			PositiveClass:    resModelPackage.PositiveClassLabel,
			NegativeClass:    resModelPackage.NegativeClassLabel,
			BinaryThreshold:  resModelPackage.PredictionThreshold,
		}

		err = domAggregateMonitor.SetAccuracyMonitoringOn(s.domMonitorAccuracySvc, reqDomAccuracySvc)
		if err != nil {
			return nil, err
		}
		domAggregateMonitor.SetAccuracyCreatedTrue()

		err = s.repo.Save(domAggregateMonitor)
		if err != nil {
			return nil, err
		}
	}

	resDTO := new(appDTO.MonitorReplaceModelResponseDTO)
	resDTO.DeploymentID = domAggregateMonitor.ID

	return resDTO, nil
}

func (s *MonitorService) Delete(req *appDTO.MonitorDeleteRequestDTO) (*appDTO.MonitorDeleteResponseDTO, error) {
	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	if domAggregateMonitor.DriftCreated == true {
		// Drift OFF
		reqDomDriftSvc := domSvcMonitorSvcDriftDTO.DataDriftDeleteRequest{
			InferenceName: req.DeploymentID,
		}
		err = domAggregateMonitor.SetFeatureDriftTrackingOff(s.domMonitorDriftSvc, reqDomDriftSvc)
		if err != nil {
			return nil, err
		}
	}

	if domAggregateMonitor.AccuracyCreated == true {
		// Accuracy OFF
		reqDomAccuracySvc := domSvcMonitorSvcAccuracyDTO.AccuracyDeleteRequest{
			InferenceName: req.DeploymentID,
		}
		err = domAggregateMonitor.SetAccuracyMonitoringOff(s.domMonitorAccuracySvc, reqDomAccuracySvc)
		if err != nil {
			return nil, err
		}
	}

	err = s.repo.Delete(domAggregateMonitor.ID)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.MonitorDeleteResponseDTO)
	resDTO.Message = "DataDrift Monitor Delete Success"

	return resDTO, nil
}

func (s *MonitorService) PatchDriftMonitorSetting(req *appDTO.MonitorDriftPatchRequestDTO) (*appDTO.MonitorDriftPatchResponseDTO, error) {
	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)
	if err != nil {
		return nil, err
	}
	resDTO := new(appDTO.MonitorDriftPatchResponseDTO)
	// false 여도 patch 가능

	if err := req.Validate(); err != nil {
		return nil, err
	}
	if err = req.DataDriftSetting.Validate(); err != nil {
		return nil, err
	}

	reqDomDriftSvc := new(domSvcMonitorSvcDriftDTO.DataDriftPatchRequest)

	reqDomDriftSvc.InferenceName = req.DeploymentID

	if req.DataDriftSetting.MonitorRange != "" {
		reqDomDriftSvc.MonitorRange = req.DataDriftSetting.MonitorRange
	} else {
		reqDomDriftSvc.MonitorRange = domAggregateMonitor.MonitorRange
	}

	if req.DataDriftSetting.DriftThreshold != nil {
		reqDomDriftSvc.DriftThreshold = *req.DataDriftSetting.DriftThreshold
	} else {
		reqDomDriftSvc.DriftThreshold = domAggregateMonitor.DriftThreshold
	}

	if req.DataDriftSetting.ImportanceThreshold != nil {
		reqDomDriftSvc.ImportanceThreshold = *req.DataDriftSetting.ImportanceThreshold
	} else {
		reqDomDriftSvc.ImportanceThreshold = domAggregateMonitor.ImportanceThreshold
	}

	if req.DataDriftSetting.LowImportanceAtRiskCount != nil {
		reqDomDriftSvc.LowImportanceAtRiskCount = *req.DataDriftSetting.LowImportanceAtRiskCount
	} else {
		reqDomDriftSvc.LowImportanceAtRiskCount = domAggregateMonitor.LowImportanceAtRiskCount
	}

	if req.DataDriftSetting.LowImportanceFailingCount != nil {
		reqDomDriftSvc.LowImportanceFailingCount = *req.DataDriftSetting.LowImportanceFailingCount
	} else {
		reqDomDriftSvc.LowImportanceFailingCount = domAggregateMonitor.LowImportanceFailingCount
	}

	if req.DataDriftSetting.HighImportanceAtRiskCount != nil {
		reqDomDriftSvc.HighImportanceAtRiskCount = *req.DataDriftSetting.HighImportanceAtRiskCount
	} else {
		reqDomDriftSvc.HighImportanceAtRiskCount = domAggregateMonitor.HighImportanceAtRiskCount
	}

	if req.DataDriftSetting.HighImportanceFailingCount != nil {
		reqDomDriftSvc.HighImportanceFailingCount = *req.DataDriftSetting.HighImportanceFailingCount
	} else {
		reqDomDriftSvc.HighImportanceFailingCount = domAggregateMonitor.HighImportanceFailingCount
	}

	err = domAggregateMonitor.PatchDataDriftSetting(s.domMonitorDriftSvc, *reqDomDriftSvc)
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
	// active 는 드리프트 셋팅과 연관이 없음. 단순 on 만 가능하게 하는데 drift monitor가 만들어지지 않을경우엔 생성, 이미 있을 경우엔 단순 on만. drift monitor 생성여부 상태값 저장해야함
	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)

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
		TrainDatasetPath:           resModelPackage.TrainingDatasetPath,
		ModelPath:                  resModelPackage.ModelFilePath,
		DriftThreshold:             domAggregateMonitor.DriftThreshold,
		ImportanceThreshold:        domAggregateMonitor.ImportanceThreshold,
		MonitorRange:               domAggregateMonitor.MonitorRange,
		LowImportanceAtRiskCount:   domAggregateMonitor.LowImportanceAtRiskCount,
		LowImportanceFailingCount:  domAggregateMonitor.LowImportanceFailingCount,
		HighImportanceAtRiskCount:  domAggregateMonitor.HighImportanceAtRiskCount,
		HighImportanceFailingCount: domAggregateMonitor.HighImportanceFailingCount,
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
	// 단순 off 만 구현
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

	if err := req.Validate(); err != nil {
		return nil, err
	}

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

	if err := req.Validate(); err != nil {
		return nil, err
	}

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
	// false 여도 patch는 가능

	if err := req.Validate(); err != nil {
		return nil, err
	}
	if err = req.AccuracySetting.Validate(); err != nil {
		return nil, err
	}

	reqDomAccuracySvc := new(domSvcMonitorSvcAccuracyDTO.AccuracyPatchRequest)

	reqDomAccuracySvc.InferenceName = req.DeploymentID

	if req.AccuracySetting.MetricType != "" {
		reqDomAccuracySvc.DriftMetrics = req.AccuracySetting.MetricType
	} else {
		reqDomAccuracySvc.DriftMetrics = domAggregateMonitor.MetricType
	}

	if req.AccuracySetting.Measurement != "" {
		reqDomAccuracySvc.DriftMeasurement = req.AccuracySetting.Measurement
	} else {
		reqDomAccuracySvc.DriftMeasurement = domAggregateMonitor.Measurement
	}

	if req.AccuracySetting.AtRiskValue != nil {
		reqDomAccuracySvc.AtriskValue = *req.AccuracySetting.AtRiskValue
	} else {
		reqDomAccuracySvc.AtriskValue = domAggregateMonitor.AtRiskValue
	}

	if req.AccuracySetting.FailingValue != nil {
		reqDomAccuracySvc.FailingValue = *req.AccuracySetting.FailingValue
	} else {
		reqDomAccuracySvc.FailingValue = domAggregateMonitor.FailingValue
	}

	err = domAggregateMonitor.PatchAccuracySetting(s.domMonitorAccuracySvc, *reqDomAccuracySvc)
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

func (s *MonitorService) UpdateAssociationID(req *appDTO.UpdateAssociationIDRequestDTO) (*appDTO.UpdateAssociationIDResponseDTO, error) {
	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	reqDomAccuracySvc := new(domSvcMonitorSvcAccuracyDTO.AccuracyUpdateAssociationIDRequest)
	reqDomAccuracySvc.InferenceName = req.DeploymentID

	if req.AssociationID != nil {
		reqDomAccuracySvc.AssociationID = *req.AssociationID
	}

	err = domAggregateMonitor.SetAssociationID(s.domMonitorAccuracySvc, *reqDomAccuracySvc)
	if err != nil {
		return nil, err
	}
	err = s.repo.Save(domAggregateMonitor)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.UpdateAssociationIDResponseDTO)
	resDTO.Message = "Association ID change Success"
	resDTO.DeploymentID = domAggregateMonitor.ID

	return resDTO, nil
}

func (s *MonitorService) SetAccuracyMonitorActive(req *appDTO.MonitorAccuracyActiveRequestDTO) (*appDTO.MonitorAccuracyActiveResponseDTO, error) {
	// drift와 동일하게 on 기능 + accuracy monitor가 없을 경우 생성 있을경우 on 만. accuracy monitor 생성 여부 상태 저장해야함
	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)

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
		ModelType:        resModelPackage.TargetType,
		Framework:        resModelPackage.ModelFrameWork,
		DriftMetrics:     domAggregateMonitor.MetricType,
		DriftMeasurement: domAggregateMonitor.Measurement,
		AtriskValue:      domAggregateMonitor.AtRiskValue,
		FailingValue:     domAggregateMonitor.FailingValue,
		PositiveClass:    resModelPackage.PositiveClassLabel,
		NegativeClass:    resModelPackage.NegativeClassLabel,
		BinaryThreshold:  resModelPackage.PredictionThreshold,
	}
	if req.AssociationID != nil {
		reqDomAccuracySvc.AssociationID = *req.AssociationID
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

	if err := req.Validate(); err != nil {
		return nil, err
	}

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

	if err := req.Validate(); err != nil {
		return nil, err
	}

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

	if err := req.Validate(); err != nil {
		return nil, err
	}

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

	if err := req.Validate(); err != nil {
		return nil, err
	}

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
		InferenceName:     req.DeploymentID,
		DatasetPath:       uploadFilePath,
		ActualResponse:    req.ActualResponse,
		AssociationColumn: req.AssociationColumn,
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

//func (s *MonitorService) GetFeatureDriftGraph(req *appDTO.DriftGraphGetRequestDTO) (*appDTO.DriftGraphGetResponseDTO, error) {
//
//	if err := req.Validate(); err != nil {
//		return nil, err
//	}
//
//	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)
//
//	if err != nil {
//		return nil, err
//	}
//	reqDomDriftGraphSvc := domSvcMonitorSvcGraphDTO.DriftGraphGetRequest{
//		InferenceName:       req.DeploymentID,
//		ModelHistoryID:      req.ModelHistoryID,
//		StartTime:           convertTimestamp(req.StartTime),
//		EndTime:             convertTimestamp(req.EndTime),
//		HostEndpoint:        req.HostEndpoint,
//		DriftThreshold:      domAggregateMonitor.DriftThreshold,
//		ImportanceThreshold: domAggregateMonitor.ImportanceThreshold,
//	}
//	res, err := domAggregateMonitor.GetDriftGraph(s.domMonitorGraphSvc, reqDomDriftGraphSvc)
//	if err != nil {
//		return nil, err
//	}
//
//	resDTO := new(appDTO.DriftGraphGetResponseDTO)
//	resDTO.Script = res.Script
//
//	return resDTO, nil
//}

//func (s *MonitorService) GetFeatureDetailGraph(req *appDTO.DetailGraphGetRequestDTO) (*appDTO.DetailGraphGetResponseDTO, error) {
//
//	if err := req.Validate(); err != nil {
//		return nil, err
//	}
//
//	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)
//	if err != nil {
//		return nil, err
//	}
//	reqDomDetailGraphSvc := domSvcMonitorSvcGraphDTO.DetailGraphGetRequest{
//		InferenceName:  req.DeploymentID,
//		ModelHistoryID: req.ModelHistoryID,
//		StartTime:      convertTimestamp(req.StartTime),
//		EndTime:        convertTimestamp(req.EndTime),
//		HostEndpoint:   req.HostEndpoint,
//	}
//	res, err := domAggregateMonitor.GetDetailGraph(s.domMonitorGraphSvc, reqDomDetailGraphSvc)
//	if err != nil {
//		return nil, err
//	}
//
//	resDTO := new(appDTO.DetailGraphGetResponseDTO)
//	resDTO.Script = res.Script
//	return resDTO, nil
//}

//func convertTimestamp(timeString string) string {
//	t, _ := time.Parse("2006-01-02:15", timeString)
//
//	t2 := t.UTC().Format("2006-01-02T15:04:05.00000000Z")
//
//	return t2
//}
