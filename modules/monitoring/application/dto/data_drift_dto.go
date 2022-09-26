package dto

type DataDriftSetting struct {
	MonitorRange               string  `json:"monitorRange" example:"7d" extensions:"x-order=0" enums:"2h, 1d, 7d, 30d, 90d, 180d, 365d"` // Monitoring 할 범위
	DriftMetricType            string  `json:"driftMetricType" example:"PSI" extensions:"x-order=1" enums:"PSI"`                          // DataDrift 측정 Metric
	DriftThreshold             float32 `json:"driftThreshold" example:"0.15" extensions:"x-order=2"`                                      // Drift 값 임계치
	ImportanceThreshold        float32 `json:"importanceThreshold" example:"0.5" extensions:"x-order=3"`                                  // Importance 값 임계치
	LowImportanceAtRiskCount   int     `json:"lowImportanceAtRiskCount" example:"1" extensions:"x-order=4"`                               // 낮은 importance feature 수의 drift at risk 임계치
	LowImportanceFailingCount  int     `json:"lowImportanceFailingCount" example:"0" extensions:"x-order=5"`                              // 낮은 importance feature 수의 drift failing 임계치
	HighImportanceAtRiskCount  int     `json:"highImportanceAtRiskCount" example:"0" extensions:"x-order=6"`                              // 높은 importance feature 수의 drift at risk 임계치
	HighImportanceFailingCount int     `json:"highImportanceFailingCount" example:"1" extensions:"x-order=7"`                             // 높은 importance feature 수의 drift failing 임계치
}

type PatchDataDriftSetting struct {
	MonitorRange               string   `json:"monitorRange" example:"7d" extensions:"x-order=0" enums:"2h, 1d, 7d, 30d, 90d, 180d, 365d"` // Monitoring 할 범위
	DriftMetricType            string   `json:"driftMetricType" example:"PSI" extensions:"x-order=1" enums:"PSI"`                          // DataDrift 측정 Metric
	DriftThreshold             *float32 `json:"driftThreshold" example:"0.15" extensions:"x-order=2"`                                      // Drift 값 임계치
	ImportanceThreshold        *float32 `json:"importanceThreshold" example:"0.5" extensions:"x-order=3"`                                  // Importance 값 임계치
	LowImportanceAtRiskCount   *int     `json:"lowImportanceAtRiskCount" example:"1" extensions:"x-order=4"`                               // 낮은 importance feature 수의 drift at risk 임계치
	LowImportanceFailingCount  *int     `json:"lowImportanceFailingCount" example:"0" extensions:"x-order=5"`                              // 낮은 importance feature 수의 drift failing 임계치
	HighImportanceAtRiskCount  *int     `json:"highImportanceAtRiskCount" example:"0" extensions:"x-order=6"`                              // 높은 importance feature 수의 drift at risk 임계치
	HighImportanceFailingCount *int     `json:"highImportanceFailingCount" example:"1" extensions:"x-order=7"`                             // 높은 importance feature 수의 drift failing 임계치
}

type MonitorDriftPatchRequestDTO struct {
	DeploymentID     string `json:"deploymentID" swaggerignore:"true"`
	DataDriftSetting PatchDataDriftSetting
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
	Message         string `json:"message" extensions:"x-order=0"`         // Message
	Data            string `json:"data" extensions:"x-order=1"`            // Feature Drift 결과 Data
	StartTime       string `json:"startTIme" extensions:"x-order=2"`       // 검색 시작 시간
	EndTime         string `json:"endTIme" extensions:"x-order=3"`         // 검색 끝 시간
	PredictionCount int    `json:"predictionCount" extensions:"x-order=4"` // 검색 시간 사이의 총 예측 수
}

type MonitorDriftActiveRequestDTO struct {
	DeploymentID   string
	ModelPackageID string
	CurrentModelID string
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
