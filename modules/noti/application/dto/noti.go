package dto

type NotiRequestDTO struct {
	DeploymentID   string
	NotiCategory   string //"Datadrift", "Accuracy", "Service"
	AdditionalData string //
}

type NotiResponseDTO struct {
	Message string
}
