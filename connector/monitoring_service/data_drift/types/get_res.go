package types

type GetFeatureDriftResponse struct {
	Message         string `json:"message"`
	Data            string `json:"data"`
	StartTime       string `json:"start_time"`
	EndTime         string `json:"end_time"`
	PredictionCount int    `json:"prediction_count"`
}
