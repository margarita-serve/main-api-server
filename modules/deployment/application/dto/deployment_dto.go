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
	FeatureDriftTracking bool   `json:"featureDriftTracking" example:"True" extensions:"x-order=8"` //데이터 드리프트 설정
	AccuracyAnalyze      bool   `json:"accuracyAnalyze" example:"True" extensions:"x-order=9"`      // 정확도 측정 설정
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
	ModelPackageID       string  `json:"modelPackageID" validate:"required" example:"cbjmmrvr2g4j4bjpq1a0" extensions:"x-order=1 x-nullable=false"` // 모델패키지 ID
	Name                 string  `json:"name" validate:"required" example:"This Is a Test Deploy" extensions:"x-order=2 x-nullable=false"`          // 배포 명
	Description          string  `json:"description" example:"deploy best model" extensions:"x-order=3"`                                            // 배포 설명
	PredictionEnvID      string  `json:"predictionEnvID" example:"cbjmmqfr2g4j4bjpq19g" extensions:"x-order=4"`                                     // 예측 환경   미 입력시 프로젝트에 설정된 기본 배포환경에 배포
	Importance           string  `json:"importance" example:"Low" extensions:"x-order=5" enums:"Low, Moderate, High, Critical"`                     // 배포중요도   미 입력시 'Moderate'로 설정
	RequestCPU           float32 `json:"requestCPU" example:"0.5" extensions:"x-order=6"`                                                           // 요청 CPU 단위)1 = 1vCore = 1000millicpu, 범위)0.001 ~, 미 입력시 1
	RequestMEM           float32 `json:"requestMEM" example:"1" extensions:"x-order=7"`                                                             // 요청 MEM 단위)1 = 1G= 1Gi  범위)0.001 ~, 미 입력시 2
	FeatureDriftTracking bool    `json:"featureDriftTracking" example:"false" extensions:"x-order=8"`                                               // 데이터 드리프트 모니터링 설정, 미 입력시 false
	AccuracyAnalyze      bool    `json:"accuracyAnalyze" example:"false" extensions:"x-order=9"`                                                    // 정확도 모니터링 설정, 미 입력시 false
	AssociationID        string  `json:"associationID" example:"" extensions:"x-order=10"`                                                          // 정확도 측정을 위해 요청데이터에서 ID로 처리할 피쳐컬럼 명, 요청 피쳐컬럼에 유니크한 ID값이 없다면 비워두고 요청데이터에 별도"association_id" key항목입력 필요
}

type CreateDeploymentResponseDTO struct {
	DeploymentID string
}

type ReplaceModelRequestDTO struct {
	//ProjectID      string `json:"projectID" validate:"false" swaggerignore:"true"`                                                           // 프로젝트 ID
	DeploymentID   string `json:"deploymentID" validate:"required" swaggerignore:"true"`                                                     // 베포 ID
	ModelPackageID string `json:"modelPackageID" validate:"required"`                                                                        // 교체 할 모델패키지 ID
	Reason         string `json:"reason" validate:"required" enums:"Accurancy, DataDrift, Errors, ScheduledRefresh, PredictionSpeed, Other"` // 교체 이유
	//ManualApplication bool
}

type ReplaceModelResponseDTO struct {
	Message string
}

type UpdateDeploymentRequestDTO struct {
	//ProjectID            string  `json:"projectID" validate:"false" swaggerignore:"true"`                                       // 프로젝트 ID
	DeploymentID         string   `json:"deploymentID" validate:"required" swaggerignore:"true"`                                 // 베포 ID
	Name                 *string  `json:"name" `                                                                                 // 베포 명
	Description          *string  `json:"description" `                                                                          // 베포 설명
	Importance           *string  `json:"importance" example:"Low" extensions:"x-order=5" enums:"Low, Moderate, High, Critical"` // 배포중요도
	RequestCPU           *float32 `json:"requestCPU" example:"1" extensions:"x-order=6"`                                         // 요청 CPU
	RequestMEM           *float32 `json:"requestMEM" example:"2" extensions:"x-order=7"`                                         // 요청 MEM
	FeatureDriftTracking *bool    `json:"featureDriftTracking" example:"True" extensions:"x-order=8"`                            // 데이터 드리프트 설정
	AccuracyAnalyze      *bool    `json:"accuracyAnalyze" example:"True" extensions:"x-order=9"`                                 // 정확도 측정 설정
	AssociationID        *string  `json:"associationID" example:"Index" extensions:"x-order=9"`                                  // 요청데이터에서 ID로 처리할 유일한 피쳐컬럼 명
}

type UpdateDeploymentResponseDTO struct {
	Message string
}

type GetDeploymentRequestDTO struct {
	//ProjectID    string
	DeploymentID string
}

type GetDeploymentResponseDTO struct {
	Deployment
}

type GetDeploymentListRequestDTO struct {
	//ProjectID string `json:"projectID" validate:"false" swaggerignore:"true"` // 프로젝트 ID
	Name  string `json:"name" `                 // 검색조건: 배포 명
	Limit int    `json:"limit" `                // 한번에 조회 할 건수
	Page  int    `json:"page" `                 // 조회 할 페이지, 첫 조회후 TotalPages 범위 내에서 선택 후 보낸다
	Sort  string `enums:"CreateAsc,CreateDesc"` //정열방식, CreateAsc: 생성시간 내림차순, CraeteDesc: 생성시간 역차순
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
	//ProjectID string
	Name string
}

type GetDeploymentByNameResponseDTO struct {
	Deployments []*Deployment
}

type DeleteDeploymentRequestDTO struct {
	//ProjectID    string
	DeploymentID string
}

type DeleteDeploymentResponseDTO struct {
	Message string
}

type ActiveDeploymentRequestDTO struct {
	//ProjectID    string
	DeploymentID string
}

type ActiveDeploymentResponseDTO struct {
	Message string
}

type InActiveDeploymentRequestDTO struct {
	//ProjectID    string
	DeploymentID string
}

type InActiveDeploymentResponseDTO struct {
	Message string
}
type SendPredictionRequestDTO struct {
	//ProjectID    string `json:"projectID" validate:"false" swaggerignore:"true"`                                                                                                                                                                                                                                            // 프로젝트 ID
	DeploymentID string `json:"deploymentID" validate:"false" swaggerignore:"true"`                                                                                              // 배포 ID
	JsonData     string `validate:"required" example:"{\"association_id\": [\"abcd1234\"], \"instances\": { [[-122.12,	37.68,	45.0,	2179.0,	401.0,	1159.0,	399.0,	3.4839]] } }"` //  예측요청데이터
}

type SendPredictionResponseDTO struct {
	Message          string
	PredictionResult string
}

type GetModelHistoryRequestDTO struct {
	//ProjectID    string
	DeploymentID string
}

type GetModelHistoryResponseDTO struct {
	ModelHistory interface{}
}

type GetGovernanceHistoryRequestDTO struct {
	//ProjectID    string
	DeploymentID string
}

type GetGovernanceHistoryResponseDTO struct {
	EventHistory interface{}
}

type AddGovernanceHistoryRequestDTO struct {
	DeploymentID string
	EventType    string
	LogMessage   string
}
