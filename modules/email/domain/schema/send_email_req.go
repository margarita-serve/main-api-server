package schema

import domSchemaET "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/domain/schema/email_template"

// SendEmailRequest type
type SendEmailRequest struct {
	TemplateCode   string                 `json:"templateCode"`
	From           *MailAddress           `json:"from"`
	To             *MailAddress           `json:"to"`
	CC             []*MailAddress         `json:"cc"`
	BCC            []*MailAddress         `json:"bcc"`
	TemplateData   map[string]interface{} `json:"templateData"`
	ProcessingType string                 `json:"processingType"`

	Template *domSchemaET.ETFindByCodeData `json:"-"`
}

// MailAddress type
type MailAddress struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// ProcessingType type
type ProcessingType string

const (
	SYNCProcess  ProcessingType = "SYNC"
	ASYNCProcess ProcessingType = "ASYNC"
)
