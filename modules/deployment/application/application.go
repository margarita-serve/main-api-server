package application

import (
	appSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

// NewDeploymentApp new DeploymentApp
func NewDeploymentApp(h *handler.Handler, deploymentSvc *appSvc.DeploymentService) (*DeploymentApp, error) {

	app := new(DeploymentApp)
	app.handler = h

	app.DeploymentSvc = deploymentSvc

	return app, nil
}

// DeploymentApp represent DDD Module: Email (Application Layer)
type DeploymentApp struct {
	handler       *handler.Handler
	DeploymentSvc *appSvc.DeploymentService
}
