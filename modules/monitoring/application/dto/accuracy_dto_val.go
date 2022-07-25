package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

func (r *MonitorAccuracyPatchRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.AccuracySetting.MetricType, validation.Required, validation.NotNil, validation.In("rmse", "rmsle", "mae", "mad", "mape", "mean_tweedie_deviance", "gamma_deviance", "tpr", "accuracy",
			"f1", "ppv", "fnr", "fpr")),
		validation.Field(&r.AccuracySetting.Measurement, validation.Required, validation.NotNil, validation.In("percent", "value")),
		validation.Field(&r.AccuracySetting.AtRiskValue, validation.Required, validation.NotNil, validation.Min(0)),
		validation.Field(&r.AccuracySetting.FailingValue, validation.Required, validation.NotNil, validation.Min(0)),
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
