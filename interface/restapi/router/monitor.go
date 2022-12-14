package router

import (
	"net/http"
	"net/url"

	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/feature"
	internalMiddleware "git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/middleware"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetMonitor(eg *echo.Group, h *handler.Handler, f *feature.FMonitor) {
	cfg, _ := h.GetConfig()
	ProxyServer := cfg.Connectors.GraphServer.Endpoint

	gc := eg.Group("/deployments/:deploymentID/monitor")
	gc.Use(internalMiddleware.JWTVerifier(f.GetHandler()))

	gc.GET("", f.GetMonitorSetting)
	gc.GET("/detail", f.GetDetail)
	gc.GET("/drift", f.GetDrift)
	gc.POST("/actual", f.UploadActual)
	gc.GET("/accuracy", f.GetAccuracy)
	gc.PATCH("", f.PatchMonitorSetting)
	//gc.PATCH("/association-id", f.UpdateAssociationID)
	//gc.POST("", f.Create)

	url1, _ := url.Parse(ProxyServer)
	targets := []*middleware.ProxyTarget{
		{
			URL: url1,
		},
	}
	transURL := middleware.ProxyWithConfig(middleware.ProxyConfig{
		Balancer: middleware.NewRoundRobinBalancer(targets),
		Rewrite: map[string]string{
			"^/api/v1/deployments/*/monitor/graph/*": "/api/v1/deployments/graph-svr/$1/$2",
		},
		Transport: MyRoundTripper{r: http.DefaultTransport},
	})

	gc2 := gc.Group("/graph", ACAOHeaderOverwriteMiddleware)
	gc2.Use(transURL)
}

type MyRoundTripper struct {
	r http.RoundTripper
}

func (mrt MyRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("X-Forwarded-Proto", r.Header.Get("X-Forwarded-Proto"))
	//fmt.Println("X-Forwarded-Proto: %s", r.Header.Get("X-Forwarded-Proto"))
	return mrt.r.RoundTrip(r)
}
