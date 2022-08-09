package types

type GetDriftGraphRequest struct {
	InferenceName       string  `json:"inferencename"`
	ModelHistoryID      string  `json:"model_history_id"`
	StartTime           string  `json:"start_time"`
	EndTime             string  `json:"end_time"`
	HostEndpoint        string  `json:"host_endpoint"`
	DriftThreshold      float32 `json:"drift_threshold"`
	ImportanceThreshold float32 `json:"importance_threshold"`
}
