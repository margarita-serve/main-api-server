package dto

type AccuracyCreateRequest struct {
	InferenceName    string
	ModelHistoryID   string
	DatasetPath      string
	ModelPath        string
	TargetLabel      string
	AssociationID    string
	ModelType        string
	Framework        string
	DriftMetrics     string
	DriftMeasurement string
	AtriskValue      float32
	FailingValue     float32
	PositiveClass    string
	NegativeClass    string
	BinaryThreshold  float32
}

type AccuracyCreateResponse struct {
	Message       string
	InferenceName string
}

type AccuracyDeleteRequest struct {
	InferenceName string
}

type AccuracyDeleteResponse struct {
	Message       string
	InferenceName string
}

type AccuracyPatchRequest struct {
	InferenceName    string
	DriftMetrics     string
	DriftMeasurement string
	AtriskValue      float32
	FailingValue     float32
}

type AccuracyPatchResponse struct {
	Message       string
	InferenceName string
}

type AccuracyGetRequest struct {
	InferenceName  string
	ModelHistoryID string
	DataType       string
	StartTime      string
	EndTime        string
}

type AccuracyGetResponse struct {
	Message   string
	Data      string
	StartTime string
	EndTIme   string
}

type AccuracyPostActualRequest struct {
	InferenceName  string
	DatasetPath    string
	ActualResponse string
}

type AccuracyPostActualResponse struct {
	Message       string `json:"message"`
	InferenceName string `json:"inferencename"`
}
