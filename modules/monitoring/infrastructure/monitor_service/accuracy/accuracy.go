package accuracy

import (
	"fmt"
	conMonitorAccuracy "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/monitoring_service/accuracy"
	domSchema "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service/accuracy/dto"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

func NewAccuracyAdapter(h *handler.Handler) (*AccuracyAdapter, error) {
	var err error

	adp := new(AccuracyAdapter)
	cfg, err := h.GetConfig()
	if err != nil {
		return nil, err
	}
	config := conMonitorAccuracy.Config{}
	config.Endpoint = cfg.Connectors.AccuracyServer.Endpoint
	adp.connector = conMonitorAccuracy.NewAccuracyMonitor(config, nil)

	return adp, err
}

type AccuracyAdapter struct {
	connector *conMonitorAccuracy.AccuracyMonitor
}

func (a *AccuracyAdapter) MonitorCreate(req *domSchema.AccuracyCreateRequest) (*domSchema.AccuracyCreateResponse, error) {
	resp := new(domSchema.AccuracyCreateResponse)

	connReq, err := MapCreateReq(req)
	if err != nil {
		return nil, err
	}

	connResp, err := a.connector.CreateAccuracyMonitor(connReq)
	if err != nil {
		return nil, err
	}

	resp, err = MapCreateRes(connResp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AccuracyAdapter) MonitorDisable(req *domSchema.AccuracyDeleteRequest) error {
	connReq, err := MapDeleteReq(req)
	if err != nil {
		return err
	}
	connResp, err := a.connector.DisableMonitor(connReq)
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

func (a *AccuracyAdapter) MonitorPatch(req *domSchema.AccuracyPatchRequest) (*domSchema.AccuracyPatchResponse, error) {
	resp := new(domSchema.AccuracyPatchResponse)

	connReq, err := MapPatchReq(req)
	if err != nil {
		return nil, err
	}

	connResp, err := a.connector.PatchAccuracyMonitor(connReq)
	if err != nil {
		return nil, err
	}

	resp, err = MapPatchRes(connResp)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

func (a *AccuracyAdapter) MonitorGetAccuracy(req *domSchema.AccuracyGetRequest) (*domSchema.AccuracyGetResponse, error) {
	resp := new(domSchema.AccuracyGetResponse)

	connReq, err := MapGetReq(req)
	if err != nil {
		return nil, err
	}

	connResp, err := a.connector.GetAccuracy(connReq)
	if err != nil {
		return nil, err
	}

	resp, err = MapGetRes(connResp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AccuracyAdapter) MonitorPostActual(req *domSchema.AccuracyPostActualRequest) (*domSchema.AccuracyPostActualResponse, error) {
	resp := new(domSchema.AccuracyPostActualResponse)
	connReq, err := MapPostReq(req)
	if err != nil {
		return nil, err
	}

	connResp, err := a.connector.PostActual(connReq)
	if err != nil {
		return nil, err
	}

	resp, err = MapPostRes(connResp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AccuracyAdapter) MonitorEnable(req *domSchema.AccuracyEnableRequest) (*domSchema.AccuracyEnableResponse, error) {
	resp := new(domSchema.AccuracyEnableResponse)
	connReq, err := MapEnableReq(req)
	if err != nil {
		return nil, err
	}

	connResp, err := a.connector.EnableMonitor(connReq)
	if err != nil {
		return nil, err
	}

	resp, err = MapEnableRes(connResp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AccuracyAdapter) MonitorAssociationIDPatch(req *domSchema.AccuracyUpdateAssociationIDRequest) (*domSchema.AccuracyUpdateAssociationIDResponse, error) {
	resp := new(domSchema.AccuracyUpdateAssociationIDResponse)
	connReq, err := MapUpdateReq(req)
	if err != nil {
		return nil, err
	}

	connResp, err := a.connector.UpdateAssociationID(connReq)
	if err != nil {
		return nil, err
	}

	resp, err = MapUpdateRes(connResp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
