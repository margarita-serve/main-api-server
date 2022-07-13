package router

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/feature"
	"github.com/labstack/echo/v4"
)

// SetAuths set Auths Router
func SetAuths(eg *echo.Group, f *feature.FAuths) {

	gc := eg.Group("/auths")

	gc.POST("/register", f.RegisterUser)
	gc.GET("/registration/activate/:activationCode/:format", f.ActivateRegistration)
	gc.POST("/login", f.Login)
	gc.POST("/login-app", f.LoginApp)
}
