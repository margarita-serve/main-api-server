package application

import (
	appSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/resource/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

// NewResourceApp new ResourceApp
func NewResourceApp(h *handler.Handler) (*ResourceApp, error) {
	var err error

	app := new(ResourceApp)
	app.handler = h

	if app.ClusterInfoSvc, err = appSvc.NewClusterInfoService(h); err != nil {
		return nil, err
	}

	if app.PredictionEnvSvc, err = appSvc.NewPredictionEnvService(h, app.ClusterInfoSvc); err != nil {
		return nil, err
	}

	return app, nil
}

// ResourceApp represent DDD Module:  (Application Layer)
type ResourceApp struct {
	handler          *handler.Handler
	ClusterInfoSvc   *appSvc.ClusterInfoService
	PredictionEnvSvc *appSvc.PredictionEnvService
}
