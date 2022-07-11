package router

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/feature"
	"github.com/labstack/echo/v4"
	//internalMiddleware "git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/middleware"
)

// SetEmail set Email Router
func SetModelPackage(eg *echo.Group, f *feature.FModelPackage) {
	gc := eg.Group("/projects/:projectID/model-packages")
	//gc.Use(internalMiddleware.JWTVerifier(f.GetHandler()))

	gc.POST("", f.Create)
	gc.GET("/:modelPackageID", f.GetByID)
	gc.GET("", f.GetList)
	gc.DELETE("/:modelPackageID", f.Delete)
	gc.PATCH("/:modelPackageID", f.Update)
	gc.PUT("/:modelPackageID/archive", f.Archive)
	gc.POST("/:modelPackageID/upload-model", f.UploadModel)
	gc.POST("/:modelPackageID/upload-training-dataset", f.UploadTrainingDataset)
	gc.POST("/:modelPackageID/upload-holdout-dataset", f.UploadHoldoutDataset)
	gc.GET("/:modelPackageID/download-model", f.DownloadModelFile)
	gc.GET("/:modelPackageID/download-training-dataset", f.DownloadTrainingDataset)
	gc.GET("/:modelPackageID/download-holdout-dataset", f.DownloadHoldoutDataset)

}
