package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

func (r *DriftGraphGetRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.ModelHistoryID, validation.Required, validation.NotNil),
		validation.Field(&r.StartTime, validation.Required, validation.NotNil),
		validation.Field(&r.EndTime, validation.Required, validation.NotNil),
	)
}

func (r *DetailGraphGetRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.ModelHistoryID, validation.Required, validation.NotNil),
		validation.Field(&r.StartTime, validation.Required, validation.NotNil),
		validation.Field(&r.EndTime, validation.Required, validation.NotNil),
	)
}
