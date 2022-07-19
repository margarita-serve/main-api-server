package dto

type Deployment struct {
	ID                   string
	ProjectID            string
	ModelPackageID       string
	PredictionEnvID      string
	PredictionEnvName    string
	URI                  string
	Name                 string
	Description          string
	Importance           string
	RequestCPU           float32
	RequestMEM           float32
	ActiveStatus         string
	ServiceStatus        string
	ChangeRequested      bool
	ServiceHealthStatus  string
	DriftStatus          string
	AccuracyStatus       string
	FeatureDriftTracking string `json:"featureDriftTracking" example:"True" extensions:"x-order=8"` //데이터 드리프트 설정
	AccuracyAnalyze      string `json:"accuracyAnalyze" example:"True" extensions:"x-order=9"`      // 정확도 측정 설정
	AssociationID        string `json:"associationID" example:"Index" extensions:"x-order=9"`       // 요청데이터에서 ID로 처리할 유일한 피쳐컬럼 명
}

type DeploymentList struct {
	ID                string
	ProjectID         string
	ModelPackageID    string
	PredictionEnvID   string
	PredictionEnvName string
	Name              string
	Importance        string
	ModelFrameWork    string
}

type CreateDeploymentRequestDTO struct {
	ProjectID            string  `json:"projectID" validate:"required" example:"ID-1234" extensions:"x-order=0 x-nullable=false" swaggerignore:"true"` // 프로젝트 ID
	ModelPackageID       string  `json:"modelPackageID" validate:"required" example:"ID-1234" extensions:"x-order=1 x-nullable=false"`                 // 모델패키지 ID
	Name                 string  `json:"name" validate:"required" example:"This Is a Test Deploy" extensions:"x-order=2 x-nullable=false"`             // 배포 명
	Description          string  `json:"description" example:"deploy best model" extensions:"x-order=3"`                                               // 배포 설명
	PredictionEnvID      string  `json:"predictionEnvID" validate:"required" example:"k8s-inference-staging" extensions:"x-order=4"`                   // 예측 환경
	Importance           string  `json:"importance" example:"Low" extensions:"x-order=5" enums:"Low, Moderate, High, Critical"`                        // 배포중요도
	RequestCPU           float32 `json:"requestCPU" example:"1" extensions:"x-order=6"`                                                                //요청 CPU
	RequestMEM           float32 `json:"requestMEM" example:"2" extensions:"x-order=7"`                                                                //요청 MEM
	FeatureDriftTracking string  `json:"featureDriftTracking" example:"True" extensions:"x-order=8"`                                                   //데이터 드리프트 설정
	AccuracyAnalyze      string  `json:"accuracyAnalyze" example:"True" extensions:"x-order=9"`                                                        // 정확도 측정 설정
	AssociationID        string  `json:"associationID" example:"Index" extensions:"x-order=9"`                                                         // 요청데이터에서 ID로 처리할 유일한 피쳐컬럼 명
}

type CreateDeploymentResponseDTO struct {
	DeploymentID string
}

type ReplaceModelRequestDTO struct {
	ProjectID      string `json:"projectID" validate:"false" swaggerignore:"true"`                                                           // 프로젝트 ID
	DeploymentID   string `json:"deploymentID" validate:"required" swaggerignore:"true"`                                                     // 베포 ID
	ModelPackageID string `json:"modelPackageID" validate:"required"`                                                                        // 교체 할 모델패키지 ID
	Reason         string `json:"reason" validate:"required" enums:"Accurancy, DataDrift, Errors, ScheduledRefresh, PredictionSpeed, Other"` // 교체 이유
	//ManualApplication bool
}

type ReplaceModelResponseDTO struct {
	Message string
}

