package types

import "encoding/json"

type CreateDataDriftRequest struct {
	TrainDatasetPath    string  `json:"train_dataset_path"`
	ModelPath           string  `json:"model_path"`
	InferenceName       string  `json:"inference_name"`
	ModelID             string  `json:"model_id"`
	TargetLabel         string  `json:"target_label"`
	ModelType           string  `json:"model_type"`
	Framework           string  `json:"framework"`
	DriftThreshold      float32 `json:"drift_threshold"`
	ImportanceThreshold float32 `json:"importance_threshold"`
	MonitorRange        string  `json:"monitor_range"`
	LowImpAtRiskCount   int     `json:"low_imp_atrisk_count"`
	LowImpFailingCount  int     `json:"low_imp_failing_count"`
	HighImpAtRiskCount  int     `json:"high_imp_atrisk_count"`
	HighImpFailingCount int     `json:"high_imp_failing_count"`
}

func (r *CreateDataDriftRequest) ToJSON() []byte {
	Mjson, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		return nil
	}
	return Mjson
}
