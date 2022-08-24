package dto

type WebHook struct {
	ID            string `json:"id" example:""`
	DeploymentID  string `json:"deploymentID"`
	TriggerSource string `json:"triggerSource" enums:"Datadrift, Accuracy" example:"Datadrift" validate:"required" extensions:"x-order=1"`
	Name          string `json:"name" validate:"required" example:"pipe-line-trigger" extensions:"x-order=1"`
	URL           string `json:"url" validate:"required" example:"http://example.com/to/webhook/client" extensions:"x-order=2"`
	Method        string `json:"method" enums:"POST,GET" example:"POST" validate:"required" extensions:"x-order=3"`
	CustomHeader  string `json:"customHeader" example:"Content-Type: application/json " extensions:"x-order=4"`
	MessageBody   string `json:"messageBody" example:"{ \"key\": \"value\"}" extensions:"x-order=5"`
}

type CreateWebHookRequestDTO struct {
	DeploymentID  string `json:"deploymentID" validate:"required" swaggerignore:"true"`
	TriggerSource string `json:"triggerSource" enums:"Datadrift, Accuracy" example:"Datadrift" validate:"required" extensions:"x-order=1"`
	Name          string `json:"name" validate:"required" example:"pipe-line-trigger" extensions:"x-order=1"`
	URL           string `json:"url" validate:"required" example:"http://example.com/to/webhook/client" extensions:"x-order=2"`
	Method        string `json:"method" enums:"POST,GET" example:"POST" validate:"required" extensions:"x-order=3"`
	CustomHeader  string `json:"customHeader" example:"Content-Type: application/json " extensions:"x-order=4"`
	MessageBody   string `json:"messageBody" example:"{ \"key\": \"value\"}" extensions:"x-order=5"`
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
	TriggerSource *string `json:"triggerSource" enums:"Datadrift, Accuracy" example:"Datadrift" validate:"required" extensions:"x-order=1"`
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
	Name         string `json:"name" extensions:"x-order=1"`                  // 검색조건: 배포 명
	Limit        int    `json:"limit" extensions:"x-order=2"`                 // 한번에 조회 할 건수
	Page         int    `json:"page" extensions:"x-order=3"`                  // 조회 할 페이지, 첫 조회후 TotalPages 범위 내에서 선택 후 보낸다
	Sort         string `enums:"CreateAsc,CreateDesc" extensions:"x-order=4"` //정열방식, CreateAsc: 생성시간 내림차순, CraeteDesc: 생성시간 역차순
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
}

type InternalGetWebHookResponseDTO struct {
	WebHookList []*WebHook
}

type TestWebHookRequestDTO struct {
	DeploymentID  string `json:"deploymentID" validate:"required" swaggerignore:"true"`
	TriggerSource string `json:"triggerSource" enums:"Datadrift, Accuracy" example:"Datadrift" validate:"required" extensions:"x-order=1"`
	Name          string `json:"name" validate:"required" example:"pipe-line-trigger" extensions:"x-order=1"`
	URL           string `json:"url" validate:"required" example:"http://example.com/to/webhook/client" extensions:"x-order=2"`
	Method        string `json:"method" enums:"POST,GET" example:"POST" validate:"required" extensions:"x-order=3"`
	CustomHeader  string `json:"customHeader" example:"Content-Type: application/json " extensions:"x-order=4"`
	MessageBody   string `json:"messageBody" example:"{ \"key\": \"value\"}" extensions:"x-order=5"`
}

type SendWebHookRequestDTO struct {
	DeploymentID  string
	TriggerSource string
}
