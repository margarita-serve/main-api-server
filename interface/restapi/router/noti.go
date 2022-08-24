package router

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/feature"
	internalMiddleware "git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/middleware"
	"github.com/labstack/echo/v4"
)

// SetAuths set Auths Router
func SetNoti(eg *echo.Group, f *feature.FNoti) {

	gc := eg.Group("/deployments/:deploymentID/noti")
	gc.Use(internalMiddleware.JWTVerifier(f.GetHandler()))

	gc.POST("/web-hooks", f.CreateWebHook)
	gc.DELETE("/web-hooks/:webHookID", f.DeleteWebHook)
	gc.PATCH("/web-hooks/:webHookID", f.UpdateWebHook)
	gc.GET("/web-hooks/:webHookID", f.GetByIDWebHook)
	gc.GET("/web-hooks", f.GetListWebHook)
	gc.POST("/web-hooks/test", f.TestWebHookSend)

}
