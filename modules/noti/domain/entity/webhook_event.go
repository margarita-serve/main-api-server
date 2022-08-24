package entity

type WebHookEvent struct {
	DeploymentID  string
	WebHookID     string
	ID            string
	URL           string
	Method        string
	CustomHeader  string
	MessageBody   string
	TriggerSource string
	SendStatus    string //Created, Success, Fail
	BaseEntity
}

//Validate
// func Validate(r *WebHookEvent) error {
// 	return validation.ValidateStruct(r,
// 		validation.Field(&r.SendStatus, validation.Required, validation.In("Created", "Success", "Fail")),
// 	)
// }

func NewWebHookEvent(id string, deploymentID string, URL string, method string, customHeader string, messageBody string, triggerSource string) (*WebHookEvent, error) {

	var baseEntity BaseEntity
	baseEntity.CreatedBy = "testuser"

	WebHookEvent := &WebHookEvent{
		ID:            id,
		DeploymentID:  deploymentID,
		URL:           URL,
		Method:        method,
		CustomHeader:  customHeader,
		MessageBody:   messageBody,
		TriggerSource: triggerSource,
		SendStatus:    "Created",
		BaseEntity:    baseEntity,
	}

	//Validate
	// err := Validate(WebHookEvent)
	// if err != nil {
	// 	return nil, err
	// }

	return WebHookEvent, nil
}

func (d *WebHookEvent) SetSendStatus(req string) {
	if req == "" {
		req = "Created"
	}
	d.SendStatus = req
}
