package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

func (r *MonitorAccuracyPatchRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
	)
}

func (r *AccuracyGetRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.ModelHistoryID, validation.Required, validation.NotNil),
		validation.Field(&r.Type, validation.Required, validation.NotNil, validation.In("timeline", "aggregation")),
		validation.Field(&r.StartTime, validation.Required, validation.NotNil),
		validation.Field(&r.EndTime, validation.Required, validation.NotNil),
	)
}

func (r *UploadActualRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.ActualResponse, validation.Required, validation.NotNil),
		validation.Field(&r.FileName, validation.Required, validation.NotNil),
		validation.Field(&r.FileName, validation.Required, validation.NotNil),
	)
}

func (r *PatchAccuracySetting) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.MetricType, validation.NotNil, validation.In("rmse", "rmsle", "mae", "mad", "mape", "mean_tweedie_deviance", "gamma_deviance", "tpr", "accuracy",
			"f1", "ppv", "fnr", "fpr")),
		validation.Field(&r.Measurement, validation.NotNil, validation.In("percent", "value")),
		validation.Field(&r.AtRiskValue, validation.Min(0.0)),
		validation.Field(&r.FailingValue, validation.Min(0.0)),
	)
}
