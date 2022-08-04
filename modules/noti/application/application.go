package application

import (
	appSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

// NewNotiApp new NotiApp
func NewNotiApp(h *handler.Handler) (*NotiApp, error) {
	var err error

	app := new(NotiApp)
	app.handler = h

	if app.NotiSvc, err = appSvc.NewNotiService(h); err != nil {
		return nil, err
	}

	return app, nil
}

// NotiApp represent DDD Module:  (Application Layer)
type NotiApp struct {
	handler *handler.Handler
	NotiSvc *appSvc.NotiService
}
