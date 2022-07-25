package dto

type DataDriftCreateRequest struct {
	InferenceName              string
	ModelHistoryID             string
	TargetLabel                string
	ModelType                  string
	Framework                  string
	TrainDatasetPath           string
	ModelPath                  string
	DriftThreshold             float32
	ImportanceThreshold        float32
	MonitorRange               string
	LowImportanceAtRiskCount   int
	LowImportanceFailingCount  int
	HighImportanceAtRiskCount  int
	HighImportanceFailingCount int
}

type DataDriftCreateResponse struct {
	Message       string
	InferenceName string
}

type DataDriftDeleteRequest struct {
	InferenceName string
}

type DataDriftDeleteResponse struct {
	Message       string
	InferenceName string
}

type DataDriftGetRequest struct {
	InferenceName  string
	ModelHistoryID string
	StartTime      string
	EndTime        string
}

type DataDriftGetResponse struct {
	Message         string
	Data            string
	StartTime       string
	EndTime         string
	PredictionCount int
}

type DataDriftPatchRequest struct {
	InferenceName              string
	DriftThreshold             float32
	ImportanceThreshold        float32
	MonitorRange               string
	LowImportanceAtRiskCount   int
	LowImportanceFailingCount  int
	HighImportanceAtRiskCount  int
	HighImportanceFailingCount int
}

type DataDriftPatchResponse struct {
	Message       string
	InferenceName string
}

type DataDriftEnableRequest struct {
	InferenceName string
}

type DataDriftEnableResponse struct {
	Message       string
	InferenceName string
}
