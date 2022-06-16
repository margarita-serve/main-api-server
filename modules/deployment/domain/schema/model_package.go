package schema

type ModelPackage struct {
	TargetType           string
	PredictionTargetName string
	PositiveClassLabel   string
	ModelFrameWork       string
	ModelURL             string
	ModelName            string
	TrainingDatasetURL   string
	HoldoutDatasetURL    string
}
