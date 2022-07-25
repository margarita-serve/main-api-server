package application

import (
	appSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

func NewMonitorApp(h *handler.Handler, modelPackageSvc appSvc.IModelPackageService) (*MonitorApp, error) {
	var err error

	app := new(MonitorApp)
	app.handler = h

	if app.MonitorSvc, err = appSvc.NewMonitorService(h, modelPackageSvc); err != nil {
		return nil, err
	}

	if err = appSvc.NewMessagingService(h, app.MonitorSvc); err != nil {
		return nil, err
	}

	return app, nil
}

type MonitorApp struct {
	handler    *handler.Handler
	MonitorSvc *appSvc.MonitorService
}
