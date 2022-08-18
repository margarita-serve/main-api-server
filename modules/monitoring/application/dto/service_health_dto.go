package dto

type MonitorServiceHealthActiveRequestDTO struct {
	DeploymentID   string
	CurrentModelID string
}

type MonitorServiceHealthActiveResponseDTO struct {
	DeploymentID string
}

type MonitorServiceHealthInActiveRequestDTO struct {
	DeploymentID string
}

type MonitorServiceHealthInActiveResponseDTO struct {
	Message string
}
