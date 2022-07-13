package repository

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/domain/schema"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
)

// IEmailRepo interface
type IEmailRepo interface {
	Send(req *schema.SendEmailRequest, i identity.Identity) (*schema.SendEmailResponse, error)
}