type UpdateDeploymentRequestDTO struct {
	ProjectID            string  `json:"projectID" validate:"false" swaggerignore:"true"`                                       // 프로젝트 ID
	DeploymentID         string  `json:"deploymentID" validate:"required" swaggerignore:"true"`                                 // 베포 ID
	Name                 string  `json:"name" `                                                                                 // 베포 명
	Description          string  `json:"description" `                                                                          // 베포 설명
	Importance           string  `json:"importance" example:"Low" extensions:"x-order=5" enums:"Low, Moderate, High, Critical"` // 배포중요도
	RequestCPU           float32 `json:"requestCPU" example:"1" extensions:"x-order=6"`                                         // 요청 CPU
	RequestMEM           float32 `json:"requestMEM" example:"2" extensions:"x-order=7"`                                         // 요청 MEM
	FeatureDriftTracking string  `json:"featureDriftTracking" example:"True" extensions:"x-order=8"`                            // 데이터 드리프트 설정
	AccuracyAnalyze      string  `json:"accuracyAnalyze" example:"True" extensions:"x-order=9"`                                 // 정확도 측정 설정
	AssociationID        string  `json:"associationID" example:"Index" extensions:"x-order=9"`                                  // 요청데이터에서 ID로 처리할 유일한 피쳐컬럼 명
}

type UpdateDeploymentResponseDTO struct {
	Message string
}

type GetDeploymentRequestDTO struct {
	ProjectID    string
	DeploymentID string
}

type GetDeploymentResponseDTO struct {
	Deployment
}

type GetDeploymentListRequestDTO struct {
	ProjectID string `json:"projectID" validate:"false" swaggerignore:"true"` // 프로젝트 ID
	Name      string `json:"name" `                                           // 검색조건: 배포 명
	Limit     int    `json:"limit" `                                          // 한번에 조회 할 건수
	Page      int    `json:"page" `                                           // 조회 할 페이지, 첫 조회후 TotalPages 범위 내에서 선택 후 보낸다
	Sort      string `enums:"CreateAsc,CreateDesc"`                           //정열방식, CreateAsc: 생성시간 내림차순, CraeteDesc: 생성시간 역차순
}
type GetDeploymentListResponseDTO struct {
	Limit      int
	Page       int
	Sort       string
	TotalRows  int64
	TotalPages int
	Rows       interface{}
}

type GetDeploymentByNametRequestDTO struct {
	ProjectID string
	Name      string
}

type GetDeploymentByNameResponseDTO struct {
	Deployments []*Deployment
}

type DeleteDeploymentRequestDTO struct {
	ProjectID    string
	DeploymentID string
}

type DeleteDeploymentResponseDTO struct {
	Message string
}

type ActiveDeploymentRequestDTO struct {
	ProjectID    string
	DeploymentID string
}

type ActiveDeploymentResponseDTO struct {
	Message string
}

type InActiveDeploymentRequestDTO struct {
	ProjectID    string
	DeploymentID string
}

type InActiveDeploymentResponseDTO struct {
	Message string
}

type GetPredictionURLRequestDTO struct {
	ProjectID    string
	DeploymentID string
}

type GetPredictionURLResponseDTO struct {
	PredictionURL string
}

type SendPredictionRequestDTO struct {
	ProjectID    string `json:"projectID" validate:"false" swaggerignore:"true"`                                                                                                                                                                                                                                            // 프로젝트 ID
	DeploymentID string `json:"deploymentID" validate:"false" swaggerignore:"true"`                                                                                                                                                                                                                                         // 배포 ID
	JsonData     string `validate:"required" example:"{\"association_id\": [\"abcd1234\", \"abcd1235\"], \"instances\": [[1.483887, 1.865988, 2.234620, 1.018782, -2.530891, -1.604642, 0.774676, -0.465148, -0.495225], [1.483887, 1.865988, 2.234620, 1.018782, -2.530891, -1.604642, 0.774676, -0.465148, -0.495225]]}"` // 배포 ID
}

type SendPredictionResponseDTO struct {
	Message          string
	PredictionResult string
}

type GetModelHistoryRequestDTO struct {
	ProjectID    string
	DeploymentID string
}

type GetModelHistoryResponseDTO struct {
	ModelHistory interface{}
}

type GetGovernanceHistoryRequestDTO struct {
	ProjectID    string
	DeploymentID string
}

type GetGovernanceHistoryResponseDTO struct {
	EventHistory interface{}
}
