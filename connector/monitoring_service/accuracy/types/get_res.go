package types

type GetAccuracyResponse struct {
	Message   string `json:"message"`
	Data      string `json:"data"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
