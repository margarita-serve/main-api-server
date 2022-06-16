package dto

type DeploymentCreateRequestDTO struct {
	ProjectID       string `json:"projectID" validate:"required" example:"ID-1234" extensions:"x-order=0 x-nullable=false"`          // 프로젝트 ID
	ModelPackageID  string `json:"modelPackageID" validate:"required" example:"ID-1234" extensions:"x-order=1 x-nullable=false"`     // 모델패키지 ID
	Name            string `json:"name" validate:"required" example:"This Is a Test Deploy" extensions:"x-order=2 x-nullable=false"` // 배포 명
	Description     string `json:"description" example:"deploy best model" extensions:"x-order=3"`                                   // 배포 설명
	PredictionEnvID string `json:"predictionEnvID" validate:"required" example:"k8s-inference-staging" extensions:"x-order=4"`       //예측 환경
	Importance      string `json:"importance" example:"" extensions:"x-order=5"`
	RequestCPU      string `json:"requestCPU" example:"" extensions:"x-order=7"`
	RequestMEM      string `json:"requestMEM" example:"" extensions:"x-order=8"`
	RequestGPU      string `json:"requestGPU" example:"" extensions:"x-order=9"`
}

type DeploymentCreateResponseDTO struct {
	DeploymentID string
}

type DeploymentGetRequestDTO struct {
	DeploymentID string
}

type DeploymentGetResponseDTO struct {
	ID              string
	ProjectID       string
	ModelPackageID  string
	PredictionEnvID string
	Name            string
	Description     string
	Importance      string
	RequestCPU      string
	RequestMEM      string
	RequestGPU      string
	ActiveStatus    string
	ServiceStatus   string
	ChangeRequested bool
}

type DeploymentGetByNametRequestDTO struct {
	Name string
}

type DeploymentGetByNameResponseDTO struct {
	Deployments []*Deployment
}

type Deployment struct {
	ID              string
	ProjectID       string
	ModelPackageID  string
	PredictionEnvID string
	Name            string
	Description     string
	Importance      string
	RequestCPU      string
	RequestMEM      string
	RequestGPU      string
	ActiveStatus    string
	ServiceStatus   string
	ChangeRequested bool
}

type DeploymentDeleteRequestDTO struct {
	DeploymentID string
}

type DeploymentDeleteResponseDTO struct {
	Message string
}

type DeploymentActiveRequestDTO struct {
	DeploymentID string
}

type DeploymentActiveResponseDTO struct {
	Message string
}
