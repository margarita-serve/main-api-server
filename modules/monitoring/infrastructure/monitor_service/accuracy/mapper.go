package accuracy

import (
	conMonitor "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/monitoring_service/accuracy/types"
	domSchema "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service/accuracy/dto"
)

func MapGetReq(req *domSchema.AccuracyGetRequest) (*conMonitor.GetAccuracyRequest, error) {
	reqCon := new(conMonitor.GetAccuracyRequest)
	reqCon.InferenceName = req.InferenceName
	reqCon.ModelHistoryID = req.ModelHistoryID
	reqCon.DataType = req.DataType
	reqCon.StartTime = req.StartTime
	reqCon.EndTime = req.EndTime

	return reqCon, nil
}

func MapGetRes(res *conMonitor.GetAccuracyResponse) (*domSchema.AccuracyGetResponse, error) {
	resDom := new(domSchema.AccuracyGetResponse)
	resDom.Message = res.Message
	resDom.Data = res.Data
	resDom.StartTime = res.StartTime
	resDom.EndTIme = res.EndTime

	return resDom, nil
}

func MapCreateReq(req *domSchema.AccuracyCreateRequest) (*conMonitor.CreateAccuracyRequest, error) {
	reqCon := new(conMonitor.CreateAccuracyRequest)
	reqCon.InferenceName = req.InferenceName
	reqCon.ModelId = req.ModelHistoryID
	reqCon.DatasetPath = req.DatasetPath
	reqCon.ModelPath = req.ModelPath
	reqCon.TargetLabel = req.TargetLabel
	reqCon.AssociationId = req.AssociationID
	reqCon.ModelType = req.ModelType
	reqCon.Framework = req.Framework
	reqCon.DriftMetric = req.DriftMetric
	reqCon.DriftMeasurement = req.DriftMeasurement
	reqCon.AtriskValue = req.AtriskValue
	reqCon.FailingValue = req.FailingValue
	reqCon.PositiveClass = req.PositiveClass
	reqCon.NegativeClass = req.NegativeClass
	reqCon.BinaryThreshold = req.BinaryThreshold

	return reqCon, nil
}

func MapCreateRes(res *conMonitor.CreateAccuracyResponse) (*domSchema.AccuracyCreateResponse, error) {
	resDom := new(domSchema.AccuracyCreateResponse)
	resDom.Message = res.Message
	resDom.InferenceName = res.InferenceName

	return resDom, nil
}

func MapDeleteReq(req *domSchema.AccuracyDeleteRequest) (*conMonitor.DisableMonitorRequest, error) {
	reqCon := new(conMonitor.DisableMonitorRequest)
	reqCon.InferenceName = req.InferenceName

	return reqCon, nil
}

func MapDeleteRes(res *conMonitor.DisableMonitorResponse) (*domSchema.AccuracyDeleteResponse, error) {
	resDom := new(domSchema.AccuracyDeleteResponse)
	resDom.Message = res.Message
	resDom.InferenceName = res.InferenceName

	return resDom, nil
}

func MapPatchReq(req *domSchema.AccuracyPatchRequest) (*conMonitor.PatchAccuracySettingRequest, error) {
	reqCon := new(conMonitor.PatchAccuracySettingRequest)
	reqCon.InferenceName = req.InferenceName
	reqCon.DriftMetric = req.DriftMetric
	reqCon.DriftMeasurement = req.DriftMeasurement
	reqCon.AtriskValue = req.AtriskValue
	reqCon.FailingValue = req.FailingValue

	return reqCon, nil
}

func MapPatchRes(res *conMonitor.PatchAccuracySettingResponse) (*domSchema.AccuracyPatchResponse, error) {
	resDom := new(domSchema.AccuracyPatchResponse)
	resDom.Message = res.Message
	resDom.InferenceName = res.InferenceName

	return resDom, nil
}

func MapPostReq(req *domSchema.AccuracyPostActualRequest) (*conMonitor.ActualRequest, error) {
	reqCon := new(conMonitor.ActualRequest)
	reqCon.InferenceName = req.InferenceName
	reqCon.DatasetPath = req.DatasetPath
	reqCon.ActualResponse = req.ActualResponse
	reqCon.AssociationColumn = req.AssociationColumn

	return reqCon, nil
}

func MapPostRes(res *conMonitor.ActualResponse) (*domSchema.AccuracyPostActualResponse, error) {
	resDom := new(domSchema.AccuracyPostActualResponse)
	resDom.Message = res.Message
	resDom.InferenceName = res.InferenceName

	return resDom, nil
}

func MapEnableReq(req *domSchema.AccuracyEnableRequest) (*conMonitor.EnableMonitorRequest, error) {
	reqCon := new(conMonitor.EnableMonitorRequest)
	reqCon.InferenceName = req.InferenceName

	return reqCon, nil

}

func MapEnableRes(res *conMonitor.EnableMonitorResponse) (*domSchema.AccuracyEnableResponse, error) {
	resDom := new(domSchema.AccuracyEnableResponse)
	resDom.Message = res.Message
	resDom.InferenceName = res.InferenceName

	return resDom, nil
}

func MapUpdateReq(req *domSchema.AccuracyUpdateAssociationIDRequest) (*conMonitor.UpdateAssociationIDRequest, error) {
	reqCon := new(conMonitor.UpdateAssociationIDRequest)
	reqCon.InferenceName = req.InferenceName
	reqCon.AssociationID = req.AssociationID

	return reqCon, nil
}

func MapUpdateRes(res *conMonitor.UpdateAssociationIDResponse) (*domSchema.AccuracyUpdateAssociationIDResponse, error) {
	resDom := new(domSchema.AccuracyUpdateAssociationIDResponse)
	resDom.Message = res.Message
	resDom.InferenceName = res.InferenceName

	return resDom, nil
}
