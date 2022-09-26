package router

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"net/url"
)

// SetGraph Graph Router
func SetGraph(eg *echo.Group, h *handler.Handler) {
	cfg, _ := h.GetConfig()
	ProxyServer := cfg.Connectors.GraphServer.Endpoint

	gc := eg.Group("/deployments/graph-svr", ACAOHeaderOverwriteMiddleware)

	url1, _ := url.Parse(ProxyServer)
	targets := []*middleware.ProxyTarget{
		{
			URL: url1,
		},
	}

	gc.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))

	ProxyServerResource := cfg.Connectors.ServiceHealthServer.Endpoint

	gc2 := eg.Group("/deployments/graph-data", ACAOHeaderOverwriteMiddleware)

	url2, _ := url.Parse(ProxyServerResource)
	targetsResource := []*middleware.ProxyTarget{
		{
			URL: url2,
		},
	}
	TransURLResource := middleware.ProxyWithConfig(middleware.ProxyConfig{
		Balancer: middleware.NewRoundRobinBalancer(targetsResource),
		Rewrite: map[string]string{
			"^/api/v1/deployments/graph-data/*": "/resource-data/$1",
		},
	})

	gc2.Use(TransURLResource)

}

func ACAOHeaderOverwriteMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Before(func() {
			setResponseACAOHeaderFromRequest(*ctx.Request(), *ctx.Response())
		})
		return next(ctx)
	}
}

func setResponseACAOHeaderFromRequest(req http.Request, resp echo.Response) {
	resp.Header().Set(echo.HeaderAccessControlAllowOrigin,
		"*")
}
