package common

import "git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"

type SendEmailReqDTO struct {
	TemplateCode   string                 `json:"templateCode"`
	From           *MailAddressDTO        `json:"from"`
	To             *MailAddressDTO        `json:"to"`
	CC             []*MailAddressDTO      `json:"cc"`
	BCC            []*MailAddressDTO      `json:"bcc"`
	TemplateData   map[string]interface{} `json:"templateData"`
	ProcessingType string                 `json:"processingType"`
	// domSchema.SendEmailRequest
}

type SendEmailResponse struct {
	TemplateCode string `json:"templateCode"`
	Status       string `json:"status"`
}

// SendEmailResDTO type
type SendEmailResDTO struct {
	SendEmailResponse
}

// MailAddressDTO type
type MailAddressDTO struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type IEmailService interface {
	SendInternal(req *SendEmailReqDTO, i identity.Identity) (*SendEmailResDTO, error)
}
