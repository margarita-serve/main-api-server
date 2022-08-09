package restapi

import (
	feature "git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/feature"
	router "git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/router"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"net/http"
)

// Template html Template
type Template struct {
	templates *template.Template
}

// Render html template
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// SetRouters is a function to ser Echo Routers
func SetRouters(e *echo.Echo, h *handler.Handler) {
	cfg, err := h.GetConfig()

	// set default middleware
	e.Pre(middleware.RemoveTrailingSlash())
	if cfg.Applications.Servers.RestAPI.Options.Middlewares.Logger.Enable {
		e.Use(middleware.Logger())
	}
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	// html template
	// t := &Template{
	// 	templates: template.Must(template.ParseGlob("www/templates/**/*.*ml")),
	// }
	// e.Renderer = t

	// Set CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))

	// features
	features, err := feature.NewFeature(h)
	if err != nil {
		panic(err)
	}

	// System
	gs := e.Group("system")
	router.SetSystem(gs, features.System)

	// OpenApi/swagger-ui
	if cfg.Applications.Servers.RestAPI.Options.DisplayOpenAPI {
		gd := e.Group("openapi")
		router.SetOpenAPI(gd, features.OpenAPI)
	}

	//Group API
	ga := e.Group("api/v1")
	router.SetDeployment(ga, features.Deployment)
	router.SetModelPackage(ga, features.ModelPackage)
	router.SetMonitor(ga, h, features.Monitor)
	router.SetAuths(ga, features.Auths)
	router.SetEmail(ga, features.Email)
	router.SetGraph(ga, h)
}
