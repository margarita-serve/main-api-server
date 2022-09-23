package application

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/common"
	appSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

// NewNotiApp new NotiApp
func NewNotiApp(h *handler.Handler, EmailSvc common.IEmailService, DeploymentSvc common.IDeploymentService, ProjectSvc common.IProjectService, AuthenticationSvc common.IAuthService) (*NotiApp, error) {
	app := new(NotiApp)
	app.handler = h

	WebHookEventSvc, err := appSvc.NewWebHookEventService(h)
	if err != nil {
		return nil, err
	}

	WebHookSvc, err := appSvc.NewWebHookService(h, WebHookEventSvc)
	if err != nil {
		return nil, err
	}

	NotiSvc, err := appSvc.NewNotiService(h, EmailSvc, DeploymentSvc, ProjectSvc, AuthenticationSvc, WebHookSvc)
	if err != nil {
		return nil, err
	}

	app.NotiSvc = NotiSvc
	app.WebHookSvc = WebHookSvc

	return app, nil
}

// NotiApp represent DDD Module:  (Application Layer)
type NotiApp struct {
	handler    *handler.Handler
	NotiSvc    *appSvc.NotiService
	WebHookSvc *appSvc.WebHookService
}
