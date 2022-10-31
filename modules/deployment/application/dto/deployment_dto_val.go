package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validate
func (r *CreateDeploymentRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		//validation.Field(&r.ProjectID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.ModelPackageID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.Name, validation.Length(0, 255)),
		validation.Field(&r.Description, validation.Length(0, 255)),
		validation.Field(&r.PredictionEnvID, validation.Length(20, 20)),
		validation.Field(&r.Importance, validation.In("Low", "Moderate", "High", "Critical")),
		validation.Field(&r.RequestCPU, validation.Min(0.1), validation.Max(2.0)),
		validation.Field(&r.RequestMEM, validation.Min(0.1), validation.Max(2.0)),
		validation.Field(&r.AssociationID, validation.When(r.AccuracyAnalyze == true, validation.Required)),
		validation.Field(&r.AssociationIDInFeature, validation.When(r.AccuracyAnalyze == true, validation.NotNil)),
	)
}

func (r *ReplaceModelRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.ModelPackageID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.Reason, validation.Required, validation.In("Accuracy", "DataDrift", "Errors", "ScheduledRefresh", "PredictionSpeed", "Other")),
		//validation.Field(&r.ManualApplication, validation.In("True", "False")),
	)
}

func (r *UpdateDeploymentRequestDTO) Validate() error {
	//var checkAcc bool
	//if r.AccuracyAnalyze != nil {
	//	checkAcc = *r.AccuracyAnalyze
	//}

	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.Name, validation.Length(0, 255)),
		validation.Field(&r.Description, validation.Length(0, 255)),
		validation.Field(&r.Importance, validation.In("Low", "Moderate", "High", "Critical")),
		validation.Field(&r.RequestCPU, validation.Min(0.1), validation.Max(2.0)),
		validation.Field(&r.RequestMEM, validation.Min(0.1), validation.Max(2.0)),
		//validation.Field(&r.AssociationID, validation.When(checkAcc == true, validation.Required, validation.NotNil)),
		//validation.Field(&r.AssociationIDInFeature, validation.When(checkAcc == true, validation.NotNil)),
	)
}

func (r *ActiveDeploymentRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
	)
}

func (r *GetDeploymentRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.DeploymentID, validation.Required, validation.NotNil, validation.Length(20, 20)),
	)
}

func (r *GetDeploymentListRequestDTO) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Name, validation.Length(0, 255)),
		validation.Field(&r.Page, validation.Min(0)),
		validation.Field(&r.Limit, validation.Min(0)),
		validation.Field(&r.Sort, validation.In("CreateDesc", "CreateASC")),
	)
}
