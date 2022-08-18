package service_health

import (
	"fmt"
	conMonitorServiceHealth "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/monitoring_service/service_health"
	domSchema "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/service/service_health/dto"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

func NewServiceHealthAdapter(h *handler.Handler) (*ServiceHealthAdapter, error) {
	var err error

	adp := new(ServiceHealthAdapter)
	cfg, err := h.GetConfig()
	if err != nil {
		return nil, err
	}
	config := conMonitorServiceHealth.Config{}
	config.Endpoint = cfg.Connectors.ServiceHealthServer.Endpoint
	adp.connector = conMonitorServiceHealth.NewServiceHealthMonitor(config, nil)

	return adp, err
}

type ServiceHealthAdapter struct {
	connector *conMonitorServiceHealth.ServiceHealthMonitor
}

func (a *ServiceHealthAdapter) MonitorCreate(req *domSchema.ServiceHealthCreateRequest) (*domSchema.ServiceHealthCreateResponse, error) {
	resp := new(domSchema.ServiceHealthCreateResponse)

	connReq, err := MapCreateReq(req)
	if err != nil {
		return nil, err
	}

	connResp, err := a.connector.CreateServiceHealthMonitor(connReq)
	if err != nil {
		return nil, err
	}

	resp, err = MapCreateRes(connResp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *ServiceHealthAdapter) MonitorDisable(req *domSchema.ServiceHealthDeleteRequest) error {
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

func (a *ServiceHealthAdapter) MonitorGetServiceHealth(req *domSchema.ServiceHealthGetRequest) (*domSchema.ServiceHealthGetResponse, error) {
	resp := new(domSchema.ServiceHealthGetResponse)

	connReq, err := MapGetReq(req)
	if err != nil {
		return nil, err
	}

	connResp, err := a.connector.GetServiceHealth(connReq)
	if err != nil {
		return nil, err
	}

	resp, err = MapGetRes(connResp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (a *ServiceHealthAdapter) MonitorEnable(req *domSchema.ServiceHealthEnableRequest) (*domSchema.ServiceHealthEnableResponse, error) {
	resp := new(domSchema.ServiceHealthEnableResponse)
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
