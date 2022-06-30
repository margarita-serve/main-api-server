package router

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/feature"
	"github.com/labstack/echo/v4"
	//internalMiddleware "git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/middleware"
)

// SetEmail set Email Router
func SetModelPackage(eg *echo.Group, f *feature.FModelPackage) {
	gc := eg.Group("/model-packages")
	//gc.Use(internalMiddleware.JWTVerifier(f.GetHandler()))

	gc.POST("", f.Create)
	gc.GET("/:modelpackageId", f.Get)
	gc.GET("", f.GetByName)
	gc.DELETE("/:modelpackageId", f.Delete)
	//e.GET("/deployments/{deploymentId}/serviceStats", GetServiceStats)
	//e.GET("/deployments/{deploymentId}/featureDrift", GetFeatureDrift)
	// gc.GET("/deployments/{deploymentId}/", f.Get)
	// gc.PATCH("/deployments/{deploymentId}", f.Patch)
	// gc.PATCH("/deployments/{deploymentId}/model", f.PatchModel)

	// gc.POST("/deployments/actuals/{}", f.Post)
}
