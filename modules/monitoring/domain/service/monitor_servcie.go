package service

import (
	domAccuracySvcDto "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service/accuracy/dto"
	domDriftSvcDto "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service/data_drift/dto"
)

type IExternalDriftMonitorAdapter interface {
	MonitorCreate(req *domDriftSvcDto.DataDriftCreateRequest) (*domDriftSvcDto.DataDriftCreateResponse, error)
	MonitorDelete(req *domDriftSvcDto.DataDriftDeleteRequest) error
	MonitorPatch(req *domDriftSvcDto.DataDriftPatchRequest) (*domDriftSvcDto.DataDriftPatchResponse, error)
	MonitorGetDetail(req *domDriftSvcDto.DataDriftGetRequest) (*domDriftSvcDto.DataDriftGetResponse, error)
	MonitorGetDrift(req *domDriftSvcDto.DataDriftGetRequest) (*domDriftSvcDto.DataDriftGetResponse, error)
}

type IExternalAccuracyMonitorAdapter interface {
	MonitorCreate(req *domAccuracySvcDto.AccuracyCreateRequest) (*domAccuracySvcDto.AccuracyCreateResponse, error)
	MonitorDelete(req *domAccuracySvcDto.AccuracyDeleteRequest) error
	MonitorPatch(req *domAccuracySvcDto.AccuracyPatchRequest) (*domAccuracySvcDto.AccuracyPatchResponse, error)
	MonitorGetAccuracy(req *domAccuracySvcDto.AccuracyGetRequest) (*domAccuracySvcDto.AccuracyGetResponse, error)
	MonitorPostActual(req *domAccuracySvcDto.AccuracyPostActualRequest) (*domAccuracySvcDto.AccuracyPostActualResponse, error)
}
