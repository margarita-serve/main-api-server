package application

import (
	appSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

// NewEmailApp new EmailApp
func NewEmailApp(h *handler.Handler, emailSvc *appSvc.EmailService, emailTemplateSvc *appSvc.EmailTemplateService) (*EmailApp, error) {

	app := new(EmailApp)
	app.handler = h

	app.EmailSvc = emailSvc
	app.EmailTemplateSvc = emailTemplateSvc

	return app, nil
}

// EmailApp represent DDD Module: Email (Application Layer)
type EmailApp struct {
	handler          *handler.Handler
	EmailSvc         *appSvc.EmailService
	EmailTemplateSvc *appSvc.EmailTemplateService
}
