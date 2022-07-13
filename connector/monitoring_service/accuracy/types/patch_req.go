package types

import "encoding/json"

type PatchAccuracySettingRequest struct {
	InferenceName    string  `json:"inference_name"`
	DriftMetrics     string  `json:"drift_metrics"`
	DriftMeasurement string  `json:"drift_measurement"`
	AtriskValue      float32 `json:"atrisk_value"`
	FailingValue     float32 `json:"failing_value"`
}

type PatchAccuracySettingRequestDTO struct {
	DriftMetrics     string  `json:"drift_metrics"`
	DriftMeasurement string  `json:"drift_measurement"`
	AtriskValue      float32 `json:"atrisk_value"`
	FailingValue     float32 `json:"failing_value"`
}

func (r *PatchAccuracySettingRequestDTO) ToJSON() []byte {
	Mjson, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		return nil
	}
	return Mjson
}
