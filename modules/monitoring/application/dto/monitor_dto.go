package dto

import (
	domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/entity"
	"io"
)

type MonitorCreateRequestDTO struct {
	DeploymentID         string `json:"deploymentID" swaggerignore:"true"`
	ModelPackageID       string `json:"modelPackageID"`
	FeatureDriftTracking bool   `json:"featureDriftTracking"`
	AccuracyMonitoring   bool   `json:"accuracyMonitoring"`
	AssociationID        string `json:"associationID"`
	DataDriftSetting     DataDriftSetting
	AccuracySetting      AccuracySetting
	ServiceHealthSetting ServiceHealthSetting
}

type MonitorStatus struct {
	DeploymentID        string
	DriftStatus         string
	AccuracyStatus      string
	ServiceHealthStatus string
}

type MonitorCreateResponseDTO struct {
	DeploymentID string
}

type MonitorDeleteRequestDTO struct {
	DeploymentID string
}

type MonitorDeleteResponseDTO struct {
	Message string
}

type MonitorGetByIDRequestDTO struct {
	ID string
}

type MonitorGetByIDResponseDTO struct {
	Monitor *domEntity.Monitor
}

type MonitorPatchRequestDTO struct {
	DeploymentID     string `json:"deploymentID" swaggerignore:"true"`
	DataDriftSetting DataDriftSetting
	AccuracySetting  AccuracySetting
}

type MonitorGetSettingRequestDTO struct {
	DeploymentID string `json:"deploymentID"`
}

type MonitorGetSettingResponseDTO struct {
	DataDriftSetting DataDriftSetting
	AccuracySetting  AccuracySetting
}

type UploadActualRequestDTO struct {
	DeploymentID   string    `json:"deploymentID" swaggerignore:"true"`
	ProjectID      string    `json:"projectID" swaggerignore:"true"`
	ActualResponse string    `json:"actualResponse"`
	File           io.Reader `swaggerignore:"true"`
	FileName       string    `swaggerignore:"true"`
}

type UploadActualResponseDTO struct {
	DeploymentID string
	Message      string
}

type MonitorGetStatusListRequestDTO struct {
	DeploymentsID string `json:"deploymentsID"`
}

type MonitorGetStatusListResponseDTO struct {
	StatusList []MonitorStatus
}
