package dto

type DetailGraphGetRequest struct {
	InferenceName  string
	ModelHistoryID string
	StartTime      string
	EndTime        string
	HostEndpoint   string
}

type DetailGraphGetResponse struct {
	Script string
}

type DriftGraphGetRequest struct {
	InferenceName       string
	ModelHistoryID      string
	StartTime           string
	EndTime             string
	HostEndpoint        string
	DriftThreshold      float32
	ImportanceThreshold float32
}

type DriftGraphGetResponse struct {
	Script string
}
