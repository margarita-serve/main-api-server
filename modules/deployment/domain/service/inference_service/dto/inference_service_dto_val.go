package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validate
func (r *InferenceServiceCreateRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.ConnectionInfo, validation.Required),
		validation.Field(&r.DeploymentID, validation.Required),
		validation.Field(&r.ModelURL, validation.Required),
		validation.Field(&r.ModelFrameWork, validation.Required),
		validation.Field(&r.Namespace, validation.Required),
	)
}
