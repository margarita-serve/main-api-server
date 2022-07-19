package dto

type AccuracySetting struct {
	MetricType   string  `json:"metricType"`
	Measurement  string  `json:"measurement"`
	AtRiskValue  float32 `json:"atRiskValue"`
	FailingValue float32 `json:"failingValue"`
}

type MonitorAccuracyPatchRequestDTO struct {
	DeploymentID    string
	AccuracySetting AccuracySetting
}

type MonitorAccuracyPatchResponseDTO struct {
	DeploymentID string
	Message      string
}

type AccuracyGetRequestDTO struct {
	DeploymentID   string
	ModelHistoryID string
	Type           string
	StartTime      string
	EndTime        string
}

type AccuracyGetResponseDTO struct {
	Message   string
	Data      string
	StartTime string
	EndTime   string
}

type MonitorAccuracyActiveRequestDTO struct {
	DeploymentID    string
	ModelPackageID  string
	AssociationID   string
	AccuracySetting AccuracySetting
	CurrentModelID  string
}

type MonitorAccuracyActiveResponseDTO struct {
	DeploymentID string
}

type MonitorAccuracyInActiveRequestDTO struct {
	DeploymentID string
}

type MonitorAccuracyInActiveResponseDTO struct {
	Message string
}

type MonitorAccuracyGetSettingRequestDTO struct {
	DeploymentID string
}

type MonitorAccuracyGetSettingResponseDTO struct {
	AccuracySetting AccuracySetting
}
