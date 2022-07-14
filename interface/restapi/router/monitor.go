package router

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/feature"
	"github.com/labstack/echo/v4"
)

func SetMonitor(eg *echo.Group, f *feature.FMonitor) {
	gc := eg.Group("/projects/:projectID/deployments/:deploymentID/monitor")

	gc.GET("/detail", f.GetDetail)
	gc.GET("/drift", f.GetDrift)
	gc.GET("/accuracy", f.GetAccuracy)
	gc.PATCH("", f.PatchMonitor)
	gc.GET("", f.GetMonitorSetting)
	gc.POST("/actual", f.UploadActual)
	gc.POST("", f.Create)
}
