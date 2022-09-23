package common

type InternalGetModelPackageResponseDTO struct {
	ModelPackageID        string
	ProjectID             string
	Name                  string
	Description           string
	ModelName             string
	ModelVersion          string
	ModelDescription      string
	TargetType            string
	PredictionTargetName  string
	ModelFrameWork        string
	ModelFrameWorkVersion string
	PredictionThreshold   float32
	PositiveClassLabel    string
	NegativeClassLabel    string
	ModelFilePath         string
	TrainingDatasetPath   string
	HoldoutDatasetPath    string
	Archived              bool
}

type IModelPackageService interface {
	GetByIDInternal(modelPackageID string) (*InternalGetModelPackageResponseDTO, error)
}
