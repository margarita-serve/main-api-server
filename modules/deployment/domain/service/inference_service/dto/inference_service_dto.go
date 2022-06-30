package dto

type InferenceServiceCreateRequest struct {
	ConnectionInfo        string
	Namespace             string
	DeploymentID          string
	ModelFrameWork        string
	ModelFrameWorkVersion string
	ModelURL              string
	ModelName             string
	RequestCPU            float32
	RequestMEM            float32
	LimitCPU              float32
	LimitMEM              float32
}

type InferenceServiceCreateResponse struct {
	Message        string
	DeploymentID   string
	ModelHistoryID string
}

type InferenceServiceDeleteRequest struct {
	ConnectionInfo string
	Namespace      string
	DeploymentID   string
}

type InferenceServiceDeleteResponse struct {
	Message      string
	DeploymentID string
}

type InferenceServiceGetRequest struct {
	ConnectionInfo string
	Namespace      string
	DeploymentID   string
}

type InferenceServiceGetResponse struct {
	Message      string
	DeploymentID string
}

type InferenceServiceReplaceModelRequest struct {
	ConnectionInfo        string
	Namespace             string
	DeploymentID          string
	ModelName             string
	ModelFrameWork        string
	ModelFrameWorkVersion string
	ModelURL              string
	RequestCPU            float32
	RequestMEM            float32
	LimitCPU              float32
	LimitMEM              float32
}

type InferenceServiceReplaceModelResponse struct {
	Namespace      string
	DeploymentID   string
	ModelHistoryID string
}

type InferenceServiceActiveRequest struct {
	ConnectionInfo        string
	Namespace             string
	DeploymentID          string
	ModelName             string
	ModelFrameWork        string
	ModelFrameWorkVersion string
	ModelURL              string
	RequestCPU            float32
	RequestMEM            float32
	LimitCPU              float32
	LimitMEM              float32
}

type InferenceServiceInActiveRequest struct {
	ConnectionInfo        string
	Namespace             string
	DeploymentID          string
	ModelName             string
	ModelFrameWork        string
	ModelFrameWorkVersion string
	ModelURL              string
	RequestCPU            float32
	RequestMEM            float32
	LimitCPU              float32
	LimitMEM              float32
}

type InferenceServiceActiveResponse struct {
	Namespace      string
	DeploymentID   string
	ModelHistoryID string
}

type InferenceServiceInActiveResponse struct {
	Namespace      string
	DeploymentID   string
	ModelHistoryID string
}
