package dto

type NotiRequestDTO struct {
	DeploymentID string
	NotiCategory string
	Data         string
}

type NotiResponseDTO struct {
	Message string
}
