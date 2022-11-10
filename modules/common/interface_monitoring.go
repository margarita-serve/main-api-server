package common

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IMonitorService interface {
	GetByIDInternal(req *MonitorGetByIDInternalRequest) (*MonitorGetByIDInternalResponse, error)
	CreateMonitoring(req *CreateMonitoringRequest) error
	UpdateMonitoringOptions(req *UpdateMonitoringOptionsRequest) error
	DeleteMonitoring(req *DeleteMonitoringRequest) error
	ReplaceModelMonitoring(req *ReplaceModelMonitoringRequest) error
}

type MonitorGetByIDInternalRequest struct {
	DeploymentID string
}

type MonitorGetByIDInternalResponse struct {
	ID                     string
	ModelPackageID         string
	FeatureDriftTracking   bool
	AccuracyMonitoring     bool
	AssociationID          string
	AssociationIDInFeature bool
	DriftStatus            string
	AccuracyStatus         string
	ServiceHealthStatus    string
	DriftCreated           bool
	AccuracyCreated        bool
	ServiceHealthCreated   bool
}

type CreateMonitoringRequest struct {
	DeploymentID           string `json:"deploymentID"   `                                                                // Deployment ID
	ModelPackageID         string `json:"modelPackageID" `                                                                // ModelPackage ID
	FeatureDriftTracking   bool   `json:"featureDriftTracking" validate:"required" example:"true" extensions:"x-order=2"` // DataDrift Monitor 활성 여부
	AccuracyMonitoring     bool   `json:"accuracyMonitoring" validate:"required" example:"true" extensions:"x-order=3"`   // Accuracy Monitor 활성 여부
	AssociationID          string `json:"associationID" example:"index" extensions:"x-order=4"`                           // Accuracy Monitor 시 연결 ID
	AssociationIDInFeature bool   `json:"associationIDInFeature" example:"true" extensions:"x-order=5"`                   // Association In Feature 여부
	ModelHistoryID         string `json:"modelHistoryID" validate:"required" example:"000001" extensions:"x-order=6"`     // Monitor할 Model History ID
}

type DeleteMonitoringRequest struct {
	DeploymentID string
}

type UpdateMonitoringOptionsRequest struct {
	Identity               identity.Identity
	DeploymentID           string
	ModelPackageID         string
	ModelHistoryID         string
	FeatureDriftTracking   *bool
	AccuracyMonitoring     *bool
	AssociationID          *string
	AssociationIDInFeature *bool
}

type ReplaceModelMonitoringRequest struct {
	DeploymentID   string `json:"deploymentID"   ` // Deployment ID
	ModelPackageID string `json:"modelPackageID" ` // ModelPackage ID
	ModelHistoryID string `json:"modelHistoryID"`  // Monitor할 Model History ID
}

func (r *CreateMonitoringRequest) Validate() error {
	// if r.AccuracyMonitoring && r.AssociationID == "" {
	// 	return validation.ErrDateInvalid
	// }

	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.ModelPackageID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.FeatureDriftTracking, validation.NotNil, validation.In(true, false)),
		validation.Field(&r.AccuracyMonitoring, validation.NotNil, validation.In(true, false)),
		validation.Field(&r.ModelHistoryID, validation.Required, validation.NotNil),
	)
}

func (r *ReplaceModelMonitoringRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.ModelPackageID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.ModelHistoryID, validation.Required, validation.NotNil),
	)
}

func (r *MonitorGetByIDInternalRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
	)
}

func (r *DeleteMonitoringRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
	)
}

func (r *UpdateMonitoringOptionsRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
	)
}
