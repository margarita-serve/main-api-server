package dto

type DataDriftSetting struct {
	MonitorRange               string  `json:"monitorRange"`
	DriftMetricType            string  `json:"driftMetricType"`
	DriftThreshold             float32 `json:"driftThreshold"`
	ImportanceThreshold        float32 `josn:"importanceThreshold"`
	LowImportanceAtRiskCount   int     `json:"lowImportanceAtRiskCount"`
	LowImportanceFailingCount  int     `json:"lowImportanceFailingCount"`
	HighImportanceAtRiskCount  int     `json:"highImportanceAtRiskCount"`
	HighImportanceFailingCount int     `json:"highImportanceFailingCount"`
}

type MonitorDriftPatchRequestDTO struct {
	DeploymentID     string
	DataDriftSetting DataDriftSetting
}

type MonitorDriftPatchResponseDTO struct {
	DeploymentID string
	Message      string
}

type FeatureDriftGetRequestDTO struct {
	DeploymentID   string
	ModelHistoryID string
	StartTime      string
	EndTime        string
}

type FeatureDriftGetResponseDTO struct {
	Message         string
	Data            string
	StartTime       string
	EndTime         string
	PredictionCount int
}

type MonitorDriftActiveRequestDTO struct {
	DeploymentID     string
	ModelPackageID   string
	DataDriftSetting DataDriftSetting
	CurrentModelID   string
}

type MonitorDriftActiveResponseDTO struct {
	DeploymentID string
}

type MonitorDriftInActiveRequestDTO struct {
	DeploymentID string
}

type MonitorDriftInActiveResponseDTO struct {
	Message string
}

type MonitorDriftGetSettingRequestDTO struct {
	DeploymentID string
}

type MonitorDriftGetSettingResponse struct {
	DataDriftSetting DataDriftSetting
}
