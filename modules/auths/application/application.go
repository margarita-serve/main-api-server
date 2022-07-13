package application

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/auths/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

// NewAuthsApp new AuthsApp
func NewAuthsApp(h *handler.Handler) (*AuthsApp, error) {
	var err error

	app := new(AuthsApp)
	app.handler = h

	if app.AuthenticationSvc, err = service.NewAuthenticationSvc(h); err != nil {
		return nil, err
	}

	return app, nil
}

// AuthsApp type
type AuthsApp struct {
	handler           *handler.Handler
	AuthenticationSvc *service.AuthenticationSvc
}
