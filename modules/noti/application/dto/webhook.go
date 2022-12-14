package dto

type WebHook struct {
	ID            string `json:"webHookID" example:"" extensions:"x-order=01"`
	DeploymentID  string `json:"deploymentID" extensions:"x-order=02"`
	TriggerSource string `json:"triggerSource" enums:"DataDrift, Accuracy" example:"DataDrift" validate:"required" extensions:"x-order=03"`
	Name          string `json:"name" validate:"required" example:"pipe-line-trigger" extensions:"x-order=05"`
	URL           string `json:"url" validate:"required" example:"http://example.com/to/webhook/client" extensions:"x-order=06"`
	Method        string `json:"method" enums:"POST,GET" example:"POST" validate:"required" extensions:"x-order=07"`
	CustomHeader  string `json:"customHeader" example:"Content-Type: application/json " extensions:"x-order=08"`
	MessageBody   string `json:"messageBody" example:"{ \"key\": \"value\"}" extensions:"x-order=09"`
	TriggerStatus string `json:"triggerStatus" enums:"AtRisk, Failing" example:"AtRisk" extensions:"x-order=04"`
}

type CreateWebHookRequestDTO struct {
	DeploymentID  string `json:"deploymentID" validate:"required" swaggerignore:"true"`
	TriggerSource string `json:"triggerSource" enums:"DataDrift, Accuracy" example:"DataDrift" validate:"required" extensions:"x-order=01"`
	Name          string `json:"name" validate:"required" example:"pipe-line-trigger" extensions:"x-order=03"`
	URL           string `json:"url" validate:"required" example:"http://example.com/to/webhook/client" extensions:"x-order=04"`
	Method        string `json:"method" enums:"POST,GET" example:"POST" validate:"required" extensions:"x-order=05"`
	CustomHeader  string `json:"customHeader" example:"Content-Type: application/json " extensions:"x-order=06"`
	MessageBody   string `json:"messageBody" example:"{ \"key\": \"value\"}" extensions:"x-order=07"`
	TriggerStatus string `json:"triggerStatus" enums:"AtRisk, Failing" example:"AtRisk" extensions:"x-order=02"`
}

type CreateWebHookResponseDTO struct {
	WebHookID string `json:"webHookID"`
}

type DeleteWebHookRequestDTO struct {
	DeploymentID string `json:"deploymentID" validate:"required" swaggerignore:"true"`
	WebHookID    string `json:"webHookID" validate:"required" swaggerignore:"true"`
}

type UpdateWebHookRequestDTO struct {
	DeploymentID  string  `json:"deploymentID" validate:"required" swaggerignore:"true"`
	TriggerSource *string `json:"triggerSource" enums:"DataDrift, Accuracy" example:"DataDrift" validate:"required" extensions:"x-order=1"`
	WebHookID     string  `json:"webHookID" validate:"required" swaggerignore:"true"`
	Name          *string `json:"name" validate:"required" example:"pipe-line-trigger" extensions:"x-order=2"`
	URL           *string `json:"url" validate:"required" example:"http://example.com/to/webhook/client" extensions:"x-order=3"`
	Method        *string `json:"method" enums:"POST,GET" example:"POST" validate:"required" extensions:"x-order=4"`
	CustomHeader  *string `json:"customHeader" example:"Content-Type: application/json " extensions:"x-order=5"`
	MessageBody   *string `json:"messageBody" example:"{ \"key\": \"value\"}" extensions:"x-order=6"`
}

type GetWebHookRequestDTO struct {
	DeploymentID string `json:"deploymentID" validate:"required" swaggerignore:"true"`
	WebHookID    string `json:"webHookID" validate:"required" swaggerignore:"true"`
}

type GetWebHookResponseDTO struct {
	WebHook
}

type GetWebHookListRequestDTO struct {
	DeploymentID string `json:"deploymentID" validate:"required" swaggerignore:"true"`
	Name         string `json:"name" extensions:"x-order=1"`                  // ????????????: ?????? ???
	Limit        int    `json:"limit" extensions:"x-order=2"`                 // ????????? ?????? ??? ??????
	Page         int    `json:"page" extensions:"x-order=3"`                  // ?????? ??? ?????????, ??? ????????? TotalPages ?????? ????????? ?????? ??? ?????????
	Sort         string `enums:"CreateAsc,CreateDesc" extensions:"x-order=4"` //????????????, CreateAsc: ???????????? ????????????, CraeteDesc: ???????????? ?????????
}

type GetWebHookListResponseDTO struct {
	Limit      int
	Page       int
	Sort       string
	TotalRows  int64
	TotalPages int
	Rows       interface{}
}

type InternalGetWebHookRequestDTO struct {
	DeploymentID  string
	TriggerSource string
	TriggerStatus string
}

type InternalGetWebHookResponseDTO struct {
	WebHookList []*WebHook
}

type TestWebHookRequestDTO struct {
	DeploymentID  string `json:"deploymentID" validate:"required" swaggerignore:"true"`                                                     // ?????? ID
	TriggerSource string `json:"triggerSource" enums:"DataDrift, Accuracy" example:"DataDrift" validate:"required" extensions:"x-order=01"` // ???????????? ??? ???????????? ??????
	Name          string `json:"name" validate:"required" example:"pipe-line-trigger" extensions:"x-order=03"`                              // ????????? ?????? ???
	URL           string `json:"url" validate:"required" example:"http://example.com/to/webhook/client" extensions:"x-order=04"`            // callback URL
	Method        string `json:"method" enums:"POST,GET" example:"POST" validate:"required" extensions:"x-order=05"`                        // http method
	CustomHeader  string `json:"customHeader" example:"Content-Type: application/json " extensions:"x-order=06"`                            // http header
	MessageBody   string `json:"messageBody" example:"{ \"key\": \"value\"}" extensions:"x-order=07"`                                       // http message body
	TriggerStatus string `json:"triggerStatus" enums:"AtRisk, Failing" example:"AtRisk" extensions:"x-order=02"`                            // ???????????? ??? ???????????? ?????? (????????? ?????? ???????????? ???????????? ???)
}

type SendWebHookRequestDTO struct {
	DeploymentID  string
	TriggerSource string
	TriggerStatus string
}
