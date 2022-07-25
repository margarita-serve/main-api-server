package graph

import (
	conMonitor "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/monitoring_service/graph/types"
	domSchema "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service/graph/dto"
)

func MapGetDetailGraphReq(req *domSchema.DetailGraphGetRequest) (*conMonitor.GetDetailGraphRequest, error) {
	reqCon := new(conMonitor.GetDetailGraphRequest)
	reqCon.InferenceName = req.InferenceName
	reqCon.StartTime = req.StartTime
	reqCon.EndTime = req.EndTime
	reqCon.ModelHistoryID = req.ModelHistoryID

	return reqCon, nil
}

func MapGetDetailGraphRes(res *conMonitor.GetDetailGraphResponse) (*domSchema.DetailGraphGetResponse, error) {
	resDom := new(domSchema.DetailGraphGetResponse)
	resDom.Script = res.Script

	return resDom, nil
}

func MapGetDriftGraphReq(req *domSchema.DriftGraphGetRequest) (*conMonitor.GetDriftGraphRequest, error) {
	reqCon := new(conMonitor.GetDriftGraphRequest)
	reqCon.InferenceName = req.InferenceName
	reqCon.ModelHistoryID = req.ModelHistoryID
	reqCon.StartTime = req.StartTime
	reqCon.EndTime = req.EndTime
	reqCon.DriftThreshold = req.DriftThreshold
	reqCon.ImportanceThreshold = req.ImportanceThreshold

	return reqCon, nil
}

func MapGetDriftGraphRes(res *conMonitor.GetDriftGraphResponse) (*domSchema.DriftGraphGetResponse, error) {
	resDom := new(domSchema.DriftGraphGetResponse)
	resDom.Script = res.Script

	return resDom, nil
}
