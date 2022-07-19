package application

import (
	appSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

// NewDeploymentApp new DeploymentApp
func NewDeploymentApp(h *handler.Handler, modelPackageSvc appSvc.IModelPackageService, monitorSvc appSvc.IMonitorService) (*DeploymentApp, error) {
	var err error

	app := new(DeploymentApp)
	app.handler = h

	if app.DeploymentSvc, err = appSvc.NewDeploymentService(h, modelPackageSvc, monitorSvc); err != nil {
		return nil, err
	}

	return app, nil
}

// DeploymentApp represent DDD Module: Email (Application Layer)
type DeploymentApp struct {
	handler       *handler.Handler
	DeploymentSvc *appSvc.DeploymentService
}
