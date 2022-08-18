package types

type GetServiceHealthRequest struct {
	InferenceName  string `json:"inferencename"`
	ModelHistoryID string `json:"model_history_id"`
	DataType       string `json:"type"`
	StartTime      string `json:"start_time"`
	EndTime        string `json:"end_time"`
}
