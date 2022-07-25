package dto

type DetailGraphGetRequestDTO struct {
	DeploymentID   string
	ModelHistoryID string
	StartTime      string
	EndTime        string
}

type DetailGraphGetResponseDTO struct {
	Script string `json:"script"` // graph JS script
}

type DriftGraphGetRequestDTO struct {
	DeploymentID   string
	ModelHistoryID string
	StartTime      string
	EndTime        string
}

type DriftGraphGetResponseDTO struct {
	Script string `json:"script"` // graph JS script
}
