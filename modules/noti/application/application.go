package application

import (
	appSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

// NewNotiApp new NotiApp
func NewNotiApp(h *handler.Handler, notiService *appSvc.NotiService, webHookService *appSvc.WebHookService) (*NotiApp, error) {
	app := new(NotiApp)
	app.handler = h

	app.NotiSvc = notiService
	app.WebHookSvc = webHookService

	return app, nil
}

// NotiApp represent DDD Module:  (Application Layer)
type NotiApp struct {
	handler    *handler.Handler
	NotiSvc    *appSvc.NotiService
	WebHookSvc *appSvc.WebHookService
}
