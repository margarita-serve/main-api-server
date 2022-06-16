package entity

import (
	"errors"
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
	BaseEntity
}

func NewModelPackage(id string, projectID string, name string, description string, modelName string, modelVersion string, modelDescription string, targetType string, predictionTargetName string, modelFrameWork string, modelFrameWorkVersion string, predictionThreshold float32, positiveClassLabel string, negativeClassLabel string, ownerID string) (*ModelPackage, error) {

	var be BaseEntity
	be.CreatedBy = ownerID

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
		BaseEntity:            be,
	}

	return modelPackage, nil
}
