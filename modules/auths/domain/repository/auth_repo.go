package repository

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/auths/domain/entity"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/auths/domain/schema"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
)

// IAuthenticationRepo interface
type IAuthenticationRepo interface {
	Register(req *schema.RegisterRequest, i identity.Identity) (*schema.RegisterResponse, error)
	ActivateRegistration(req *schema.ActivateRegistrationRequest, i identity.Identity) (*schema.ActivateRegistrationResponse, error)
	Login(req *schema.LoginRequest, i identity.Identity) (*schema.LoginResponse, error)
	LoginApp(req *schema.LoginAppRequest, i identity.Identity) (*schema.LoginAppResponse, error)
	GetUserByName(req string) (*entity.SysUser, error)
}
