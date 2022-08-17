package service

import (
	"testing"

	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/config"
	//"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
)

func newConfig(t *testing.T) (*config.Config, error) {
	c, _, err := config.NewConfig("../../../../conf")
	if err != nil {
		return nil, err
	}
	//c.IAM.Casbin.ModelPath = "../../../../../conf/casbin/casbin_rbac_rest_model.conf"
	return c, nil
}

// func newIdentity(h *handler.Handler, t *testing.T) identity.Identity {
// 	j, err := identity.NewJWT(h)
// 	if err != nil {
// 		return identity.Identity{}
// 	}
// 	claims, token, _, err := j.GenerateAnonymousToken()
// 	if err != nil {
// 		return identity.Identity{}
// 	}
// 	i, err := identity.NewIdentity(identity.DefaultIdentity, identity.TokenJWT, token, claims, context.NewCtx(context.SystemCtx), h)
// 	if err != nil {
// 		return identity.Identity{}
// 	}
// 	return i
// }
