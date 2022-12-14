package service

import (
	domAccuracySvcDto "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service/accuracy/dto"
	domDriftSvcDto "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service/data_drift/dto"
	domServiceHealthSvcDto "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service/service_health/dto"
)

type IExternalDriftMonitorAdapter interface {
	MonitorCreate(req *domDriftSvcDto.DataDriftCreateRequest) (*domDriftSvcDto.DataDriftCreateResponse, error)
	MonitorPatch(req *domDriftSvcDto.DataDriftPatchRequest) (*domDriftSvcDto.DataDriftPatchResponse, error)
	MonitorGetDetail(req *domDriftSvcDto.DataDriftGetRequest) (*domDriftSvcDto.DataDriftGetResponse, error)
	MonitorGetDrift(req *domDriftSvcDto.DataDriftGetRequest) (*domDriftSvcDto.DataDriftGetResponse, error)
	MonitorEnable(req *domDriftSvcDto.DataDriftEnableRequest) (*domDriftSvcDto.DataDriftEnableResponse, error)
	MonitorDisable(req *domDriftSvcDto.DataDriftDeleteRequest) error
}

type IExternalAccuracyMonitorAdapter interface {
	MonitorCreate(req *domAccuracySvcDto.AccuracyCreateRequest) (*domAccuracySvcDto.AccuracyCreateResponse, error)
	MonitorPatch(req *domAccuracySvcDto.AccuracyPatchRequest) (*domAccuracySvcDto.AccuracyPatchResponse, error)
	MonitorGetAccuracy(req *domAccuracySvcDto.AccuracyGetRequest) (*domAccuracySvcDto.AccuracyGetResponse, error)
	MonitorPostActual(req *domAccuracySvcDto.AccuracyPostActualRequest) (*domAccuracySvcDto.AccuracyPostActualResponse, error)
	MonitorDisable(req *domAccuracySvcDto.AccuracyDeleteRequest) error
	MonitorEnable(req *domAccuracySvcDto.AccuracyEnableRequest) (*domAccuracySvcDto.AccuracyEnableResponse, error)
	MonitorAssociationIDPatch(req *domAccuracySvcDto.AccuracyUpdateAssociationIDRequest) (*domAccuracySvcDto.AccuracyUpdateAssociationIDResponse, error)
}

type IExternalServiceHealthMonitorAdapter interface {
	MonitorCreate(req *domServiceHealthSvcDto.ServiceHealthCreateRequest) (*domServiceHealthSvcDto.ServiceHealthCreateResponse, error)
	MonitorGetServiceHealth(req *domServiceHealthSvcDto.ServiceHealthGetRequest) (*domServiceHealthSvcDto.ServiceHealthGetResponse, error)
	MonitorEnable(req *domServiceHealthSvcDto.ServiceHealthEnableRequest) (*domServiceHealthSvcDto.ServiceHealthEnableResponse, error)
	MonitorDisable(req *domServiceHealthSvcDto.ServiceHealthDeleteRequest) error
}
