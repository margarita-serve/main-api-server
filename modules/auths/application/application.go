package application

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/auths/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

// NewAuthsApp new AuthsApp
func NewAuthsApp(h *handler.Handler, appAuthsSvc *service.AuthenticationSvc) (*AuthsApp, error) {

	app := new(AuthsApp)
	app.handler = h

	app.AuthenticationSvc = appAuthsSvc

	return app, nil
}

// AuthsApp type
type AuthsApp struct {
	handler           *handler.Handler
	AuthenticationSvc *service.AuthenticationSvc
}
