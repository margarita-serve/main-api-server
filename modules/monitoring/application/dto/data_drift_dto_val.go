package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

func (r *MonitorDriftPatchRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
	)
}

func (r *FeatureDriftGetRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.ModelHistoryID, validation.Required, validation.NotNil),
		validation.Field(&r.StartTime, validation.Required, validation.NotNil),
		validation.Field(&r.EndTime, validation.Required, validation.NotNil),
	)
}

func (r *PatchDataDriftSetting) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DriftThreshold, validation.Min(0.0), validation.Max(1.0)),
		validation.Field(&r.ImportanceThreshold, validation.Min(0.0), validation.Max(1.0)),
		validation.Field(&r.DriftMetricType, validation.In("PSI")),
		validation.Field(&r.LowImportanceAtRiskCount, validation.Min(0)),
		validation.Field(&r.LowImportanceFailingCount, validation.Min(0)),
		validation.Field(&r.HighImportanceAtRiskCount, validation.Min(0)),
		validation.Field(&r.HighImportanceFailingCount, validation.Min(0)),
		validation.Field(&r.MonitorRange, validation.In("2h", "1d", "7d", "30d", "90d", "180d", "365d")),
	)
}
