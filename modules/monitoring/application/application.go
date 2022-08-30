package application

import (
	appSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

func NewMonitorApp(h *handler.Handler, monitorService *appSvc.MonitorService, messagingService *appSvc.MessagingService) (*MonitorApp, error) {
	var err error

	app := new(MonitorApp)
	app.handler = h

	app.MonitorSvc = monitorService

	err = messagingService.MessageConsume()
	if err != nil {
		return nil, err
	}

	return app, nil
}

type MonitorApp struct {
	handler    *handler.Handler
	MonitorSvc *appSvc.MonitorService
}
