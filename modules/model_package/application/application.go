package application

import (
	appSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

// NewModelPackageApp new ModelPackageApp
func NewModelPackageApp(h *handler.Handler) (*ModelPackageApp, error) {
	var err error

	app := new(ModelPackageApp)
	app.handler = h

	if app.ModelPackageSvc, err = appSvc.NewModelPackageService(h); err != nil {
		return nil, err
	}

	return app, nil
}

// ModelPackageApp represent DDD Module: Email (Application Layer)
type ModelPackageApp struct {
	handler         *handler.Handler
	ModelPackageSvc *appSvc.ModelPackageService
}
