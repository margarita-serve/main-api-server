package model_package

import (
	"time"
)

type ModelPackageID string
type ProjectID string

type TargetType string //[‘Binary’, ‘Regression’, ‘Multiclass’, ‘Anomaly’, ‘Transform’, ‘Unstructured’]
const (
	Binary       TargetType = "binary"
	Regression   TargetType = "regression"
	Multiclass   TargetType = "multiclass"
	Anomaly      TargetType = "anomaly"
	Transform    TargetType = "anomaly"
	Unstructured TargetType = "unstructured"
)

type UserID string

type ModelPackage struct {
	ID                   ModelPackageID `gorm:"primarykey"`
	ProjectID            ProjectID
	Name                 string
	Description          string
	Archived             bool
	TargetType           TargetType
	PredictionTargetName string
	PositiveClassLabel   string
	NegativeClassLabel   string
	PredictionThreshold  int
	ClassLabels          string
	RuntimeFrameWork     string
	CreatedAt            time.Time
	CreatedBy            UserID
	Model                Model             `gorm:"embedded;embeddedPrefix:model_"`
	TrainingDataset      TrainingDataset   `gorm:"embedded;embeddedPrefix:training_dataset_"`
	HoldoutDataset       HoldoutDataset    `gorm:"embedded;embeddedPrefix:holdout_dataset_"`
	TransformerFiles     []TransformerFile `gorm:"foreignKey:ID"`
}
