package application

import (
	appSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

func NewMonitorApp(h *handler.Handler) (*MonitorApp, error) {
	var err error

	app := new(MonitorApp)
	app.handler = h

	if app.MonitorSvc, err = appSvc.NewMonitorService(h); err != nil {
		return nil, err
	}
	appSvc.NewMessagingService(h)

	return app, nil
}

type MonitorApp struct {
	handler    *handler.Handler
	MonitorSvc *appSvc.MonitorService
}
