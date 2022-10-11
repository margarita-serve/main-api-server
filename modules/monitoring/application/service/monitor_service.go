package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sync"

	infStorageClient "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/storage/minio"
	common "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/common"
	appDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/application/dto"
	domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/entity"
	domRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/repository"
	domSvcMonitorSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service"
	domSvcMonitorSvcAccuracyDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service/accuracy/dto"
	domSvcMonitorSvcDriftDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service/data_drift/dto"
	domSvcMonitorSvcServiceHealthDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service/service_health/dto"
	infAccuracySvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/infrastructure/monitor_service/accuracy"
	infDriftSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/infrastructure/monitor_service/data_drift"
	infServiceHealthSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/infrastructure/monitor_service/service_health"

	//appNotiDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/application/dto"
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

// type INotiService interface {
// 	SendNoti(req *appNotiDTO.NotiRequestDTO, i identity.Identity) error
// }

type StorageClient interface {
	UploadFile(ioReader interface{}, filePath string) error
	DeleteFile(filePath string) error
	GetFile(filePath string) (io.Reader, error)
}

type MonitorService struct {
	BaseService
	domMonitorDriftSvc         domSvcMonitorSvc.IExternalDriftMonitorAdapter
	domMonitorAccuracySvc      domSvcMonitorSvc.IExternalAccuracyMonitorAdapter
	domMonitorServiceHealthSvc domSvcMonitorSvc.IExternalServiceHealthMonitorAdapter
	modelPackageSvc            common.IModelPackageService
	//notiSvc                    INotiService
	repo          domRepo.IMonitorRepo
	storageClient StorageClient
	publisher     common.EventPublisher
}

//func NewMonitorService(h *handler.Handler, modelPackageSvc common.IModelPackageService, notiSvc INotiService, publisher common.EventPublisher) (*MonitorService, error) {

func NewMonitorService(h *handler.Handler, modelPackageSvc common.IModelPackageService, publisher common.EventPublisher) (*MonitorService, error) {
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

	if svc.domMonitorServiceHealthSvc, err = infServiceHealthSvc.NewServiceHealthAdapter(h); err != nil {
		return nil, err
	}

	svc.modelPackageSvc = modelPackageSvc

	//svc.notiSvc = notiSvc

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

	svc.publisher = publisher

	return svc, nil
}

