package application

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/common"
	appSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

func NewMonitorApp(h *handler.Handler, modelPackageSvc common.IModelPackageService, publisher common.EventPublisher) (*MonitorApp, error) {
	var err error

	app := new(MonitorApp)
	app.handler = h

	if app.MonitorSvc, err = appSvc.NewMonitorService(h, modelPackageSvc, publisher); err != nil {
		return nil, err
	}

	messagingService, err := appSvc.NewMessagingService(h, app.MonitorSvc)
	if err != nil {
		return nil, err
	}

	//app.MonitorSvc = monitorService

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
