package dto

type ModelPackageCreateRequestDTO struct {
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
}

type ModelPackageCreateResponseDTO struct {
	ModelPackageID string
}

type ModelPackageGetRequestDTO struct {
	ModelPackageID string
}

type ModelPackageGetResponseDTO struct {
	ID                    string
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
}

type ModelPackageGetByNametRequestDTO struct {
	Name string
}

type ModelPackageGetByNameResponseDTO struct {
	ModelPackages []*ModelPackage
}

type ModelPackage struct {
	ID                    string
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
}

type ModelPackageDeleteRequestDTO struct {
	ModelPackageID string
}

type ModelPackageDeleteResponseDTO struct {
	Message string
}

type ModelPackageActiveRequestDTO struct {
	ModelPackageID string
}

type ModelPackageActiveResponseDTO struct {
	Message string
}

type ModelPackageGetInternalRequestDTO struct {
	ModelPackageID string
}

type ModelPackageGetInternalResponseDTO struct {
	ID                    string
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
	ModelURL              string
	TrainingDatasetURL    string
	HoldoutDatasetURL     string
}
