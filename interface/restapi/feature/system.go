package feature

import (
	response "git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/response"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"github.com/labstack/echo/v4"
)

// NewSystem new FSystem
func NewSystem(h *handler.Handler) (*FSystem, error) {

	f := new(FSystem)
	f.handler = h

	return f, nil
}

// FSystem represent FSystem
type FSystem struct {
	BaseFeature
}

// HealthCheck display system health check
func (f *FSystem) HealthCheck(c echo.Context) error {
	data := map[string]interface{}{"serviceStatus": "OK"}
	return response.OkWithData(data, c)
}

// // GenerateCaptcha generate Captcha
// func (f *FSystem) GenerateCaptcha(c echo.Context) error {
// 	cfg, err := f.handler.GetConfig()
// 	if err != nil {
// 		return err
// 	}

// 	resp := captcha.GenerateCaptchaID(cfg, c)

// 	return response.OkWithData(resp, c)
// }

// // GenerateCaptchaImage generate CaptchaImage
// func (f *FSystem) GenerateCaptchaImage(c echo.Context) error {
// 	cfg, err := f.handler.GetConfig()
// 	if err != nil {
// 		return err
// 	}

// 	return captcha.CaptchaServeHTTP(cfg, c)
// }
