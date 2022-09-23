package service

import (
	common "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/common"
	appDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/application/dto"
	domRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/domain/repository"
	domSchema "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/domain/schema"
	domSchemaET "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/domain/schema/email_template"
	infRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/infrastructure/repository"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
)

// NewEmailService new EmailService
func NewEmailService(h *handler.Handler) (*EmailService, error) {
	var err error

	svc := new(EmailService)
	svc.handler = h
	if err := svc.initBaseService(); err != nil {
		return nil, err
	}

	if svc.repoEmailTpl, err = infRepo.NewEmailTemplateRepo(h); err != nil {
		return nil, err
	}
	if svc.repoEmail, err = infRepo.NewEmailRepo(h); err != nil {
		return nil, err
	}

	return svc, nil
}

// EmailService type
type EmailService struct {
	BaseService
	repoEmail    domRepo.IEmailRepo
	repoEmailTpl domRepo.IEmailTemplateRepo
}

//Send send Email
func (s *EmailService) Send(req *appDTO.SendEmailReqDTO, i identity.Identity) (*appDTO.SendEmailResDTO, error) {
	// authorization
	// if (i.CanAccessCurrentRequest() == false) && (i.CanAccess("", "system.module.email.send", "EXECUTE", nil) == false) {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	// request domain
	reqDom := domSchema.SendEmailRequest{
		TemplateCode:   req.TemplateCode,
		From:           &domSchema.MailAddress{Email: req.From.Email, Name: req.From.Name},
		To:             &domSchema.MailAddress{Email: req.To.Email, Name: req.To.Name},
		CC:             req.ConvertCC2Domain(),
		BCC:            req.ConvertBCC2Domain(),
		TemplateData:   req.TemplateData,
		ProcessingType: req.ProcessingType,
	}

	if err := reqDom.Validate(); err != nil {
		return nil, err
	}

	// retrieve and assign email template
	// -->
	reqET := domSchemaET.ETFindByCodeRequest{
		Code: req.TemplateCode,
	}
	tpl, err := s.repoEmailTpl.FindByCode(&reqET, s.systemIdentity)
	if err != nil {
		return nil, err
	}
	reqDom.Template = &tpl.Data
	// <--

	res, err := s.repoEmail.Send(&reqDom, i)
	if err != nil {
		return nil, err
	}

	// response - dto
	resDTO := new(appDTO.SendEmailResDTO)
	resDTO.TemplateCode = res.TemplateCode
	resDTO.Status = res.Status

	return resDTO, nil
}

//Send(req SendEmailReqDTO, i identity.Identity) (SendEmailResDTO, error)
func (s *EmailService) SendInternal(req *common.SendEmailReqDTO, i identity.Identity) (*common.SendEmailResDTO, error) {
	// authorization
	// if (i.CanAccessCurrentRequest() == false) && (i.CanAccess("", "system.module.email.send", "EXECUTE", nil) == false) {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	// request domain
	reqSend := appDTO.SendEmailReqDTO{
		TemplateCode: req.TemplateCode,
		From:         (*appDTO.MailAddressDTO)(req.From),
		To:           (*appDTO.MailAddressDTO)(req.To),
		// CC:             req.CC,
		// BCC:            req.BCC,
		TemplateData:   req.TemplateData,
		ProcessingType: req.ProcessingType,
	}

	res, err := s.Send(&reqSend, i)
	if err != nil {
		return nil, err
	}

	// response - dto
	resDTO := new(common.SendEmailResDTO)
	resDTO.TemplateCode = res.TemplateCode
	resDTO.Status = res.Status

	return resDTO, nil
}
