package data_drift

import (
	"fmt"

	conMonitorDataDrift "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/monitoring_service/data_drift"
	domSchema "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring_mockup/domain/service/data_drift/dto"
)

func NewDataDriftAdapter() (*DataDriftAdapter, error) {
	var err error

	adp := new(DataDriftAdapter)
	config := conMonitorDataDrift.Config{}
	adp.connector = conMonitorDataDrift.NewDriftMonitor(config, nil)

	return adp, err
}

type DataDriftAdapter struct {
	connector *conMonitorDataDrift.DriftMonitor
}

func (a *DataDriftAdapter) MonitorCreate(req *domSchema.DataDriftCreateRequest) (*domSchema.DataDriftCreateResponse, error) {
	resp := new(domSchema.DataDriftCreateResponse)

	connReq, err := MapCreateReq(req)
	if err != nil {
		return nil, err
	}

	connResp, err := a.connector.CreateDriftMonitor(connReq)
	if err != nil {
		return nil, err
	}

	resp, err = MapCreateRes(connResp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *DataDriftAdapter) MonitorDelete(req *domSchema.DataDriftDeleteRequest) error {
	connReq, err := MapDeleteReq(req)
	if err != nil {
		return err
	}
	connResp, err := a.connector.DeleteDriftMonitor(connReq)
	if err != nil {
		return err
	}

	resp, err := MapDeleteRes(connResp)
	if err != nil {
		fmt.Printf("resp: %v\n", resp)
		return err
	}
	return nil
}

func (a *DataDriftAdapter) MonitorPatch(req *domSchema.DataDriftPatchRequest) (*domSchema.DataDriftPatchResponse, error) {
	resp := new(domSchema.DataDriftPatchResponse)

	connReq, err := MapPatchReq(req)
	if err != nil {
		return nil, err
	}

	connResp, err := a.connector.PatchDriftMonitor(connReq)
	if err != nil {
		return nil, err
	}

	resp, err = MapPatchRes(connResp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *DataDriftAdapter) MonitorGetDrift(req *domSchema.DataDriftGetRequest) (*domSchema.DataDriftGetResponse, error) {
	resp := new(domSchema.DataDriftGetResponse)
	connReq, err := MapGetReq(req)
	if err != nil {
		return nil, err
	}

	connResp, err := a.connector.GetFeatureDrift(connReq)
	if err != nil {
		return nil, err
	}

	resp, err = MapGetRes(connResp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *DataDriftAdapter) MonitorGetDetail(req *domSchema.DataDriftGetRequest) (*domSchema.DataDriftGetResponse, error) {
	resp := new(domSchema.DataDriftGetResponse)
	connReq, err := MapGetReq(req)
	if err != nil {
		return nil, err
	}

	connResp, err := a.connector.GetFeatureDetail(connReq)
	if err != nil {
		return nil, err
	}

	resp, err = MapGetRes(connResp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
