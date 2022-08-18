package dto

type ServiceHealthCreateRequest struct {
	InferenceName  string
	ModelHistoryID string
}

type ServiceHealthCreateResponse struct {
	Message       string
	InferenceName string
}

type ServiceHealthDeleteRequest struct {
	InferenceName string
}

type ServiceHealthDeleteResponse struct {
	Message       string
	InferenceName string
}

type ServiceHealthGetRequest struct {
	InferenceName  string
	ModelHistoryID string
	DataType       string
	StartTime      string
	EndTime        string
}

type ServiceHealthGetResponse struct {
	Message   string
	Data      string
	StartTime string
	EndTIme   string
}

type ServiceHealthEnableRequest struct {
	InferenceName string
}

type ServiceHealthEnableResponse struct {
	Message       string
	InferenceName string
}
