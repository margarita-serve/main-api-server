package router

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/feature"
	"github.com/labstack/echo/v4"
	//internalMiddleware "git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/middleware"
)

// SetDeployment Deployment Router
func SetDeployment(eg *echo.Group, f *feature.FDeployment) {
	gc := eg.Group("/projects/:projectID/deployments")
	//gc.Use(internalMiddleware.JWTVerifier(f.GetHandler()))

	gc.POST("", f.Create)
	gc.GET("/:deploymentID", f.GetByID)
	gc.GET("", f.GetList)
	gc.DELETE("/:deploymentID", f.Delete)
	gc.PATCH("/:deploymentID/replace-model", f.ReplaceModel)
	gc.PATCH("/:deploymentID", f.Update)
	gc.PUT("/:deploymentID/active", f.Active)
	gc.PUT("/:deploymentID/inactive", f.InActive)
	//gc.GET("/:deploymentID/prediction-url", f.GetPredictionURL)
	gc.POST("/:deploymentID/predict", f.SendPrediction)
	gc.GET("/:deploymentID/governance-log", f.GetGovernanceHistory)
	gc.GET("/:deploymentID/model-history", f.GetModelHistory)

	gc2 := eg.Group("/deployments")
	//gc.Use(internalMiddleware.JWTVerifier(f.GetHandler()))

	gc2.GET("/:deploymentID", f.GetByID)
	gc2.GET("", f.GetList)
	gc2.DELETE("/:deploymentID", f.Delete)
	gc2.PATCH("/:deploymentID/replace-model", f.ReplaceModel)
	gc2.PATCH("/:deploymentID", f.Update)
	gc2.PUT("/:deploymentID/active", f.Active)
	gc2.PUT("/:deploymentID/inactive", f.InActive)
	gc2.POST("/:deploymentID/predict", f.SendPrediction)
	gc2.GET("/:deploymentID/governance-log", f.GetGovernanceHistory)
	gc2.GET("/:deploymentID/model-history", f.GetModelHistory)

}
