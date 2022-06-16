package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validate
func (r *ModelPackageCreateRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.ProjectID, validation.Required),
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.Description, validation.Required),
		validation.Field(&r.ModelName),
		validation.Field(&r.ModelVersion),
		validation.Field(&r.ModelDescription),
		validation.Field(&r.TargetType, validation.Required, validation.In("Binary", "Regression")),
		validation.Field(&r.PredictionTargetName),
		validation.Field(&r.ModelFrameWork, validation.Required, validation.In("TensorFlow", "PyTorch", "SkLearn", "XGBoost", "LightGBM")),
		validation.Field(&r.ModelFrameWorkVersion, validation.Required),
		validation.Field(&r.PredictionThreshold),
		validation.Field(&r.PositiveClassLabel),
		validation.Field(&r.NegativeClassLabel),
	)
}

func (r *ModelPackageGetRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.ModelPackageID, validation.Required, validation.NotNil, validation.Length(20, 20)),
	)
}
