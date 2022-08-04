package dto

type AccuracySetting struct {
	MetricType   string  `json:"metricType" example:"rmse" extensions:"x-order=0" enums:"rmse, rmsle, mae, mad, mape, mean_tweedie_deviance, gamma_deviance, tpr, accuracy, f1, ppv, fnr, fpr"` // Accuracy 측정 메트릭 종류
	Measurement  string  `json:"measurement" example:"percent" extensions:"x-order=1" enums:"percent, value"`                                                                                   // 메트릭 Value 타입
	AtRiskValue  float32 `json:"atRiskValue" example:"5" extensions:"x-order=2"`                                                                                                                // 메트릭의 AtRisk Value
	FailingValue float32 `json:"failingValue" example:"10" extensions:"x-order=3"`                                                                                                              // 메트릭의 Failing Value
}

type PatchAccuracySetting struct {
	MetricType   string   `json:"metricType" example:"rmse" extensions:"x-order=0" enums:"rmse, rmsle, mae, mad, mape, mean_tweedie_deviance, gamma_deviance, tpr, accuracy, f1, ppv, fnr, fpr"` // Accuracy 측정 메트릭 종류
	Measurement  string   `json:"measurement" example:"percent" extensions:"x-order=1" enums:"percent, value"`                                                                                   // 메트릭 Value 타입
	AtRiskValue  *float32 `json:"atRiskValue" example:"5" extensions:"x-order=2"`                                                                                                                // 메트릭의 AtRisk Value
	FailingValue *float32 `json:"failingValue" example:"10" extensions:"x-order=3"`                                                                                                              // 메트릭의 Failing Value
}

type MonitorAccuracyPatchRequestDTO struct {
	DeploymentID    string `json:"deploymentID" swaggerignore:"true"`
	AccuracySetting PatchAccuracySetting
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
	Message   string `json:"message" extensions:"x-order=0"`   // 결과 message
	Data      string `json:"data" extensions:"x-order=1"`      // 응답 결과 data
	StartTime string `json:"startTime" extensions:"x-order=2"` // 검색 시작 시간
	EndTime   string `json:"endTIme" extensions:"x-order=3"`   // 검색 끝 시간
}

type MonitorAccuracyActiveRequestDTO struct {
	DeploymentID   string
	ModelPackageID string
	AssociationID  *string
	CurrentModelID string
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

type UpdateAssociationIDRequestDTO struct {
	DeploymentID  string  `json:"deploymentID" swaggerignore:"true"`
	AssociationID *string `json:"associationID" example:"index" extensions:"x-order=0"`
}

type UpdateAssociationIDResponseDTO struct {
	DeploymentID string
	Message      string
}
