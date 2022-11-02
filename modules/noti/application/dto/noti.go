package dto

type NotiRequestDTO struct {
	DeploymentID   string
	NotiCategory   string //"DataDrift", "Accuracy", "Service"
	AdditionalData string //
}

type NotiResponseDTO struct {
	Message string
}
