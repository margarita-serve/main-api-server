package repository

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/domain/entity"
	domRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/domain/repository"
	domSchema "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/domain/schema"
	sysError "git.k3.acornsoft.io/msit-auto-ml/koreserv/system/error"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/utils"
)

// NewEmailRepo new EmailRepo
func NewEmailRepo(h *handler.Handler) (domRepo.IEmailRepo, error) {

	repo := new(EmailRepo)
	repo.handler = h

	cfg, err := h.GetConfig()
	if err != nil {
		return nil, err
	}
	repo.SetDBConnectionName(cfg.Databases.EmailDB.ConnectionName)

	repo.smtp, err = service.NewSMTPSender(h)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

// EmailRepo type Implement IEmailRepo
type EmailRepo struct {
	BaseRepo
	smtp *service.SMTPSender
}

// Send send Email
func (r *EmailRepo) Send(req *domSchema.SendEmailRequest, i identity.Identity) (*domSchema.SendEmailResponse, error) {
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, err
	}

	subjEmail := req.Template.DefaultTemplateVersion.SubjectTpl
	bodyEmail, err := r.compileEmailBody(req.Template.DefaultTemplateVersion.BodyTpl, req.TemplateData)
	if err != nil {
		return nil, err
	}

	// send email depend on processing type : SYNC/ASYNC
	processingStatus := req.ProcessingType
	if domSchema.ProcessingType(req.ProcessingType) == domSchema.SYNCProcess {
		if err := r.sendEmails(req.From, req.To, req.CC, req.BCC, subjEmail, string(bodyEmail), service.EmailFormat(req.Template.EmailFormat)); err != nil {
			return nil, err
		}
	} else {
		go r.sendEmails(req.From, req.To, req.CC, req.BCC, subjEmail, string(bodyEmail), service.EmailFormat(req.Template.EmailFormat))
		// time.Sleep(100 * time.Millisecond)
	}

	// save email to db
	emailEtt := domEntity.Email{
		TemplateID: req.Template.ID,
		UUID:       utils.GenerateUUID(),
		From:       req.From.Email,
		FromName:   req.From.Name,
		To:         req.To.Email,
		ToName:     req.To.Name,
		CC:         r.compileEmail(req.CC),
		BCC:        r.compileEmail(req.BCC),
		Subject:    subjEmail,
		Body:       bodyEmail,
		Status:     fmt.Sprintf("SENT.%s", processingStatus),
	}
	emailEtt.SentBy = i.Claims.Username
	emailEtt.CreatedBy = fmt.Sprintf("%s@%s", i.Claims.Username, i.ClientDevices.IPAddress)

	if err := dbCon.Create(&emailEtt).Error; err != nil {
		if strings.Index(err.Error(), "Error 1062: Duplicate entry") > -1 {
			return nil, &sysError.SystemError{StatusCode: http.StatusConflict, Err: err}
		}
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	// response
	resp := new(domSchema.SendEmailResponse)
	resp.TemplateCode = req.TemplateCode
	resp.Status = emailEtt.Status

	return resp, nil
}

func (r *EmailRepo) compileEmail(mailAddresses []*domSchema.MailAddress) string {
	var str string
	for _, v := range mailAddresses {
		if str != "" {
			str = fmt.Sprintf("%s,", str)
		}
		str = fmt.Sprintf("%s%s <%s>", str, v.Name, v.Email)
	}
	return str
}

func (r *EmailRepo) compileEmailBody(tpl string, data map[string]interface{}) ([]byte, error) {
	t, err := template.New("foo").Parse(tpl)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = t.ExecuteTemplate(buf, "T", data)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (r *EmailRepo) sendEmails(fromEmail *domSchema.MailAddress, toEmail *domSchema.MailAddress, cc []*domSchema.MailAddress, bcc []*domSchema.MailAddress, subject string, body string, format service.EmailFormat) error {
	var toEmails []string
	toEmails = append(toEmails, toEmail.Email)
	for _, e := range cc {
		toEmails = append(toEmails, e.Email)
	}
	for _, e := range bcc {
		toEmails = append(toEmails, e.Email)
	}

	hFrom := r.compileEmail([]*domSchema.MailAddress{fromEmail})
	hTo := r.compileEmail([]*domSchema.MailAddress{toEmail})
	hCc := r.compileEmail(cc)
	hBcc := r.compileEmail(bcc)

	err := r.smtp.SendEmails(toEmails, hFrom, hTo, hCc, hBcc, subject, body, format)
	if err != nil {
		return err
	}

	return nil
}
