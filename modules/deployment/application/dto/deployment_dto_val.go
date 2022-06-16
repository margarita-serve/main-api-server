package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validate
func (r *DeploymentCreateRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.ProjectID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.ModelPackageID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.Name, validation.Length(0, 255)),
		validation.Field(&r.Description, validation.Length(0, 255)),
		validation.Field(&r.PredictionEnvID, validation.Required, validation.NotNil, validation.Length(12, 12)),
		validation.Field(&r.Importance, validation.In("Low", "Moderate", "High", "Critical")),
		validation.Field(&r.RequestCPU, validation.Length(12, 12)),
		validation.Field(&r.RequestMEM, validation.Length(12, 12)),
		validation.Field(&r.RequestGPU, validation.Length(12, 12)),
	)
}

func (r *DeploymentActiveRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
	)
}

func (r *DeploymentGetRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
	)
}
