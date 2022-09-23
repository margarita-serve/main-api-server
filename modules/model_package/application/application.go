package application

import (
	common "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/common"
	appSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

// NewModelPackageApp new ModelPackageApp
func NewModelPackageApp(h *handler.Handler, projectSvc common.IProjectService) (*ModelPackageApp, error) {
	var err error

	app := new(ModelPackageApp)
	app.handler = h

	if app.ModelPackageSvc, err = appSvc.NewModelPackageService(h, projectSvc); err != nil {
		return nil, err
	}
	//app.ModelPackageSvc = modelPackageSvc
	return app, nil
}

// ModelPackageApp represent DDD Module: Email (Application Layer)
type ModelPackageApp struct {
	handler         *handler.Handler
	ModelPackageSvc *appSvc.ModelPackageService
}
