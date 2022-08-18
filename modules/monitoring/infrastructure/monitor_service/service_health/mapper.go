package service_health

import (
	conMonitor "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/monitoring_service/service_health/types"
	domSchema "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service/service_health/dto"
)

func MapGetReq(req *domSchema.ServiceHealthGetRequest) (*conMonitor.GetServiceHealthRequest, error) {
	reqCon := new(conMonitor.GetServiceHealthRequest)
	reqCon.InferenceName = req.InferenceName
	reqCon.ModelHistoryID = req.ModelHistoryID
	reqCon.DataType = req.DataType
	reqCon.StartTime = req.StartTime
	reqCon.EndTime = req.EndTime

	return reqCon, nil
}

func MapGetRes(res *conMonitor.GetServiceHealthResponse) (*domSchema.ServiceHealthGetResponse, error) {
	resDom := new(domSchema.ServiceHealthGetResponse)
	resDom.Message = res.Message
	resDom.Data = res.Data
	resDom.StartTime = res.StartTime
	resDom.EndTIme = res.EndTime

	return resDom, nil
}

func MapCreateReq(req *domSchema.ServiceHealthCreateRequest) (*conMonitor.CreateServiceHealthRequest, error) {
	reqCon := new(conMonitor.CreateServiceHealthRequest)
	reqCon.InferenceName = req.InferenceName
	reqCon.ModelId = req.ModelHistoryID

	return reqCon, nil
}

func MapCreateRes(res *conMonitor.CreateServiceHealthResponse) (*domSchema.ServiceHealthCreateResponse, error) {
	resDom := new(domSchema.ServiceHealthCreateResponse)
	resDom.Message = res.Message
	resDom.InferenceName = res.InferenceName

	return resDom, nil
}

func MapDeleteReq(req *domSchema.ServiceHealthDeleteRequest) (*conMonitor.DisableMonitorRequest, error) {
	reqCon := new(conMonitor.DisableMonitorRequest)
	reqCon.InferenceName = req.InferenceName

	return reqCon, nil
}

func MapDeleteRes(res *conMonitor.DisableMonitorResponse) (*domSchema.ServiceHealthDeleteResponse, error) {
	resDom := new(domSchema.ServiceHealthDeleteResponse)
	resDom.Message = res.Message
	resDom.InferenceName = res.InferenceName

	return resDom, nil
}

func MapEnableReq(req *domSchema.ServiceHealthEnableRequest) (*conMonitor.EnableMonitorRequest, error) {
	reqCon := new(conMonitor.EnableMonitorRequest)
	reqCon.InferenceName = req.InferenceName

	return reqCon, nil

}

func MapEnableRes(res *conMonitor.EnableMonitorResponse) (*domSchema.ServiceHealthEnableResponse, error) {
	resDom := new(domSchema.ServiceHealthEnableResponse)
	resDom.Message = res.Message
	resDom.InferenceName = res.InferenceName

	return resDom, nil
}
