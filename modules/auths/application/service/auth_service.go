package service

import (
	"fmt"

	appDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/auths/application/dto"
	domRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/auths/domain/repository"
	domSchema "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/auths/domain/schema"
	infRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/auths/infrastructure/repository"
	appEmailDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/application/dto"
	appEmailSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/email/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
)

// NewAuthenticationSvc new AuthenticationSvc
func NewAuthenticationSvc(h *handler.Handler) (*AuthenticationSvc, error) {
	var err error

	svc := new(AuthenticationSvc)
	svc.handler = h
	if err := svc.initBaseService(); err != nil {
		return nil, err
	}

	if svc.repo, err = infRepo.NewAuthenticationRepo(h); err != nil {
		return nil, err
	}

	if svc.emailService, err = appEmailSvc.NewEmailService(h); err != nil {
		return nil, err
	}

	return svc, nil
}

// AuthenticationSvc type
type AuthenticationSvc struct {
	BaseService
	repo         domRepo.IAuthenticationRepo
	emailService *appEmailSvc.EmailService
}

// Register user
func (s *AuthenticationSvc) Register(req *appDTO.RegisterReqDTO, i identity.Identity) (*appDTO.RegisterResDTO, error) {
	reqDom := domSchema.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		NickName: req.NickName,
		// Captcha:   req.Captcha,
		// CaptchaID: req.CaptchaID,
	}

	if err := reqDom.Validate(); err != nil {
		return nil, err
	}

	res, err := s.repo.Register(&reqDom, i)
	if err != nil {
		return nil, err
	}

	// send email via email module (email - generic sub domain)
	// -->
	if err := s._sendActivationCodeEmail(req, res, i); err != nil {
		return nil, err
	}
	// <--

	resDTO := new(appDTO.RegisterResDTO)
	resDTO.Email = res.Email

	return resDTO, nil
}

func (s *AuthenticationSvc) _sendActivationCodeEmail(reqReg *appDTO.RegisterReqDTO, resReg *domSchema.RegisterResponse, i identity.Identity) error {
	cfg, err := s.handler.GetConfig()
	if err != nil {
		return err
	}
	fromEmail := cfg.SMTPServers.DefaultSMTP.SenderEmail
	fromName := cfg.SMTPServers.DefaultSMTP.SenderName

	url := fmt.Sprintf(cfg.IAM.Registration.ActivationURL, i.RequestInfo.Host)
	activationURL := fmt.Sprintf("%s/%s/%s", url, resReg.ActivationCode, "html")

	// send activate registration via email (sub)domain [email module]
	reqEmail := new(appEmailDTO.SendEmailReqDTO)
	reqEmail.TemplateCode = "activate-registration-html"
	reqEmail.From = &appEmailDTO.MailAddressDTO{Email: fromEmail, Name: fromName}
	reqEmail.To = &appEmailDTO.MailAddressDTO{Email: reqReg.Email, Name: reqReg.NickName}
	reqEmail.TemplateData = map[string]interface{}{
		"Header.Name":        reqReg.NickName,
		"Body.UserAccount":   reqReg.Username,
		"Body.ActivationURL": activationURL,
		"Footer.Name":        fromName,
	}
	reqEmail.ProcessingType = "ASYNC"

	if _, err := s.emailService.Send(reqEmail, s.systemIdentity); err != nil {
		return err
	}

	return nil
}

// ActivateRegistration activate user registration
func (s *AuthenticationSvc) ActivateRegistration(req *appDTO.ActivateRegistrationReqDTO, i identity.Identity) (*appDTO.ActivateRegistrationResDTO, error) {
	reqDom := domSchema.ActivateRegistrationRequest{
		ActivationCode: req.ActivationCode,
	}

	if err := reqDom.Validate(); err != nil {
		return nil, err
	}

	res, err := s.repo.ActivateRegistration(&reqDom, i)
	if err != nil {
		return nil, err
	}

	// send email via email module (email - generic sub domain)
	// -->
	if err := s._sendActivationResultEmail(req, res, i); err != nil {
		return nil, err
	}
	// <--

	resDTO := new(appDTO.ActivateRegistrationResDTO)
	resDTO.Email = res.Email
	resDTO.NickName = res.NickName
	resDTO.DefaultRole = res.DefaultRole

	return resDTO, nil
}

func (s *AuthenticationSvc) _sendActivationResultEmail(reqActReg *appDTO.ActivateRegistrationReqDTO, resAtcReg *domSchema.ActivateRegistrationResponse, i identity.Identity) error {
	cfg, err := s.handler.GetConfig()
	if err != nil {
		return err
	}
	fromEmail := cfg.SMTPServers.DefaultSMTP.SenderEmail
	fromName := cfg.SMTPServers.DefaultSMTP.SenderName

	// send activate registration via email (sub)domain [email module]
	reqEmail := new(appEmailDTO.SendEmailReqDTO)
	reqEmail.TemplateCode = "account-activation-html"
	reqEmail.From = &appEmailDTO.MailAddressDTO{Email: fromEmail, Name: fromName}
	reqEmail.To = &appEmailDTO.MailAddressDTO{Email: resAtcReg.Email, Name: resAtcReg.NickName}
	reqEmail.TemplateData = map[string]interface{}{
		"Header.Name": resAtcReg.NickName,
		"Footer.Name": fromName,
	}
	reqEmail.ProcessingType = "ASYNC"

	if _, err := s.emailService.Send(reqEmail, s.systemIdentity); err != nil {
		return err
	}

	return nil
}

// Login user
func (s *AuthenticationSvc) Login(req *appDTO.LoginReqDTO, i identity.Identity) (*appDTO.LoginResDTO, error) {
	reqDom := domSchema.LoginRequest{
		Username: req.Username,
		Password: req.Password,
		// Captcha:   req.Captcha,
		// CaptchaID: req.CaptchaID,
	}

	if err := reqDom.Validate(); err != nil {
		return nil, err
	}

	res, err := s.repo.Login(&reqDom, i)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.LoginResDTO)
	resDTO.TokenType = res.TokenType
	resDTO.Token = res.Token
	resDTO.ExpiredAt = res.ExpiredAt

	return resDTO, nil
}

// LoginApp login app
func (s *AuthenticationSvc) LoginApp(req *appDTO.LoginAppReqDTO, i identity.Identity) (*appDTO.LoginAppResDTO, error) {
	reqDom := domSchema.LoginAppRequest{
		ClientKey: req.ClientKey,
		SecretKey: req.SecretKey,
	}

	if err := reqDom.Validate(); err != nil {
		return nil, err
	}

	res, err := s.repo.LoginApp(&reqDom, i)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.LoginAppResDTO)
	resDTO.TokenType = res.TokenType
	resDTO.ClientAppCode = res.ClientAppCode
	resDTO.ClientAppName = res.ClientAppName
	resDTO.Token = res.Token
	resDTO.ExpiredAt = res.ExpiredAt

	return resDTO, nil
}

// Get userInfo
func (s *AuthenticationSvc) GetUserByName(req *appDTO.GetUserByNameReqDTO, i identity.Identity) (*appDTO.GetUserByNameResDTO, error) {
	res, err := s.repo.GetUserByName(req.UserName, i)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.GetUserByNameResDTO)
	resDTO.Email = res.Email
	return resDTO, nil
}
