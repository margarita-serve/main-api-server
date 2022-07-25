package types

type GetDetailGraphRequest struct {
	InferenceName  string `json:"inferencename"`
	ModelHistoryID string `json:"model_history_id"`
	StartTime      string `json:"start_time"`
	EndTime        string `json:"end_time"`
}
