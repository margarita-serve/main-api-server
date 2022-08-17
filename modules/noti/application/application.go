package application

import (
	appSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

// NewNotiApp new NotiApp
func NewNotiApp(h *handler.Handler, emailSvc appSvc.IEmailService, deploymentSvc appSvc.IDeploymentService, projectSvc appSvc.IProjectService, authSvc appSvc.IAuthService, governanceHistorySvc appSvc.IGovernanceHistoryService) (*NotiApp, error) {
	var err error

	app := new(NotiApp)
	app.handler = h

	if app.NotiSvc, err = appSvc.NewNotiService(h, emailSvc, deploymentSvc, projectSvc, authSvc, governanceHistorySvc); err != nil {
		return nil, err
	}

	return app, nil
}

// NotiApp represent DDD Module:  (Application Layer)
type NotiApp struct {
	handler *handler.Handler
	NotiSvc *appSvc.NotiService
}
