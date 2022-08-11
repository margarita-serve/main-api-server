package feature

import (
	docs "git.k3.acornsoft.io/msit-auto-ml/koreserv/docs"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// NewOpenAPI new FOpenAPI
func NewOpenAPI(h *handler.Handler) (*FOpenAPI, error) {

	f := new(FOpenAPI)
	f.handler = h

	return f, nil
}

// FOpenAPI represent FOpenAPI
type FOpenAPI struct {
	BaseFeature
}

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// GenOpenAPI generate openapi definition
func (f *FOpenAPI) GenOpenAPI(c echo.Context) error {
	cfg, err := f.handler.GetConfig()
	if err != nil {
		return err
	}

	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	docs.SwaggerInfo.Title = cfg.Applications.Servers.RestAPI.Options.OpenAPIDefinition.Info.Title
	docs.SwaggerInfo.Description = cfg.Applications.Servers.RestAPI.Options.OpenAPIDefinition.Info.Description
	docs.SwaggerInfo.Version = cfg.Applications.Servers.RestAPI.Options.OpenAPIDefinition.Info.Version
	docs.SwaggerInfo.Host = c.Request().Host
	docs.SwaggerInfo.BasePath = cfg.Applications.Servers.RestAPI.Options.OpenAPIDefinition.Info.BasePath

	return echoSwagger.WrapHandler(c)

}
