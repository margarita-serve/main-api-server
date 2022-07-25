package router

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"net/url"
)

func SetGraph(eg *echo.Group, h *handler.Handler) {
	cfg, _ := h.GetConfig()
	ProxyServer := cfg.Connectors.BokehServer.Endpoint

	eg.GET("/featuredetail/autoload.js",
		Noop,
		ACAOHeaderOverwriteMiddleware,
		middleware.ProxyWithConfig(middleware.ProxyConfig{
			Balancer: singleTargetBalancer(ProxyServer),
		}),
	)

	eg.GET("/featuredetail/ws",
		Noop,
		ACAOHeaderOverwriteMiddleware,
		middleware.ProxyWithConfig(middleware.ProxyConfig{
			Balancer: singleTargetBalancer(ProxyServer),
		}),
	)

	eg.GET("/featuredrift/autoload.js",
		Noop,
		ACAOHeaderOverwriteMiddleware,
		middleware.ProxyWithConfig(middleware.ProxyConfig{
			Balancer: singleTargetBalancer(ProxyServer),
		}),
	)

	eg.GET("/featuredrift/ws",
		Noop,
		ACAOHeaderOverwriteMiddleware,
		middleware.ProxyWithConfig(middleware.ProxyConfig{
			Balancer: singleTargetBalancer(ProxyServer),
		}),
	)
}

func SetGraphJS(eg *echo.Group, h *handler.Handler) {
	cfg, _ := h.GetConfig()
	ProxyServer := cfg.Connectors.BokehServer.Endpoint

	eg.GET("/bokeh.min.js",
		Noop,
		ACAOHeaderOverwriteMiddleware,
		middleware.ProxyWithConfig(middleware.ProxyConfig{
			Balancer: singleTargetBalancer(ProxyServer),
		}),
	)

	eg.GET("/bokeh-gl.min.js",
		Noop,
		ACAOHeaderOverwriteMiddleware,
		middleware.ProxyWithConfig(middleware.ProxyConfig{
			Balancer: singleTargetBalancer(ProxyServer),
		}),
	)

	eg.GET("/bokeh-widgets.min.js",
		Noop,
		ACAOHeaderOverwriteMiddleware,
		middleware.ProxyWithConfig(middleware.ProxyConfig{
			Balancer: singleTargetBalancer(ProxyServer),
		}),
	)

	eg.GET("/bokeh-tables.min.js",
		Noop,
		ACAOHeaderOverwriteMiddleware,
		middleware.ProxyWithConfig(middleware.ProxyConfig{
			Balancer: singleTargetBalancer(ProxyServer),
		}),
	)

	eg.GET("/bokeh-mathjax.min.js",
		Noop,
		ACAOHeaderOverwriteMiddleware,
		middleware.ProxyWithConfig(middleware.ProxyConfig{
			Balancer: singleTargetBalancer(ProxyServer),
		}),
	)
}

func Noop(ctx echo.Context) (err error) {
	ctx.String(
		http.StatusNotImplemented,
		"No op handler should never be reached!",
	)

	return err
}

func setResponseACAOHeaderFromRequest(req http.Request, resp echo.Response) {
	resp.Header().Set(echo.HeaderAccessControlAllowOrigin,
		"*")
}

func ACAOHeaderOverwriteMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Before(func() {
			setResponseACAOHeaderFromRequest(*ctx.Request(), *ctx.Response())
		})
		return next(ctx)
	}
}

func singleTargetBalancer(server string) middleware.ProxyBalancer {
	url, _ := url.Parse(server)
	targetURL := []*middleware.ProxyTarget{
		{
			URL: url,
		},
	}
	return middleware.NewRoundRobinBalancer(targetURL)
}
