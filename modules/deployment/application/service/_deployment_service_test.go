package service

import (
	"encoding/json"
	"testing"

	common "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/common"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/application/dto"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/initialize"
)

type modelPackageSvc struct{}

func (s modelPackageSvc) GetByIDInternal(modelPackageID string) (*common.InternalGetModelPackageResponseDTO, error) {
	return &common.InternalGetModelPackageResponseDTO{
		ModelPackageID: modelPackageID,
	}, nil
}

type projectSvc struct{}

func (s projectSvc) GetListInternal(userID string) (*common.GetProjectListResponseDTO, error) {
	return &common.GetProjectListResponseDTO{
		Rows: "test",
	}, nil
}
func (s projectSvc) GetByIDInternal(projectID string) (*common.GetProjectInternalResponseDTO, error) {

}

func newDeploymentSvc(t *testing.T) (*DeploymentService, *handler.Handler, error) {
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

	// modelPackageSvc, err := appModelPackageSvc.NewModelPackageService(h)
	// if err != nil {
	// 	return nil, nil, err
	// }

	// monitorSvc, err := appMonitoringSvc.NewMonitorService(h, modelPackageSvc)
	// if err != nil {
	// 	return nil, nil, err
	// }

	r, err := NewDeploymentService(h, projectSvc, new(modelPackageSvc), monitorSvc, publisher)
	if err != nil {
		return nil, nil, err
	}

	return r, h, nil
}

func TestDeploymentSvc_Create(t *testing.T) {
	svc, h, err := newDeploymentSvc(t)
	if err != nil {
		t.Errorf("newDeploymentSvc: %s", err.Error())
		return
	}
	print(h.GetConfig())

	req := dto.CreateDeploymentRequestDTO{}
	req.Name = "Test Deployment"
	req.Description = "this is test"
	req.Importance = "Low"
	req.ModelPackageID = "calvv97r2g4o4gmdmre0"
	req.PredictionEnvID = "12345678901234567890"
	req.RequestCPU = 0.1
	req.RequestMEM = 0.5

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

func TestDeploymentSvc_GetByID(t *testing.T) {
	svc, h, err := newDeploymentSvc(t)
	if err != nil {
		t.Errorf("newDeploymentSvc: %s", err.Error())
		return
	}
	print(h.GetConfig())

	req := dto.GetDeploymentRequestDTO{}
	req.DeploymentID = "capcdnvr2g4rignpkhqg"

	resp, err := svc.GetByID(&req)
	if err != nil {
		t.Errorf("GetByID: %s", err.Error())
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

func TestDeploymentSvc_GetList(t *testing.T) {
	svc, h, err := newDeploymentSvc(t)
	if err != nil {
		t.Errorf("newDeploymentSvc: %s", err.Error())
		return
	}
	print(h.GetConfig())

	req := dto.GetDeploymentListRequestDTO{}
	req.Page = 1
	req.Limit = 10
	req.Sort = ""

	resp, err := svc.GetList(&req)
	if err != nil {
		t.Errorf("GetList: %s", err.Error())
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

func TestDeploymentSvc_ReplaceModel(t *testing.T) {
	svc, h, err := newDeploymentSvc(t)
	if err != nil {
		t.Errorf("newDeploymentSvc: %s", err.Error())
		return
	}
	print(h.GetConfig())

	req := dto.ReplaceModelRequestDTO{}
	req.ModelPackageID = "calvv97r2g4o4gmdmre0"
	req.Reason = "PredictionSpeed"
	req.DeploymentID = "capcdnvr2g4rignpkhqg"

	resp, err := svc.ReplaceModel(&req)
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

func TestDeploymentSvc_Update(t *testing.T) {
	svc, h, err := newDeploymentSvc(t)
	if err != nil {
		t.Errorf("newDeploymentSvc: %s", err.Error())
		return
	}
	print(h.GetConfig())

	req := dto.UpdateDeploymentRequestDTO{}
	req.DeploymentID = "cap8nvvr2g4ptouoq6k0"
	req.Description = "Edited description"
	req.Importance = "High"
	req.Name = "Edited Deploy Name"

	resp, err := svc.UpdateDeployment(&req)
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

// func TestDeploymentSvc_SendPrediction(t *testing.T) {
// 	svc, h, err := newDeploymentSvc(t)
// 	if err != nil {
// 		t.Errorf("newDeploymentSvc: %s", err.Error())
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
