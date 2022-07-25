package data_drift

import (
	conMonitor "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/monitoring_service/data_drift/types"
	domSchema "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service/data_drift/dto"
)

func MapGetReq(req *domSchema.DataDriftGetRequest) (*conMonitor.GetFeatureDriftRequest, error) {
	reqCon := new(conMonitor.GetFeatureDriftRequest)
	reqCon.InferenceName = req.InferenceName
	reqCon.ModelHistoryID = req.ModelHistoryID
	reqCon.StartTime = req.StartTime
	reqCon.EndTime = req.EndTime

	return reqCon, nil
}

func MapGetRes(res *conMonitor.GetFeatureDriftResponse) (*domSchema.DataDriftGetResponse, error) {
	resDom := new(domSchema.DataDriftGetResponse)
	resDom.Message = res.Message
	resDom.Data = res.Data
	resDom.StartTime = res.StartTime
	resDom.EndTime = res.EndTime
	resDom.PredictionCount = res.PredictionCount

	return resDom, nil
}

func MapCreateReq(req *domSchema.DataDriftCreateRequest) (*conMonitor.CreateDataDriftRequest, error) {
	reqCon := new(conMonitor.CreateDataDriftRequest)
	reqCon.TrainDatasetPath = req.TrainDatasetPath
	reqCon.ModelPath = req.ModelPath
	reqCon.InferenceName = req.InferenceName
	reqCon.ModelID = req.ModelHistoryID
	reqCon.TargetLabel = req.TargetLabel
	reqCon.ModelType = req.ModelType
	reqCon.Framework = req.Framework
	reqCon.DriftThreshold = req.DriftThreshold
	reqCon.ImportanceThreshold = req.ImportanceThreshold
	reqCon.MonitorRange = req.MonitorRange
	reqCon.LowImpAtRiskCount = req.LowImportanceAtRiskCount
	reqCon.LowImpFailingCount = req.LowImportanceFailingCount
	reqCon.HighImpAtRiskCount = req.HighImportanceAtRiskCount
	reqCon.HighImpFailingCount = req.HighImportanceFailingCount

	return reqCon, nil
}

func MapCreateRes(res *conMonitor.CreateDataDriftResponse) (*domSchema.DataDriftCreateResponse, error) {
	resDom := new(domSchema.DataDriftCreateResponse)
	resDom.Message = res.Message
	resDom.InferenceName = res.InferenceName

	return resDom, nil
}

func MapDeleteReq(req *domSchema.DataDriftDeleteRequest) (*conMonitor.DisableMonitorRequest, error) {
	reqCon := new(conMonitor.DisableMonitorRequest)
	reqCon.InferenceName = req.InferenceName

	return reqCon, nil
}

func MapDeleteRes(res *conMonitor.DisableMonitorResponse) (*domSchema.DataDriftDeleteResponse, error) {
	resDom := new(domSchema.DataDriftDeleteResponse)
	resDom.Message = res.Message
	resDom.InferenceName = res.InferenceName

	return resDom, nil
}

func MapPatchReq(req *domSchema.DataDriftPatchRequest) (*conMonitor.PatchDriftMonitorSettingRequest, error) {
	reqCon := new(conMonitor.PatchDriftMonitorSettingRequest)
	reqCon.InferenceName = req.InferenceName
	reqCon.DriftThreshold = req.DriftThreshold
	reqCon.ImportanceThreshold = req.ImportanceThreshold
	reqCon.MonitorRange = req.MonitorRange
	reqCon.LowImpAtRiskCount = req.LowImportanceAtRiskCount
	reqCon.LowImpFailingCount = req.LowImportanceFailingCount
	reqCon.HighImpAtRiskCount = req.HighImportanceAtRiskCount
	reqCon.HighImpFailingCount = req.HighImportanceFailingCount

	return reqCon, nil
}

func MapPatchRes(res *conMonitor.PatchDriftMonitorSettingResponse) (*domSchema.DataDriftPatchResponse, error) {
	resDom := new(domSchema.DataDriftPatchResponse)
	resDom.Message = res.Message
	resDom.InferenceName = res.InferenceName

	return resDom, nil
}

func MapEnableReq(req *domSchema.DataDriftEnableRequest) (*conMonitor.EnableMonitorRequest, error) {
	reqCon := new(conMonitor.EnableMonitorRequest)
	reqCon.InferenceName = req.InferenceName

	return reqCon, nil
}

func MapEnableRes(res *conMonitor.EnableMonitorResponse) (*domSchema.DataDriftEnableResponse, error) {
	resDom := new(domSchema.DataDriftEnableResponse)
	resDom.Message = res.Message
	resDom.InferenceName = res.InferenceName

	return resDom, nil
}
