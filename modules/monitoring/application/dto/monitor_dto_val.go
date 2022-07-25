package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

func (r *MonitorCreateRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.ModelPackageID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.FeatureDriftTracking, validation.Required, validation.NotNil, validation.In(true, false)),
		validation.Field(&r.AccuracyMonitoring, validation.Required, validation.NotNil, validation.In(true, false)),
		validation.Field(&r.ModelHistoryID, validation.Required, validation.NotNil),
	)
}

func (r *MonitorReplaceModelRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.ModelPackageID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.ModelHistoryID, validation.Required, validation.NotNil),
	)
}

func (r *MonitorGetByIDRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.ID, validation.Required, validation.NotNil, validation.Length(20, 20)),
	)
}

func (r *MonitorGetSettingRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
	)
}