func (s *MonitorService) Create(req *appDTO.MonitorCreateRequestDTO) (*appDTO.MonitorCreateResponseDTO, error) {

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
	wait.Add(3)

	errs := make(chan error, 3)

	// 서비스 상태 모니터는 항상 생성한다
	go func() {
		defer wait.Done()
		reqDomServiceHealthSvc := domSvcMonitorSvcServiceHealthDTO.ServiceHealthCreateRequest{
			InferenceName:  req.DeploymentID,
			ModelHistoryID: req.ModelHistoryID,
		}
		err = domAggregateMonitor.SetServiceHealthMonitorTrackingOn(s.domMonitorServiceHealthSvc, reqDomServiceHealthSvc)
		if err != nil {
			errs <- err
		} else {
			domAggregateMonitor.SetServiceHealthCreatedTrue()
		}
	}()

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
			InferenceName:          req.DeploymentID,
			ModelHistoryID:         req.ModelHistoryID,
			DatasetPath:            resModelPackage.HoldoutDatasetPath,
			ModelPath:              resModelPackage.ModelFilePath,
			TargetLabel:            resModelPackage.PredictionTargetName,
			AssociationID:          *req.AssociationID,
			AssociationIDInFeature: req.AssociationIDInFeature,
			ModelType:              resModelPackage.TargetType,
			Framework:              resModelPackage.ModelFrameWork,
			DriftMetric:            domAggregateMonitor.MetricType,
			DriftMeasurement:       domAggregateMonitor.Measurement,
			AtriskValue:            domAggregateMonitor.AtRiskValue,
			FailingValue:           domAggregateMonitor.FailingValue,
			PositiveClass:          resModelPackage.PositiveClassLabel,
			NegativeClass:          resModelPackage.NegativeClassLabel,
			BinaryThreshold:        resModelPackage.PredictionThreshold,
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

	err = s.repo.Save(domAggregateMonitor)
	if err != nil {
		return nil, err
	}

	var checkErrMsg error = <-errs
	if checkErrMsg != nil {
		//// 에러시 만들어진 모니터 삭제 및 취소
		//reqDomServiceHealthSvc := domSvcMonitorSvcServiceHealthDTO.ServiceHealthDeleteRequest{
		//	InferenceName: req.DeploymentID,
		//}
		//err = domAggregateMonitor.SetServiceHealthMonitorTrackingOff(s.domMonitorServiceHealthSvc, reqDomServiceHealthSvc)
		//if err != nil {
		//	return nil, err
		//}
		//
		//if domAggregateMonitor.DriftCreated == true {
		//	reqDomDriftSvc := domSvcMonitorSvcDriftDTO.DataDriftDeleteRequest{
		//		InferenceName: req.DeploymentID,
		//	}
		//	err = domAggregateMonitor.SetFeatureDriftTrackingOff(s.domMonitorDriftSvc, reqDomDriftSvc)
		//	if err != nil {
		//		return nil, err
		//	}
		//}
		//if domAggregateMonitor.AccuracyCreated == true {
		//	reqDomAccuracySvc := domSvcMonitorSvcAccuracyDTO.AccuracyDeleteRequest{
		//		InferenceName: req.DeploymentID,
		//	}
		//	err = domAggregateMonitor.SetAccuracyMonitoringOff(s.domMonitorAccuracySvc, reqDomAccuracySvc)
		//	if err != nil {
		//		return nil, err
		//	}
		//}
		//err = s.repo.Delete(domAggregateMonitor.ID)
		return nil, checkErrMsg
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

	// ServiceHealth
	domAggregateMonitor.SetServiceHealthCreatedFalse()
	reqDomServiceHealthSvc := domSvcMonitorSvcServiceHealthDTO.ServiceHealthCreateRequest{
		InferenceName:  req.DeploymentID,
		ModelHistoryID: req.ModelHistoryID,
	}
	err = domAggregateMonitor.SetServiceHealthMonitorTrackingOn(s.domMonitorServiceHealthSvc, reqDomServiceHealthSvc)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(domAggregateMonitor)
	if err != nil {
		return nil, err
	}

	// Drift
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

	// Accuracy
	if domAggregateMonitor.AccuracyMonitoring == true {
		domAggregateMonitor.SetAccuracyCreatedFalse()
		reqDomAccuracySvc := domSvcMonitorSvcAccuracyDTO.AccuracyCreateRequest{
			InferenceName:          req.DeploymentID,
			ModelHistoryID:         req.ModelHistoryID,
			DatasetPath:            resModelPackage.HoldoutDatasetPath,
			ModelPath:              resModelPackage.ModelFilePath,
			TargetLabel:            resModelPackage.PredictionTargetName,
			AssociationID:          domAggregateMonitor.AssociationID,
			AssociationIDInFeature: domAggregateMonitor.AssociationIDInFeature,
			ModelType:              resModelPackage.TargetType,
			Framework:              resModelPackage.ModelFrameWork,
			DriftMetric:            domAggregateMonitor.MetricType,
			DriftMeasurement:       domAggregateMonitor.Measurement,
			AtriskValue:            domAggregateMonitor.AtRiskValue,
			FailingValue:           domAggregateMonitor.FailingValue,
			PositiveClass:          resModelPackage.PositiveClassLabel,
			NegativeClass:          resModelPackage.NegativeClassLabel,
			BinaryThreshold:        resModelPackage.PredictionThreshold,
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

	reqDomServiceHealthSvc := domSvcMonitorSvcServiceHealthDTO.ServiceHealthDeleteRequest{
		InferenceName: req.DeploymentID,
	}
	err = domAggregateMonitor.SetServiceHealthMonitorTrackingOff(s.domMonitorServiceHealthSvc, reqDomServiceHealthSvc)
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

func (s *MonitorService) SetServiceHealthMonitorActive(req *appDTO.MonitorServiceHealthActiveRequestDTO) (*appDTO.MonitorServiceHealthActiveResponseDTO, error) {
	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	reqDomServiceHealthSvc := domSvcMonitorSvcServiceHealthDTO.ServiceHealthCreateRequest{
		InferenceName:  req.DeploymentID,
		ModelHistoryID: req.CurrentModelID,
	}

	err = domAggregateMonitor.SetServiceHealthMonitorTrackingOn(s.domMonitorServiceHealthSvc, reqDomServiceHealthSvc)
	if err != nil {
		return nil, err
	}
	err = s.repo.Save(domAggregateMonitor)
	if err != nil {
		return nil, err
	}
	resDTO := new(appDTO.MonitorServiceHealthActiveResponseDTO)
	resDTO.DeploymentID = domAggregateMonitor.ID

	return resDTO, nil
}

func (s *MonitorService) SetServiceHealthMonitorInActive(req *appDTO.MonitorServiceHealthInActiveRequestDTO) (*appDTO.MonitorServiceHealthInActiveResponseDTO, error) {
	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)
	if err != nil {
		return nil, err
	}
	reqDomServiceHealthSvc := domSvcMonitorSvcServiceHealthDTO.ServiceHealthDeleteRequest{
		InferenceName: req.DeploymentID,
	}
	err = domAggregateMonitor.SetServiceHealthMonitorTrackingOff(s.domMonitorServiceHealthSvc, reqDomServiceHealthSvc)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(domAggregateMonitor)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.MonitorServiceHealthInActiveResponseDTO)
	resDTO.Message = "ServiceHealth Monitor InActive Success"

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
	resDTO.Message = "DataDrift Monitor InActive Success"

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

func (s *MonitorService) PatchMonitorSetting(req *appDTO.MonitorPatchRequestDTO) (*appDTO.MonitorPatchResponseDTO, error) {

	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)
	if err != nil {
		return nil, err
	}
	resDTO := new(appDTO.MonitorPatchResponseDTO)
	// false 여도 patch는 가능

	if err := req.Validate(); err != nil {
		return nil, err
	}
	if err = req.DataDriftSetting.Validate(); err != nil {
		return nil, err
	}

	// data drift monitor setting
	reqDomDriftSvc := new(domSvcMonitorSvcDriftDTO.DataDriftPatchRequest)
	reqBackupDrift := new(domSvcMonitorSvcDriftDTO.DataDriftPatchRequest)
	reqBackupDrift.DriftThreshold = domAggregateMonitor.DriftThreshold
	reqBackupDrift.InferenceName = req.DeploymentID
	reqBackupDrift.MonitorRange = domAggregateMonitor.MonitorRange
	reqBackupDrift.HighImportanceFailingCount = domAggregateMonitor.HighImportanceFailingCount
	reqBackupDrift.HighImportanceAtRiskCount = domAggregateMonitor.HighImportanceAtRiskCount
	reqBackupDrift.LowImportanceFailingCount = domAggregateMonitor.LowImportanceFailingCount
	reqBackupDrift.LowImportanceAtRiskCount = domAggregateMonitor.LowImportanceAtRiskCount
	reqBackupDrift.ImportanceThreshold = domAggregateMonitor.ImportanceThreshold

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
		return nil, fmt.Errorf("monitor Setting Patch Fail, error:%s", err)
	}

	// accuracy monitor setting
	reqDomAccuracySvc := new(domSvcMonitorSvcAccuracyDTO.AccuracyPatchRequest)

	reqDomAccuracySvc.InferenceName = req.DeploymentID

	if req.AccuracySetting.MetricType != "" {
		reqDomAccuracySvc.DriftMetric = req.AccuracySetting.MetricType
	} else {
		reqDomAccuracySvc.DriftMetric = domAggregateMonitor.MetricType
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
		err = domAggregateMonitor.PatchDataDriftSetting(s.domMonitorDriftSvc, *reqBackupDrift)
		if err != nil {
			err = s.repo.Save(domAggregateMonitor)
			if err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("data drift setting patch success, but accuracy setting path fail")
		}
		return nil, fmt.Errorf("monitor Setting Patch Fail, error:%s", err)
	}

	err = s.repo.Save(domAggregateMonitor)
	if err != nil {
		return nil, err
	}

	resDTO.DeploymentID = req.DeploymentID
	resDTO.Message = "Monitor Setting Patch Success"

	return resDTO, nil
}

//func (s *MonitorService) UpdateAssociationID(req *appDTO.UpdateAssociationIDRequestDTO) (*appDTO.UpdateAssociationIDResponseDTO, error) {
//	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)
//	if err != nil {
//		return nil, err
//	}
//
//	reqDomAccuracySvc := new(domSvcMonitorSvcAccuracyDTO.AccuracyUpdateAssociationIDRequest)
//	reqDomAccuracySvc.InferenceName = req.DeploymentID
//
//	if req.AssociationID != nil {
//		reqDomAccuracySvc.AssociationID = *req.AssociationID
//	}
//
//	err = domAggregateMonitor.SetAssociationID(s.domMonitorAccuracySvc, *reqDomAccuracySvc)
//	if err != nil {
//		return nil, err
//	}
//	err = s.repo.Save(domAggregateMonitor)
//	if err != nil {
//		return nil, err
//	}
//
//	resDTO := new(appDTO.UpdateAssociationIDResponseDTO)
//	resDTO.Message = "Association ID change Success"
//	resDTO.DeploymentID = domAggregateMonitor.ID
//
//	return resDTO, nil
//}

func (s *MonitorService) SetAccuracyMonitorActive(req *appDTO.MonitorAccuracyActiveRequestDTO) (*appDTO.MonitorAccuracyActiveResponseDTO, error) {
	// drift와 동일하게 on 기능 + accuracy monitor가 없을 경우 생성 있을경우 on 만. accuracy monitor 생성 여부 상태 저장해야함
	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)

	resModelPackage, err := s.getModelPackageByID(req.ModelPackageID)
	if err != nil {
		return nil, err
	}

	if domAggregateMonitor.AssociationID != "" {
		if *req.AssociationID != "" {
			return nil, fmt.Errorf("AssociationID(AssociationIDInFeature) value already exists. AssociationID(AssociationIDInFeature) cannot be changed")
		}
	}

	reqDomAccuracySvc := domSvcMonitorSvcAccuracyDTO.AccuracyCreateRequest{
		InferenceName:          req.DeploymentID,
		ModelHistoryID:         req.CurrentModelID,
		DatasetPath:            resModelPackage.HoldoutDatasetPath,
		ModelPath:              resModelPackage.ModelFilePath,
		TargetLabel:            resModelPackage.PredictionTargetName,
		AssociationID:          *req.AssociationID,
		AssociationIDInFeature: req.AssociationIDInFeature,
		ModelType:              resModelPackage.TargetType,
		Framework:              resModelPackage.ModelFrameWork,
		DriftMetric:            domAggregateMonitor.MetricType,
		DriftMeasurement:       domAggregateMonitor.Measurement,
		AtriskValue:            domAggregateMonitor.AtRiskValue,
		FailingValue:           domAggregateMonitor.FailingValue,
		PositiveClass:          resModelPackage.PositiveClassLabel,
		NegativeClass:          resModelPackage.NegativeClassLabel,
		BinaryThreshold:        resModelPackage.PredictionThreshold,
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
	resDTO.Message = "Accuracy Monitor InActive Success"

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

func (s *MonitorService) getModelPackageByID(modelPackageID string) (*common.InternalGetModelPackageResponseDTO, error) {
	resModelPackage, err := s.modelPackageSvc.GetByIDInternal(modelPackageID)
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

	if domAggregateMonitor.AccuracyMonitoring == false {
		return nil, errors.New("accuracy monitoring is turned off")
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
		errs := s.storageClient.DeleteFile(uploadFilePath)
		if errs != nil {
			return nil, errs
		}
		return nil, err
	}
	resDTO := new(appDTO.UploadActualResponseDTO)
	resDTO.Message = res.Message
	resDTO.DeploymentID = res.InferenceName

	return resDTO, nil
}

func (s *MonitorService) monitorStatusCheck(req *appDTO.MonitorStatusCheckRequestDTO) error {
	// 데이터 드리프트 = 사욪자 지정한 드리프트 모니터 셋팅에 따라 결정 /30초간격
	// 정확도 = 사용자 지정한 정확도 모니터 셋팅에 따라 결정 /30초간격
	// 서비스 상태 = 고정값으로 현재부터 24시간까지의 데이터에 따라 결정. /60초간격
	//            24시간동안 요청이 없을경우 = unknown, 4xx >=1 인경우 = warning, 5xx >=1 인경우 = failing, 4xx or 5xx 없을경우 = pass
	// 데이터 드리프트와 정확도 모니터링은 모델의 버전에 따라 모니터링 하지만 서비스 상태의 경우 모델의 버전과 상관없이 해당 배포의 24시간 범위를 모니터링 한다.
	// noti 규칙 => 상태가 변할때 단 한번 노티를 한다. 하지만  (모든상태 -> unknown), (모든상태 -> pass) 일 경우에는 노티하지 않는다.
	domAggregateMonitor, err := s.repo.Get(req.DeploymentID)
	if err != nil {
		return err
	}
	if req.Kind == "datadrift" {
		result, noti := domAggregateMonitor.CheckDriftStatus(req.Status)
		if result == true {
			err = s.repo.Save(domAggregateMonitor)
			if err != nil {
				return err
			}
			if noti == true {
				// reqNotiSvc := appNotiDTO.NotiRequestDTO{
				// 	DeploymentID:   req.DeploymentID,
				// 	NotiCategory:   "Datadrift",
				// 	AdditionalData: fmt.Sprintf("status : %s", req.Status),
				// }
				// err = s.notiSvc.SendNoti(&reqNotiSvc, s.systemIdentity)
				// if err != nil {
				// 	return err
				// }

				switch req.Status {
				case "failing":
					s.publisher.Notify(common.NewEventMonitoringDataDriftStatusChangedToFailing(req.DeploymentID))
				case "atrisk":
					s.publisher.Notify(common.NewEventMonitoringDataDriftStatusChangedToAtrisk(req.DeploymentID))
				default:
					return errors.New("datadrift status check process error")
				}
			}
		}
	} else if req.Kind == "accuracy" {
		result, noti := domAggregateMonitor.CheckAccuracyStatus(req.Status)
		if result == true {
			err = s.repo.Save(domAggregateMonitor)
			if err != nil {
				return err
			}
			if noti == true {
				// reqNotiSvc := appNotiDTO.NotiRequestDTO{
				// 	DeploymentID:   req.DeploymentID,
				// 	NotiCategory:   "Accuracy",
				// 	AdditionalData: fmt.Sprintf("status : %s", req.Status),
				// }
				// err = s.notiSvc.SendNoti(&reqNotiSvc, s.systemIdentity)
				// if err != nil {
				// 	return err
				// }
				switch req.Status {
				case "failing":
					s.publisher.Notify(common.NewEventMonitoringAccuracyStatusChangedToFailing(req.DeploymentID))
				case "atrisk":
					s.publisher.Notify(common.NewEventMonitoringAccuracyStatusChangedToAtrisk(req.DeploymentID))
				default:
					return errors.New("accurancy status check process error")
				}

			}
		}
	} else if req.Kind == "servicehealth" {
		result, noti := domAggregateMonitor.CheckServiceHealthStatus(req.Status)
		if result == true {
			err = s.repo.Save(domAggregateMonitor)
			if err != nil {
				return err
			}
			if noti == true {
				// reqNotiSvc := appNotiDTO.NotiRequestDTO{
				// 	DeploymentID:   req.DeploymentID,
				// 	NotiCategory:   "Service",
				// 	AdditionalData: fmt.Sprintf("status : %s", req.Status),
				// }
				// err = s.notiSvc.SendNoti(&reqNotiSvc, s.systemIdentity)
				// if err != nil {
				// 	return err
				// }

				switch req.Status {
				case "failing":
					s.publisher.Notify(common.NewEventMonitoringServiceHealthStatusChangedToFailing(req.DeploymentID))
				case "atrisk":
					s.publisher.Notify(common.NewEventMonitoringServiceHealthStatusChangedToAtrisk(req.DeploymentID))
				default:
					return errors.New("service health status check process error")
				}
			}

		}
	} else {
		return fmt.Errorf("kind error")
	}

	return nil

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

func (s *MonitorService) Update(event common.Event) {
	switch actualEvent := event.(type) {
	case common.DeploymentInferenceServiceCreated:
		//
		wtfAssociaionID := actualEvent.AssociationID()

		req := &appDTO.MonitorCreateRequestDTO{
			DeploymentID:           actualEvent.DeploymentID(),
			ModelPackageID:         actualEvent.ModelPackageID(),
			ModelHistoryID:         actualEvent.ModelHistoryID(),
			AccuracyMonitoring:     actualEvent.AccuracyMonitoring(),
			FeatureDriftTracking:   actualEvent.FeatureDriftTracking(),
			AssociationID:          &wtfAssociaionID,
			AssociationIDInFeature: actualEvent.AssociationIDInFeature(),
		}
		_, err := s.Create(req)
		if err != nil {
			s.publisher.Notify(common.NewEventMonitoringCreateFailed(actualEvent.DeploymentID(), err))
			return
		}

		s.publisher.Notify(common.NewEventMonitoringCreated(actualEvent.DeploymentID()))

	case common.DeploymentModelReplaced:
		//
		req := &appDTO.MonitorReplaceModelRequestDTO{
			DeploymentID:   actualEvent.DeploymentID(),
			ModelPackageID: actualEvent.ModelPackageID(),
			ModelHistoryID: actualEvent.ModelHistoryID(),
		}
		s.MonitorReplaceModel(req)

	//case common.DeploymentAssociationIDUpdated:
	//	//
	//	strAssociaionID := actualEvent.AssociationID()
	//	reqUpdateAssociationID := new(appDTO.UpdateAssociationIDRequestDTO)
	//	reqUpdateAssociationID.DeploymentID = actualEvent.DeploymentID()
	//	reqUpdateAssociationID.AssociationID = &strAssociaionID
	//	//toBe..
	//	//reqUpdateAssociationID.AssociationIDInFeature = req.AssociationIDInFeature
	//
	//	_, err := s.UpdateAssociationID(reqUpdateAssociationID)
	//	if err != nil {
	//		fmt.Printf("err: %v\n", err)
	//	}

	case common.DeploymentFeatureDriftTrackingEnabled:
		println("FeatureDriftTracking true")

		reqDriftActive := new(appDTO.MonitorDriftActiveRequestDTO)
		reqDriftActive.DeploymentID = actualEvent.DeploymentID()
		reqDriftActive.ModelPackageID = actualEvent.ModelPackageID()
		reqDriftActive.CurrentModelID = actualEvent.CurrentModelID()

		_, err := s.SetDriftMonitorActive(reqDriftActive)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

	case common.DeploymentFeatureDriftTrackingDisabled:
		reqDriftInActive := new(appDTO.MonitorDriftInActiveRequestDTO)
		reqDriftInActive.DeploymentID = actualEvent.DeploymentID()

		_, err := s.SetDriftMonitorInActive(reqDriftInActive)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

	case common.DeploymentAccuracyAnalyzeEnabled:
		println("AccuracyAnalyze true")
		strAssociaionID := actualEvent.AssociationID()
		reqAccuracyActive := new(appDTO.MonitorAccuracyActiveRequestDTO)
		reqAccuracyActive.DeploymentID = actualEvent.DeploymentID()
		reqAccuracyActive.ModelPackageID = actualEvent.ModelPackageID()
		reqAccuracyActive.AssociationID = &strAssociaionID
		reqAccuracyActive.AssociationIDInFeature = actualEvent.AssociationIDInFeature()
		reqAccuracyActive.CurrentModelID = actualEvent.CurrentModelID()
		//toBe..
		//reqAccuracyActive.AssociationIDInFeature = req.AssociationIDInFeature

		_, err := s.SetAccuracyMonitorActive(reqAccuracyActive)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

	case common.DeploymentAccuracyAnalyzeDisabled:
		reqAccuracyInActive := new(appDTO.MonitorAccuracyInActiveRequestDTO)
		reqAccuracyInActive.DeploymentID = actualEvent.DeploymentID()

		_, err := s.SetAccuracyMonitorInActive(reqAccuracyInActive)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

	case common.DeploymentDeleted:
		//
		reqDeleteMonitoring := &appDTO.MonitorDeleteRequestDTO{
			DeploymentID: actualEvent.DeploymentID(),
		}

		_, err := s.Delete(reqDeleteMonitoring)
		if err != nil {
			//return nil, fmt.Errorf("monitoring delete error: %s", err)
			fmt.Errorf("monitoring delete error: %s", err)
		}

	case common.DeploymentActived:
		//monitor active
		reqMonitor := &appDTO.MonitorGetByIDRequestDTO{
			ID: actualEvent.DeploymentID(),
		}

		resMonitor, err := s.GetByID(reqMonitor)

		reqServiceHealth := &appDTO.MonitorServiceHealthActiveRequestDTO{
			DeploymentID:   actualEvent.DeploymentID(),
			CurrentModelID: "",
		}
		_, err = s.SetServiceHealthMonitorActive(reqServiceHealth)
		if err != nil {
			fmt.Errorf("monitoring active error: %s", err)
		}
		if resMonitor.Monitor.FeatureDriftTracking == true {
			reqDrift := &appDTO.MonitorDriftActiveRequestDTO{
				DeploymentID:   actualEvent.DeploymentID(),
				ModelPackageID: "",
				CurrentModelID: "",
			}
			_, err = s.SetDriftMonitorActive(reqDrift)
			if err != nil {
				fmt.Errorf("monitoring active error: %s", err)
			}
		}
		if resMonitor.Monitor.AccuracyMonitoring == true {
			reqAccuracy := &appDTO.MonitorAccuracyActiveRequestDTO{
				DeploymentID:           actualEvent.DeploymentID(),
				ModelPackageID:         "",
				AssociationID:          nil,
				AssociationIDInFeature: false,
				CurrentModelID:         "",
			}
			_, err = s.SetAccuracyMonitorActive(reqAccuracy)
		}
		if err != nil {
			fmt.Errorf("monitoring active error: %s", err)
		}
	case common.DeploymentInActived:
		//monitor inactive
		reqMonitor := &appDTO.MonitorGetByIDRequestDTO{
			ID: actualEvent.DeploymentID(),
		}

		resMonitor, err := s.GetByID(reqMonitor)
		reqServiceHealth := &appDTO.MonitorServiceHealthInActiveRequestDTO{
			DeploymentID: actualEvent.DeploymentID(),
		}
		_, err = s.SetServiceHealthMonitorInActive(reqServiceHealth)
		if err != nil {
			fmt.Errorf("monitoring inactive error: %s", err)
		}
		if resMonitor.Monitor.FeatureDriftTracking == true {
			reqDrift := &appDTO.MonitorDriftInActiveRequestDTO{
				DeploymentID: actualEvent.DeploymentID(),
			}
			_, err = s.SetDriftMonitorInActive(reqDrift)
			if err != nil {
				fmt.Errorf("monitoring inactive error: %s", err)
			}
		}
		if resMonitor.Monitor.AccuracyMonitoring == true {
			reqAccuracy := &appDTO.MonitorAccuracyInActiveRequestDTO{
				DeploymentID: actualEvent.DeploymentID(),
			}
			_, err = s.SetAccuracyMonitorInActive(reqAccuracy)
		}
		if err != nil {
			fmt.Errorf("monitoring inactive error: %s", err)
		}

	default:
		return

	}
}

func (s *MonitorService) GetByIDInternal(monitoringID string) (*common.MonitorGetByIDInternalResponseDTO, error) {

	// if err := req.Validate(); err != nil {
	// 	return nil, err
	// }

	res, err := s.repo.Get(monitoringID)
	if err != nil {
		return nil, err
	}

	resDTO := new(common.MonitorGetByIDInternalResponseDTO)
	resDTO.ID = res.ID
	resDTO.ServiceHealthStatus = res.ServiceHealthStatus
	resDTO.AccuracyStatus = res.AccuracyStatus
	resDTO.DriftStatus = res.DriftStatus
	resDTO.FeatureDriftTracking = res.FeatureDriftTracking
	resDTO.AccuracyMonitoring = res.AccuracyMonitoring
	resDTO.AssociationID = res.AssociationID
	resDTO.AssociationIDInFeature = res.AssociationIDInFeature

	return resDTO, nil
}
