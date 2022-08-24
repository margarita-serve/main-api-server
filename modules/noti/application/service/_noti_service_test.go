package service

import (
	"encoding/json"
	"testing"

	appAuthSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/auths/application/service"
	appDeploySvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/application/service"
	appEmailSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/application/service"
	appModelPackageSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/application/service"
	appMonitoringSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/application/dto"
	appProjectSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/project/application/service"
	appResourceEnvSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/resource/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/initialize"
)

func newNotiSvc(t *testing.T) (*NotiService, *handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, nil, err
	}

	c, err := newConfig(t)
	if err != nil {
		return nil, nil, err
	}

	h.SetConfig(c)
	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		return nil, nil, err
	}

	emailSvc, err := appEmailSvc.NewEmailService(h)
	if err != nil {
		return nil, nil, err
	}

	clusterInfoSvc, err := appResourceEnvSvc.NewClusterInfoService(h)
	if err != nil {
		return nil, nil, err
	}

	predictionEnvSvc, err := appResourceEnvSvc.NewPredictionEnvService(h, clusterInfoSvc)
	if err != nil {
		return nil, nil, err
	}

	projectSvc, err := appProjectSvc.NewProjectService(h)
	if err != nil {
		return nil, nil, err
	}

	modelPackageSvc, err := appModelPackageSvc.NewModelPackageService(h, projectSvc)
	if err != nil {
		return nil, nil, err
	}

	monitorSvc, err := appMonitoringSvc.NewMonitorService(h, modelPackageSvc)
	if err != nil {
		return nil, nil, err
	}

	deploymentSvc, err := appDeploySvc.NewDeploymentService(h, predictionEnvSvc, projectSvc, modelPackageSvc, monitorSvc)
	if err != nil {
		return nil, nil, err
	}

	authSvc, err := appAuthSvc.NewAuthenticationSvc(h)
	if err != nil {
		return nil, nil, err
	}

	r, err := NewNotiService(h, emailSvc, deploymentSvc, projectSvc, authSvc)
	if err != nil {
		return nil, nil, err
	}

	return r, h, nil
}

func TestNotiSvc_SendNoti(t *testing.T) {
	svc, h, err := newNotiSvc(t)
	if err != nil {
		t.Errorf("newNotiSvc: %s", err.Error())
		return
	}
	print(h.GetConfig())

	req := dto.NotiRequestDTO{}
	req.DeploymentID = "cbpjvgfr2g4ng38tb86g"
	req.NotiCategory = "data-drift"
	req.Data = map[string]interface{}{
		"bacon": "delicious",
	}

	resp, err := svc.SendNoti(&req, svc.systemIdentity)
	if err != nil {
		t.Errorf("Send Err: %s", err.Error())
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

// func TestDeploymentSvc_GetByID(t *testing.T) {
// 	svc, h, err := newNotiSvc(t)
// 	if err != nil {
// 		t.Errorf("newNotiSvc: %s", err.Error())
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
// 	svc, h, err := newNotiSvc(t)
// 	if err != nil {
// 		t.Errorf("newNotiSvc: %s", err.Error())
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
// 	svc, h, err := newNotiSvc(t)
// 	if err != nil {
// 		t.Errorf("newNotiSvc: %s", err.Error())
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
// 	svc, h, err := newNotiSvc(t)
// 	if err != nil {
// 		t.Errorf("newNotiSvc: %s", err.Error())
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
// 	svc, h, err := newNotiSvc(t)
// 	if err != nil {
// 		t.Errorf("newNotiSvc: %s", err.Error())
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
