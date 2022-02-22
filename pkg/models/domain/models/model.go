package models

type ModelId string
type TargetType string //[‘Binary’, ‘Regression’, ‘Multiclass’, ‘Anomaly’, ‘Transform’, ‘Unstructured’]

type Model struct {
	id                  ModelId
	name                string
	description         string
	language            string
	targetType          TargetType
	positiveClassLabel  string
	negativeClassLabel  string
	classLabels         string
	predictionThreshold string
}
