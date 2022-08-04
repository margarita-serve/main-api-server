package router

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/feature"
	internalMiddleware "git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/middleware"
	"github.com/labstack/echo/v4"
)

// SetDeployment Deployment Router
func SetDeployment(eg *echo.Group, f *feature.FDeployment) {
	gc := eg.Group("/deployments")
	gc.Use(internalMiddleware.JWTVerifier(f.GetHandler()))

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

}
