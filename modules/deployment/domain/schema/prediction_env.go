package schema

type PredictionEnv struct {
	Namespace               string
	ConnectionInfo          string
	InfereceSvcHostName     string
	InferenceSvcIngressHost string
	InferenceSvcIngressPort string
}
