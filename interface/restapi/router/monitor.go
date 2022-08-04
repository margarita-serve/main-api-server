package router

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/feature"
	internalMiddleware "git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/middleware"
	"github.com/labstack/echo/v4"
)

func SetMonitor(eg *echo.Group, f *feature.FMonitor) {
	gc := eg.Group("/deployments/:deploymentID/monitor")
	gc.Use(internalMiddleware.JWTVerifier(f.GetHandler()))

	gc.GET("", f.GetMonitorSetting)
	gc.GET("/detail", f.GetDetail)
	gc.GET("/detail/graph", f.GetFeatureDetailGraph)
	gc.GET("/drift", f.GetDrift)
	gc.PATCH("/drift", f.PatchDriftSetting)
	gc.GET("/drift/graph", f.GetDriftGraph)
	gc.POST("/actual", f.UploadActual)
	gc.GET("/accuracy", f.GetAccuracy)
	gc.PATCH("/accuracy", f.PatchAccuracySetting)
	gc.PATCH("/association-id", f.UpdateAssociationID)
	//gc.POST("", f.Create)
}
