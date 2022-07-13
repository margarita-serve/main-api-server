package repository

import (
	schema "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/domain/schema/email_template"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
)

// IEmailTemplateRepo interface
type IEmailTemplateRepo interface {
	ListAll(i identity.Identity) (*schema.ETListAllResponse, error)
	FindByCode(req *schema.ETFindByCodeRequest, i identity.Identity) (*schema.ETFindByCodeResponse, error)
	Create(req *schema.ETCreateRequest, i identity.Identity) (*schema.ETCreateResponse, error)
	Update(req *schema.ETUpdateRequest, i identity.Identity) (*schema.ETUpdateResponse, error)
	SetActive(req *schema.ETSetActiveRequest, i identity.Identity) (*schema.ETSetActiveResponse, error)
	Delete(req *schema.ETDeleteRequest, i identity.Identity) (*schema.ETDeleteResponse, error)
}
