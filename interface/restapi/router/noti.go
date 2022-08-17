package router

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/feature"
	"github.com/labstack/echo/v4"
)

// SetAuths set Auths Router
func SetNoti(eg *echo.Group, f *feature.FNoti) {

	gc := eg.Group("/noti")

	gc.POST("/register", f.Create)

}
