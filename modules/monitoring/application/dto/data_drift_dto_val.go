package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

func (r *MonitorDriftPatchRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.DataDriftSetting.DriftThreshold, validation.Required, validation.NotNil, validation.Min(0), validation.Max(1)),
		validation.Field(&r.DataDriftSetting.ImportanceThreshold, validation.Required, validation.NotNil, validation.Min(0), validation.Max(1)),
		validation.Field(&r.DataDriftSetting.DriftMetricType, validation.Required, validation.NotNil, validation.In("PSI")),
		validation.Field(&r.DataDriftSetting.LowImportanceAtRiskCount, validation.Required, validation.NotNil, validation.Min(0)),
		validation.Field(&r.DataDriftSetting.LowImportanceFailingCount, validation.Required, validation.NotNil, validation.Min(0)),
		validation.Field(&r.DataDriftSetting.HighImportanceAtRiskCount, validation.Required, validation.NotNil, validation.Min(0)),
		validation.Field(&r.DataDriftSetting.HighImportanceFailingCount, validation.Required, validation.NotNil, validation.Min(0)),
		validation.Field(&r.DataDriftSetting.MonitorRange, validation.Required, validation.NotNil, validation.In("2h", "1d", "7d", "30d", "90d", "180d", "365d")),
	)
}

func (r *FeatureDriftGetRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.ModelHistoryID, validation.Required, validation.NotNil),
		validation.Field(&r.DriftThreshold, validation.Required, validation.NotNil, validation.Min(0), validation.Max(1)),
		validation.Field(&r.ImportanceThreshold, validation.Required, validation.NotNil, validation.Min(0), validation.Max(1)),
		validation.Field(&r.StartTime, validation.Required, validation.NotNil),
		validation.Field(&r.EndTime, validation.Required, validation.NotNil),
	)
}
