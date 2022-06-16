package router

import (
	_ "git.k3.acornsoft.io/msit-auto-ml/koreserv/docs"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/feature"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// SetOpenAPI set OpenAPI Router
func SetOpenAPI(eg *echo.Group, f *feature.FOpenAPI) {
	// @title KoreServ Swagger API
	// @version 1.0
	// @BasePath /api/v1
	eg.GET("/swagger/*", echoSwagger.WrapHandler)
}
