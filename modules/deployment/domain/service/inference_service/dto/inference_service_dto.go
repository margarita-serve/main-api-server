package dto

type InferenceServiceCreateRequest struct {
	ConnectionInfo        string
	Namespace             string
	Inferencename         string
	ModelFrameWork        string
	ModelFrameWorkVersion string
	ModelURL              string
	ModelName             string
}

type InferenceServiceCreateResponse struct {
	Message       string
	Inferencename string
}

type InferenceServiceDeleteRequest struct {
	ConnectionInfo string
	Namespace      string
	Inferencename  string
}

type InferenceServiceDeleteResponse struct {
	Message       string
	Inferencename string
}

type InferenceServiceGetRequest struct {
	ConnectionInfo string
	Namespace      string
	Inferencename  string
}

type InferenceServiceGetResponse struct {
	Message       string
	Inferencename string
}

type InferenceServiceModelReplaceRequest struct {
	ConnectionInfo        string
	Namespace             string
	Inferencename         string
	ModelName             string
	ModelFrameWork        string
	ModelFrameWorkVersion string
	ModelURL              string
}

type InferenceServiceModelReplaceResponse struct {
	Namespace      string
	Inferencename  string
	ModelHistoryID string
}
