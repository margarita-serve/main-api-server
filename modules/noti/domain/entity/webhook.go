package entity

import validation "github.com/go-ozzo/ozzo-validation/v4"

type WebHook struct {
	ID            string
	Name          string
	DeploymentID  string
	TriggerSource string //Datadrift, Accuracy
	URL           string //Webhook target URL, including the http:// or https:// prefix along with the url params
	Method        string //GET, POST - Default is POST
	CustomHeader  string //Specify any custom header lines here
	MessageBody   string //Put the body of your message here
	//CustomParametersAndSecrets string //Custom parameters and secrets allow you to add unique parameters and secure elements such as passwords
	TriggerStatus string // Trigger Status AtRisk, Failing
	BaseEntity
}

// Validate
func Validate(r *WebHook) error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required),
		validation.Field(&r.TriggerSource, validation.Required, validation.In("Datadrift", "Accuracy")),
		validation.Field(&r.URL, validation.Required),
		validation.Field(&r.Method, validation.Required, validation.In("POST", "GET")),
		validation.Field(&r.TriggerStatus, validation.Required, validation.In("AtRisk", "Failing")),
	)
}

func NewWebHook(id string, name string, deploymentID string, triggerSource string, URL string, method string, customHeader string, messageBody string, createUser string, triggerStatus string) (*WebHook, error) {

	var baseEntity BaseEntity
	baseEntity.CreatedBy = createUser

	if triggerStatus == "" {
		triggerStatus = "AtRisk"
	}

	WebHook := &WebHook{
		ID:            id,
		Name:          name,
		DeploymentID:  deploymentID,
		TriggerSource: triggerSource,
		TriggerStatus: triggerStatus,
		URL:           URL,
		Method:        method,
		CustomHeader:  customHeader,
		MessageBody:   messageBody,
		BaseEntity:    baseEntity,
	}

	// Validate
	err := Validate(WebHook)
	if err != nil {
		return nil, err
	}

	return WebHook, nil
}

func (d *WebHook) SetURL(req string) {
	d.URL = req
}

func (d *WebHook) SetCustomHeader(req string) {
	d.CustomHeader = req
}

func (d *WebHook) SetMessageBody(req string) {
	d.MessageBody = req
}

func (d *WebHook) SetMethod(req string) {
	if req == "" {
		req = "POST"
	}
	d.Method = req
}

func (d *WebHook) SetTriggerSource(req string) {
	d.TriggerSource = req
}

func (d *WebHook) SetName(req string) {
	if req == "" {
		req = "default"
	}
	d.Name = req
}
