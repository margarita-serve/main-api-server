package application

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/common"
	appSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

// NewDeploymentApp new DeploymentApp
func NewDeploymentApp(h *handler.Handler, ProjectSvc common.IProjectService, ModelPackageSvc common.IModelPackageService, MonitorSvc common.IMonitorService, publisher common.EventPublisher) (*DeploymentApp, error) {
	var err error

	app := new(DeploymentApp)
	app.handler = h

	if app.DeploymentSvc, err = appSvc.NewDeploymentService(h, ProjectSvc, ModelPackageSvc, MonitorSvc, publisher); err != nil {
		return nil, err
	}

	return app, err
}

// DeploymentApp represent DDD Module: Email (Application Layer)
type DeploymentApp struct {
	handler       *handler.Handler
	DeploymentSvc *appSvc.DeploymentService
}
