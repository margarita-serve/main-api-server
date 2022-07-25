package service

import (
	"encoding/json"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/application/dto"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/initialize"
	"testing"
)

func newMonitorSvc(t *testing.T) (*MonitorService, *handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		t.Errorf("eeeeeeeeeeee")
		return nil, nil, err
	}

	c, err := newConfig(t)
	if err != nil {
		t.Errorf("error!!!!!!!!!!!!!!")
		return nil, nil, err
	}

	h.SetConfig(c)
	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		return nil, nil, err
	}

	r, err := NewMonitorService(h, nil)
	if err != nil {
		return nil, nil, err
	}
	return r, h, nil
}

func TestMonitorService_Create(t *testing.T) {
	svc, h, err := newMonitorSvc(t)
	if err != nil {
		t.Errorf("newMonitorSvc: %s", err.Error())
		return
	}
	print(h.GetConfig())

	req := dto.MonitorCreateRequestDTO{}
	req.DeploymentID = "Test Deployment"
	req.ModelPackageID = "calvv97r2g4o4gmdmre0"
	req.FeatureDriftTracking = true
	req.AccuracyMonitoring = false

	resp, err := svc.Create(&req)
	if err != nil {
		t.Errorf("Create: %s", err.Error())
		return
	}
	if resp != nil {
		respJSON, err := json.Marshal(resp)
		if err != nil {
			t.Errorf("respJSON: %s", err.Error())
		}
		t.Logf("Resp: %s", respJSON)
	}
}
