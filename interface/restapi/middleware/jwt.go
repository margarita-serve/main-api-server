package middleware

import (
	"net/http"
	"strings"

	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/response"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/utils"
	"github.com/labstack/echo/v4"
)

// JWTVerifier verify JWT token from internal Identity Provider
func JWTVerifier(h *handler.Handler) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authHeader := c.Request().Header.Get("Authorization")
			//is bearer
			if !strings.HasPrefix(strings.ToLower(authHeader), "bearer") {
				return response.FailWithMessageWithCode(http.StatusForbidden, "Authorization: Bearer not found", c)
			}

			// get token
			authHeaderPart := strings.Split(authHeader, " ")
			token := authHeaderPart[len(authHeaderPart)-1]
			if token == "" {
				return response.FailWithMessageWithCode(http.StatusForbidden, "Invalid token or illegas access", c)
			}

			j, err := identity.NewJWT(h)
			if err != nil {
				return response.FailWithMessageWithCode(http.StatusInternalServerError, err.Error(), c)
			}

			// parseToken parses the information contained in the token
			claims, err := j.ParseToken(token)
			if err != nil {
				data := map[string]interface{}{
					"reload": true,
				}
				if err == identity.ErrTokenExpired {
					return response.FailWithDetailed(response.ERROR, data, "Authorization has expired", c)
				}
				return response.FailWithDetailed(response.ERROR, data, err.Error(), c)
			}

			// 과제 API연동시 업체측 어려움으로 고정 토큰생성 후 파싱만 진행으로 변경
			// verify token to persistent storage
			// exist, err := isSessionExist(token, h)
			// if err != nil {
			// 	return response.FailWithMessageWithCode(http.StatusInternalServerError, err.Error(), c)
			// }
			// if !exist {
			// 	return response.FailWithMessageWithCode(http.StatusInternalServerError, fmt.Sprintf("Identity Provider Error [%s]", "Invalid Token"), c)
			// }

			// // if expired
			// now := time.Now().Unix()
			// if claims.ExpiresAt < now {
			// 	return response.FailWithMessageWithCode(http.StatusInternalServerError, fmt.Sprintf("Identity Provider Error [%s]", "Token Expired"), c)
			// }

			// // if not valid before
			// if claims.NotBefore > now {
			// 	return response.FailWithMessageWithCode(http.StatusInternalServerError, fmt.Sprintf("Identity Provider Error [%s]", "Token Not Valid Berofe"), c)
			// }
			c.Set("identity.token.jwt", token)
			c.Set("identity.token.jwt.claims", claims)

			return next(c)
		}
	}
}

func isSessionExist(sessionValue string, h *handler.Handler) (bool, error) {
	cfg, err := h.GetConfig()
	if err != nil {
		return false, err
	}
	ce, err := h.GetCacher(cfg.Caches.SessionCache.ConnectionName)
	ce.Context = "interface"
	ce.Container = "session"
	ce.Component = "jwt"

	sessionKey := utils.MD5([]byte(sessionValue))

	return ce.IsExist(sessionKey), nil
}
