package router

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/feature"
	"github.com/labstack/echo/v4"
)

// SetSystem set FSystem Router
func SetSystem(eg *echo.Group, f *feature.FSystem) {
	eg.GET("/health", f.HealthCheck)
	// eg.GET("/captcha/generate", f.GenerateCaptcha)
	// eg.GET("/captcha/image/:captchaID", f.GenerateCaptchaImage)
}
