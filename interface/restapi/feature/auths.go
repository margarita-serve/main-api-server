package feature

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	//captcha "git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/feature/captcha"

	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/response"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/auths/application"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/auths/application/dto"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"github.com/labstack/echo/v4"
)

// NewFAuths new  FAuths
func NewFAuths(h *handler.Handler) (*FAuths, error) {
	var err error

	f := new(FAuths)
	f.handler = h

	if f.appAuths, err = application.NewAuthsApp(h); err != nil {
		return nil, err
	}

	return f, nil
}

// FAuths feature Auths
type FAuths struct {
	BaseFeature
	appAuths *application.AuthsApp
}

// RegisterUser register user
// @Summary register user
// @Description  유저 생성
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body dto.RegisterReqDTO true "register user"
// @Success 200 {object} dto.RegisterResDTO
// @Router      /auths/register [post]
func (f *FAuths) RegisterUser(c echo.Context) error {

	req := new(dto.RegisterReqDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	// if !f.inTestMode() {
	// 	decodedCaptcha, err := captcha.DecodeCaptcha(req.CaptchaID, c.RealIP())
	// 	if err != nil {
	// 		return response.FailWithMessage(err.Error(), c)
	// 	}

	// 	if !captcha.VerifyString(decodedCaptcha, req.Captcha) {
	// 		return response.FailWithMessage("Captcha verification code error", c)
	// 	}
	// }

	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp, err := f.appAuths.AuthenticationSvc.Register(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.CreatedWithData(resp, c)
}

// ActivateRegistration activater user registration
func (f *FAuths) ActivateRegistration(c echo.Context) error {

	//params
	format := strings.ToLower(c.Param("format"))

	req := new(dto.ActivateRegistrationReqDTO)
	req.ActivationCode = c.Param("activationCode")

	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp, err := f.appAuths.AuthenticationSvc.ActivateRegistration(req, i)
	if err != nil {
		if format == "html" {
			data := map[string]interface{}{
				"message": err.Error(),
			}
			return c.Render(http.StatusBadRequest, "auths/activate.registration", data)
		}

		return f.translateErrorMessage(err, c)
	}

	if format == "html" {
		data := map[string]interface{}{
			"message": fmt.Sprintf("Your user [%s] are now active", resp.Email),
		}
		return c.Render(http.StatusOK, "auths/activate.registration", data)
	}

	return response.OkWithData(resp, c)
}

// Login user Login
// @Summary Login user
// @Description  유저 로그인
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body dto.LoginReqDTO true "login user"
// @Success 200 {object} dto.LoginResDTO
// @Router      /auths/login [post]
func (f *FAuths) Login(c echo.Context) error {

	req := new(dto.LoginReqDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	// if !f.inTestMode() {
	// 	decodedCaptcha, err := captcha.DecodeCaptcha(req.CaptchaID, c.RealIP())
	// 	if err != nil {
	// 		return response.FailWithMessage(err.Error(), c)
	// 	}

	// 	if !captcha.VerifyString(decodedCaptcha, req.Captcha) {
	// 		return response.FailWithMessage("Captcha verification code error", c)
	// 	}
	// }

	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp, err := f.appAuths.AuthenticationSvc.Login(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	// set interface-session-jwt on cacher
	if err := f._setSession(resp.Token, resp.ExpiredAt); err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// LoginApp login client app
func (f *FAuths) LoginApp(c echo.Context) error {

	req := new(dto.LoginAppReqDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp, err := f.appAuths.AuthenticationSvc.LoginApp(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	// set interface-session-jwt on cacher
	if err := f._setSession(resp.Token, resp.ExpiredAt); err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

func (f *FAuths) _setSession(token string, expiredAt int64) error {
	sessionValue := token
	expiredAtDT := time.Unix(expiredAt/1000, 0) // ExpiredAt = UnixTimeStamp
	expiration := expiredAtDT.Sub(time.Now()).Seconds()
	if err := f.SetSession(sessionValue, int64(expiration)); err != nil {
		return err
	}
	return nil
}
