package router

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/feature"
	internalMiddleware "git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/middleware"
	"github.com/labstack/echo/v4"
)

// SetProject set Project Router
func SetProject(eg *echo.Group, f *feature.FProject) {
	gc := eg.Group("/projects")
	gc.Use(internalMiddleware.JWTVerifier(f.GetHandler()))

	gc.POST("", f.Create)
	gc.GET("/:projectID", f.GetByID)
	gc.GET("", f.GetList)
	gc.DELETE("/:projectID", f.Delete)
	gc.PATCH("/:projectID", f.Update)
}
