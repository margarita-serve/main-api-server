package types

import "encoding/json"

type CreateAccuracyRequest struct {
	DatasetPath      string  `json:"dataset_path"`
	ModelPath        string  `json:"model_path"`
	InferenceName    string  `json:"inference_name"`
	ModelId          string  `json:"model_id"`
	TargetLabel      string  `json:"target_label"`
	AssociationId    string  `json:"association_id"`
	ModelType        string  `json:"model_type"`
	Framework        string  `json:"framework"`
	DriftMetric      string  `json:"drift_metric"`
	DriftMeasurement string  `json:"drift_measurement"`
	AtriskValue      float32 `json:"atrisk_value"`
	FailingValue     float32 `json:"failing_value"`
	PositiveClass    string  `json:"positive_class"`
	NegativeClass    string  `json:"negative_class"`
	BinaryThreshold  float32 `json:"binary_threshold"`
}

func (r *CreateAccuracyRequest) ToJSON() []byte {
	Mjson, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		return nil
	}
	return Mjson
}
