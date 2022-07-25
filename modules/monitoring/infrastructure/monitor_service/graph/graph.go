package graph

import (
	conMonitorGraph "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/monitoring_service/graph"
	domSchema "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service/graph/dto"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

func NewGraphAdapter(h *handler.Handler) (*GraphAdapter, error) {
	var err error

	adp := new(GraphAdapter)
	cfg, err := h.GetConfig()
	if err != nil {
		return nil, err
	}
	config := conMonitorGraph.Config{}
	config.Endpoint = cfg.Connectors.GraphServer.Endpoint
	adp.connector = conMonitorGraph.NewGraphMonitor(config, nil)

	return adp, err
}

type GraphAdapter struct {
	connector *conMonitorGraph.GraphMonitor
}

func (a *GraphAdapter) MonitorGetDetailGraph(req *domSchema.DetailGraphGetRequest) (*domSchema.DetailGraphGetResponse, error) {
	resp := new(domSchema.DetailGraphGetResponse)

	connReq, err := MapGetDetailGraphReq(req)
	if err != nil {
		return nil, err
	}

	connResp, err := a.connector.GetFeatureDetailGraph(connReq)
	if err != nil {
		return nil, err
	}

	resp, err = MapGetDetailGraphRes(connResp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *GraphAdapter) MonitorGetDriftGraph(req *domSchema.DriftGraphGetRequest) (*domSchema.DriftGraphGetResponse, error) {
	resp := new(domSchema.DriftGraphGetResponse)

	connReq, err := MapGetDriftGraphReq(req)
	if err != nil {
		return nil, err
	}

	connResp, err := a.connector.GetDataDriftGraph(connReq)
	if err != nil {
		return nil, err
	}

	resp, err = MapGetDriftGraphRes(connResp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
