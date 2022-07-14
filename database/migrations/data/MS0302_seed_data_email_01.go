package data

import (
	domEmailEtt "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/domain/entity"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/utils"
)

// EmailTemplate01 data (TEXT)
func EmailTemplate01() domEmailEtt.EmailTemplate {
	return domEmailEtt.EmailTemplate{
		UUID:        utils.GenerateUUID(),
		Code:        "activate-registration-plaintext",
		Name:        "Activate Registration Email (PlainText)",
		IsActive:    true,
		EmailFormat: "TEXT",
	}
}

// EmailTemplate01Version data
func EmailTemplate01Version() domEmailEtt.EmailTemplateVersion {
	return domEmailEtt.EmailTemplateVersion{
		Version:    utils.GenSemVersion(""),
		SubjectTpl: "Activate Registration",
		BodyTpl: `{{define "T"}}Dear {{index . "Header.Name"}},

Please click on the url bellow to complete the verification process for account "{{index . "Body.UserAccount"}}":

{{index . "Body.ActivationURL"}}

If you didn't attempt to verify your email address with our service, delete this email.

Cheers,

{{index . "Footer.Name"}}
{{end}}`,
	}
}
