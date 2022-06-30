package router

import (
	_ "git.k3.acornsoft.io/msit-auto-ml/koreserv/docs"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/feature"
	"github.com/labstack/echo/v4"
)

// SetOpenAPI set OpenAPI Router
func SetOpenAPI(eg *echo.Group, f *feature.FOpenAPI) {
	eg.GET("/swagger/*", f.GenOpenAPI)
}
