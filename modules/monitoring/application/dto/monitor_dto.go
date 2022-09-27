package dto

import (
	domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/entity"
	"io"
)

type MonitorCreateRequestDTO struct {
	DeploymentID           string  `json:"deploymentID" validate:"required" example:"cbgfbddvqc7mecjqbc9g" extensions:"x-order=0" swaggerignore:"true"` // Deployment ID
	ModelPackageID         string  `json:"modelPackageID" validate:"required" example:"cbbrc45vqc7ks0qlldfg" extensions:"x-order=1"`                    // ModelPackage ID
	FeatureDriftTracking   bool    `json:"featureDriftTracking" validate:"required" example:"true" extensions:"x-order=2"`                              // DataDrift Monitor 활성 여부
	AccuracyMonitoring     bool    `json:"accuracyMonitoring" validate:"required" example:"true" extensions:"x-order=3"`                                // Accuracy Monitor 활성 여부
	AssociationID          *string `json:"associationID" example:"index" extensions:"x-order=4"`                                                        // Accuracy Monitor 시 연결 ID
	AssociationIDInFeature bool    `json:"associationIDInFeature" example:"true" extensions:"x-order=5"`                                                // Association In Feature 여부
	ModelHistoryID         string  `json:"modelHistoryID" validate:"required" example:"000001" extensions:"x-order=6"`                                  // Monitor할 Model History ID
}

type MonitorStatus struct {
	DeploymentID        string
	DriftStatus         string
	AccuracyStatus      string
	ServiceHealthStatus string
}

type MonitorCreateResponseDTO struct {
	DeploymentID string `json:"deploymentID"`
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

type MonitorGetSettingRequestDTO struct {
	DeploymentID string `json:"deploymentID"`
}

type MonitorGetSettingResponseDTO struct {
	DataDriftSetting DataDriftSetting
	AccuracySetting  AccuracySetting
}

type UploadActualRequestDTO struct {
	DeploymentID      string    `json:"deploymentID" swaggerignore:"true"`
	ActualResponse    string    `json:"actualResponse"`
	AssociationColumn string    `json:"associationColumn"`
	File              io.Reader `swaggerignore:"true"`
	FileName          string    `swaggerignore:"true"`
}

type UploadActualResponseDTO struct {
	DeploymentID string
	Message      string
}

type MonitorReplaceModelRequestDTO struct {
	DeploymentID   string
	ModelPackageID string
	ModelHistoryID string
}

type MonitorReplaceModelResponseDTO struct {
	DeploymentID string
}

type MonitorStatusCheckRequestDTO struct {
	DeploymentID string
	Status       string
	Kind         string
}

type MonitorPatchRequestDTO struct {
	DeploymentID     string `json:"deploymentID" swaggerignore:"true"`
	AccuracySetting  PatchAccuracySetting
	DataDriftSetting PatchDataDriftSetting
}

type MonitorPatchResponseDTO struct {
	DeploymentID string
	Message      string
}
