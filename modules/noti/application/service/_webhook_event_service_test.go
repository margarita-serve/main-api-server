package service

import (
	"testing"

	"git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/application/dto"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/initialize"
)

func newWebHookEventSvc(t *testing.T) (*WebHookEventService, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, err
	}

	c, err := newConfig(t)
	if err != nil {
		return nil, err
	}

	h.SetConfig(c)
	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		return nil, err
	}

	webHookSvc, err := NewWebHookService(h)
	if err != nil {
		return nil, err
	}

	r, err := NewWebHookEventService(h, webHookSvc)
	if err != nil {
		return nil, err
	}

	return r, err
}

func TestWebHookEventSvc_SendEvent(t *testing.T) {
	svc, err := newWebHookEventSvc(t)
	if err != nil {
		t.Errorf("newWebHookEventSvc: %s", err.Error())
		return
	}
	// print(h.GetConfig())

	req := dto.CreateWebHookEventRequestDTO{}
	req.DeploymentID = "cbq6c77r2g4prn3pmqjg"
	req.TriggerSource = "DataDrift"

	err = svc.SendWebHookEvent(&req, svc.systemIdentity)
	if err != nil {
		t.Errorf("Send Err: %s", err.Error())
		return
	}

	if err != nil {
		t.Errorf("respJSON: %s", err.Error())
	}

}

// func TestDeploymentSvc_GetByID(t *testing.T) {
// 	svc, h, err := newWebHookEventSvc(t)
// 	if err != nil {
// 		t.Errorf("newWebHookEventSvc: %s", err.Error())
// 		return
// 	}
// 	print(h.GetConfig())

// 	req := dto.GetDeploymentRequestDTO{}
// 	req.DeploymentID = "capcdnvr2g4rignpkhqg"

// 	resp, err := svc.GetByID(&req)
// 	if err != nil {
// 		t.Errorf("GetByID: %s", err.Error())
// 		return
// 	}

// 	if resp != nil {
// 		respJSON, err := json.Marshal(resp)
// 		if err != nil {
// 			t.Errorf("respJSON: %s", err.Error())
// 		}
// 		t.Logf("Resp: %s", respJSON)
// 	}
// }

// func TestDeploymentSvc_GetList(t *testing.T) {
// 	svc, h, err := newWebHookEventSvc(t)
// 	if err != nil {
// 		t.Errorf("newWebHookEventSvc: %s", err.Error())
// 		return
// 	}
// 	print(h.GetConfig())

// 	req := dto.GetDeploymentListRequestDTO{}
// 	req.Page = 1
// 	req.Limit = 10
// 	req.Sort = ""

// 	resp, err := svc.GetList(&req)
// 	if err != nil {
// 		t.Errorf("GetList: %s", err.Error())
// 		return
// 	}

// 	if resp != nil {
// 		respJSON, err := json.Marshal(resp)
// 		if err != nil {
// 			t.Errorf("respJSON: %s", err.Error())
// 		}
// 		t.Logf("Resp: %s", respJSON)
// 	}
// }

// func TestDeploymentSvc_ReplaceModel(t *testing.T) {
// 	svc, h, err := newWebHookEventSvc(t)
// 	if err != nil {
// 		t.Errorf("newWebHookEventSvc: %s", err.Error())
// 		return
// 	}
// 	print(h.GetConfig())

// 	req := dto.ReplaceModelRequestDTO{}
// 	req.ModelPackageID = "calvv97r2g4o4gmdmre0"
// 	req.Reason = "PredictionSpeed"
// 	req.DeploymentID = "capcdnvr2g4rignpkhqg"

// 	resp, err := svc.ReplaceModel(&req)
// 	if err != nil {
// 		t.Errorf("Create: %s", err.Error())
// 		return
// 	}

// 	if resp != nil {
// 		respJSON, err := json.Marshal(resp)
// 		if err != nil {
// 			t.Errorf("respJSON: %s", err.Error())
// 		}
// 		t.Logf("Resp: %s", respJSON)
// 	}
// }

// func TestDeploymentSvc_Update(t *testing.T) {
// 	svc, h, err := newWebHookEventSvc(t)
// 	if err != nil {
// 		t.Errorf("newWebHookEventSvc: %s", err.Error())
// 		return
// 	}
// 	print(h.GetConfig())

// 	req := dto.UpdateDeploymentRequestDTO{}
// 	req.DeploymentID = "cap8nvvr2g4ptouoq6k0"
// 	req.Description = "Edited description"
// 	req.Importance = "High"
// 	req.Name = "Edited Deploy Name"

// 	resp, err := svc.UpdateDeployment(&req)
// 	if err != nil {
// 		t.Errorf("Create: %s", err.Error())
// 		return
// 	}

// 	if resp != nil {
// 		respJSON, err := json.Marshal(resp)
// 		if err != nil {
// 			t.Errorf("respJSON: %s", err.Error())
// 		}
// 		t.Logf("Resp: %s", respJSON)
// 	}
// }

// func TestDeploymentSvc_SendPrediction(t *testing.T) {
// 	svc, h, err := newWebHookEventSvc(t)
// 	if err != nil {
// 		t.Errorf("newWebHookEventSvc: %s", err.Error())
// 		return
// 	}
// 	print(h.GetConfig())

// 	req := dto.SendPredictionRequestDTO{}
// 	req.DeploymentID = "capalanr2g4qu7i76v30"
// 	req.JsonData = "{\"instances\": [[1.483887, 1.865988, 2.234620, 1.018782, -2.530891, -1.604642, 0.774676, -0.465148, -0.495225], [1.483887, 1.865988, 2.234620, 1.018782, -2.530891, -1.604642, 0.774676, -0.465148, -0.495225]]}"

// 	resp, err := svc.SendPrediction(&req)
// 	if err != nil {
// 		t.Errorf("SendPrediction: %s", err.Error())
// 		return
// 	}

// 	if resp != nil {
// 		respJSON, err := json.Marshal(resp)
// 		if err != nil {
// 			t.Errorf("respJSON: %s", err.Error())
// 		}
// 		t.Logf("Resp: %s", respJSON)
// 	}
// }
