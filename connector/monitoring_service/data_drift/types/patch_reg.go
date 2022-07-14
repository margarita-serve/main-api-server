package types

import "encoding/json"

type PatchDriftMonitorSettingRequest struct {
	InferenceName       string  `json:"inference_name"`
	DriftThreshold      float32 `json:"drift_threshold"`
	ImportanceThreshold float32 `json:"importance_threshold"`
	MonitorRange        string  `json:"monitor_range"`
	LowImpAtRiskCount   int     `json:"low_imp_atrisk_count"`
	LowImpFailingCount  int     `json:"low_imp_failing_count"`
	HighImpAtRiskCount  int     `json:"high_imp_atrisk_count"`
	HighImpFailingCount int     `json:"high_imp_failing_count"`
}

type PatchDriftMonitorSettingRequestDTO struct {
	DriftThreshold      float32 `json:"drift_threshold"`
	ImportanceThreshold float32 `json:"importance_threshold"`
	MonitorRange        string  `json:"monitor_range"`
	LowImpAtRiskCount   int     `json:"low_imp_atrisk_count"`
	LowImpFailingCount  int     `json:"low_imp_failing_count"`
	HighImpAtRiskCount  int     `json:"high_imp_atrisk_count"`
	HighImpFailingCount int     `json:"high_imp_failing_count"`
}

func (r *PatchDriftMonitorSettingRequestDTO) ToJSON() []byte {
	Mjson, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		return nil
	}
	return Mjson
}
