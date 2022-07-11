package entity

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var ErrDeploymentID = errors.New("empty deployment id")

type ModelPackage struct {
	ID                    string `gorm:"size:256"`
	ProjectID             string `gorm:"size:256"`
	Name                  string `gorm:"size:256"`
	Description           string `gorm:"size:256"`
	ModelName             string `gorm:"size:256"`
	ModelVersion          string `gorm:"size:256"`
	ModelDescription      string `gorm:"size:256"`
	TargetType            string `gorm:"size:256"`
	PredictionTargetName  string `gorm:"size:256"`
	ModelFrameWork        string `gorm:"size:256"`
	ModelFrameWorkVersion string `gorm:"size:256"`
	PredictionThreshold   float32
	PositiveClassLabel    string `gorm:"size:256"`
	NegativeClassLabel    string `gorm:"size:256"`
	ModelFilePath         string
	TrainingDatasetPath   string
	HoldoutDatasetPath    string
	Archived              bool
	DeployCount           int
	BaseEntity
}

// Validate
func Validate(r *ModelPackage) error {
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
		validation.Field(&r.PredictionThreshold, validation.When(r.TargetType == "Binary", validation.Required).Else(validation.Empty)),
		validation.Field(&r.PositiveClassLabel, validation.When(r.TargetType == "Binary", validation.Required).Else(validation.Empty)),
		validation.Field(&r.NegativeClassLabel, validation.When(r.TargetType == "Binary", validation.Required).Else(validation.Empty)),
	)
}

func NewModelPackage(id string, projectID string, name string, description string, modelName string, modelVersion string, modelDescription string, targetType string, predictionTargetName string, modelFrameWork string, modelFrameWorkVersion string, predictionThreshold float32, positiveClassLabel string, negativeClassLabel string, ownerID string) (*ModelPackage, error) {

	var baseEntity BaseEntity
	baseEntity.CreatedBy = "testuser"

	modelPackage := &ModelPackage{
		ID:                    id,
		ProjectID:             projectID,
		Name:                  name,
		Description:           description,
		ModelName:             modelName,
		ModelVersion:          modelVersion,
		ModelDescription:      modelDescription,
		TargetType:            targetType,
		PredictionTargetName:  predictionTargetName,
		ModelFrameWork:        modelFrameWork,
		ModelFrameWorkVersion: modelFrameWorkVersion,
		PredictionThreshold:   predictionThreshold,
		PositiveClassLabel:    positiveClassLabel,
		NegativeClassLabel:    negativeClassLabel,
		DeployCount:           0,
		BaseEntity:            baseEntity,
	}

	// Validate
	err := Validate(modelPackage)
	if err != nil {
		return nil, err
	}

	return modelPackage, nil
}

func (d *ModelPackage) SetName(req string) {
	d.Name = req
}

func (d *ModelPackage) SetDescription(req string) {
	d.Description = req
}

func (d *ModelPackage) SetModelName(req string) {
	d.ModelName = req
}

func (d *ModelPackage) SetModelVersion(req string) {
	d.ModelVersion = req
}

func (d *ModelPackage) SetModelDescription(req string) {
	d.ModelDescription = req
}

func (d *ModelPackage) SetTargetType(targetType string) {
	d.TargetType = targetType
}

func (d *ModelPackage) SetPredictionTargetName(targetName string) {
	d.PredictionTargetName = targetName
}

func (d *ModelPackage) SetModelFrameWork(modelFrameWork string) {
	d.ModelFrameWork = modelFrameWork
}

func (d *ModelPackage) SetModelFrameWorkVersion(modelFrameWorkVersion string) {
	d.ModelFrameWorkVersion = modelFrameWorkVersion
}

func (d *ModelPackage) SetPredictionThreshold(threshold float32) {
	d.PredictionThreshold = threshold
}

func (d *ModelPackage) SetPositiveClassLabel(label string) {
	d.PositiveClassLabel = label
}

func (d *ModelPackage) SetNegativeClassLabel(label string) {
	d.NegativeClassLabel = label
}

func (d *ModelPackage) SetModelPath(path string) {
	d.ModelFilePath = path
}

func (d *ModelPackage) SetTraningDatasetPath(path string) {
	d.TrainingDatasetPath = path
}

func (d *ModelPackage) SetHoldoutDatasetPath(path string) {
	d.HoldoutDatasetPath = path
}

func (d *ModelPackage) SetArchived() {
	d.Archived = true
}

func (d *ModelPackage) AddDeployCount() {
	d.DeployCount++
}

func (d *ModelPackage) IsValidForDelete() bool {
	if d.DeployCount <= 0 && !d.Archived {
		return true
	}
	return false
}

func (d *ModelPackage) IsValidForUpdate() bool {
	return !d.Archived
}
