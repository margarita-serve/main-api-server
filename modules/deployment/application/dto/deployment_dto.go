package dto

type Deployment struct {
	ID                     string  `json:"deploymentID"  example:"cbjmmqfr2g4j4bjpq19g" extensions:"x-order=01"`         // deployment ID
	ProjectID              string  `json:"projectID"  example:"cbjmmqfr2g4j4bjpq19g" extensions:"x-order=02"`            // 프로젝트 ID
	ProjectName            string  `json:"projectName"  example:"House Price Project" extensions:"x-order=03"`           // 프로젝트 명
	ModelPackageID         string  `json:"modelPackageID" example:"cbjmmrvr2g4j4bjpq1a0" extensions:"x-order=04"`        // 모델패키지 ID
	ModelPackageName       string  `json:"modelPackageName" example:"House Price Best Acuuracy" extensions:"x-order=05"` // 모델패키지 ID
	PredictionEnvID        string  `json:"predictionEnvID" example:"cbjmmqfr2g4j4bjpq19g" extensions:"x-order=06"`       // 예측 환경   미 입력시 프로젝트에 설정된 기본 배포환경에 배포
	PredictionEnvName      string  `json:"predictionEnvName" example:"Production Inference Cluster" extensions:"x-order=07"`
	URI                    string  `json:"URI" example:"http://example.com/predict" extensions:"x-order=08"`                       // End Point 예측 요청 URL
	Name                   string  `json:"name" validate:"required" example:"This Is a Test Deploy" extensions:"x-order=09"`       // 배포 명
	Description            string  `json:"description" example:"deploy best model" extensions:"x-order=10"`                        // 배포 설명
	Importance             string  `json:"importance" example:"Low" extensions:"x-order=11" enums:"Low, Moderate, High, Critical"` // 배포중요도   미 입력시 'Moderate'로 설정
	RequestCPU             float32 `json:"requestCPU" example:"0.5" extensions:"x-order=12"`                                       // 요청 CPU 단위)1 = 1vCore = 1000millicpu, 범위)0.001 ~, 미 입력시 1
	RequestMEM             float32 `json:"requestMEM" example:"1" extensions:"x-order=13"`                                         // 요청 MEM 단위)1 = 1G= 1Gi  범위)0.001 ~, 미 입력시 2
	ActiveStatus           string  `json:"activeStatus" example:"active" extensions:"x-order=14"`                                  // 배포서비스 활성화 상태
	ServiceStatus          string  `json:"serviceStatus" example:"Ready" extensions:"x-order=15"`                                  // 배포서비스 내부 처리 상태
	ChangeRequested        bool    // 미사용 컬럼
	ServiceHealthStatus    string  `json:"serviceHealthStatus" example:"pass" extensions:"x-order=16"`   //서비스 상태, 24시간 기준 요청이 없을경우 = unknown, 4xx >=1 인경우 = warning, 5xx >=1 인경우 = failing, 4xx or 5xx 없을경우 = pass
	DriftStatus            string  `json:"driftStatus" example:"pass" extensions:"x-order=17"`           //데이터 드리프트 상태, 사용자가 지정한 드리프트 모니터 셋팅에 따라 결정 /30초간격, unknown, warning, failing,  정상 = pass
	AccuracyStatus         string  `json:"accuracyStatus" example:"pass" extensions:"x-order=18"`        //모델 정확도 상태, 사용자가 지정한 정확도 모니터 셋팅에 따라 결정 /30초간격, unknown, warning, failing,  정상 = pass
	FeatureDriftTracking   bool    `json:"featureDriftTracking" example:"false" extensions:"x-order=19"` //데이터 드리프트 설정
	AccuracyAnalyze        bool    `json:"accuracyAnalyze" example:"false" extensions:"x-order=20"`      // 정확도 측정 설정
	AssociationID          string  `json:"associationID" example:"Index" extensions:"x-order=21"`        // 요청데이터에서 ID로 처리할 유일한 피쳐컬럼 명
	AssociationIDInFeature bool    `json:"associationIDInFeature" example:"false" extensions:"x-order=22"`
}

type DeploymentList struct {
	ID                  string `json:"deploymentID"  example:"cbjmmqfr2g4j4bjpq19g" extensions:"x-order=1"`         // deployment ID
	ProjectID           string `json:"projectID"  example:"cbjmmqfr2g4j4bjpq19g" extensions:"x-order=2"`            // 프로젝트 ID
	ProjectName         string `json:"projectName"  example:"House Price Project" extensions:"x-order=3"`           // 프로젝트 명
	ModelPackageID      string `json:"modelPackageID" example:"cbjmmrvr2g4j4bjpq1a0" extensions:"x-order=4"`        // 모델패키지 ID
	ModelPackageName    string `json:"modelPackageName" example:"House Price Best Acuuracy" extensions:"x-order=4"` // 모델패키지 ID
	PredictionEnvID     string `json:"predictionEnvID" example:"cbjmmqfr2g4j4bjpq19g" extensions:"x-order=5"`       // 예측 환경   미 입력시 프로젝트에 설정된 기본 배포환경에 배포
	PredictionEnvName   string `json:"predictionEnvName" example:"Production Inference Cluster" extensions:"x-order=6"`
	Name                string `json:"name" validate:"required" example:"This Is a Test Deploy" extensions:"x-order=8"`        // 배포 명
	Description         string `json:"description" example:"deploy best model" extensions:"x-order=9"`                         // 배포 설명
	Importance          string `json:"importance" example:"Low" extensions:"x-order=10" enums:"Low, Moderate, High, Critical"` // 배포중요도   미 입력시 'Moderate'로 설정
	ActiveStatus        string `json:"activeStatus" example:"active" extensions:"x-order=13"`                                  // 배포서비스 활성화 상태
	ServiceStatus       string `json:"serviceStatus" example:"Ready" extensions:"x-order=14"`
	ServiceHealthStatus string `json:"serviceHealthStatus" example:"pass"`                    //서비스 상태, 24시간 기준 요청이 없을경우 = unknown, 4xx >=1 인경우 = warning, 5xx >=1 인경우 = failing, 4xx or 5xx 없을경우 = pass
	DriftStatus         string `json:"driftStatus" example:"pass" extensions:"x-order=17"`    //데이터 드리프트 상태, 사용자가 지정한 드리프트 모니터 셋팅에 따라 결정 /30초간격, unknown, warning, failing,  정상 = pass
	AccuracyStatus      string `json:"accuracyStatus" example:"pass" extensions:"x-order=18"` //모델 정확도 상태, 사용자가 지정한 정확도 모니터 셋팅에 따라 결정 /30초간격, unknown, warning, failing,  정상 = pass
	ModelFrameWork      string `json:"modelFrameWork" example:"pass" extensions:"x-order=18"` //모델 런타임 프레임워크
}

type CreateDeploymentRequestDTO struct {
	ModelPackageID         string  `json:"modelPackageID" validate:"required" example:"cbjmmrvr2g4j4bjpq1a0" extensions:"x-order=01 x-nullable=false"` // 모델패키지 ID
	Name                   string  `json:"name" validate:"required" example:"This Is a Test Deploy" extensions:"x-order=02 x-nullable=false"`          // 배포 명
	Description            string  `json:"description" example:"deploy best model" extensions:"x-order=03"`                                            // 배포 설명
	PredictionEnvID        string  `json:"predictionEnvID" example:"cbjmmqfr2g4j4bjpq19g" extensions:"x-order=04" swaggerignore:"true"`                // 예측 환경   미 입력시 프로젝트에 설정된 기본 배포환경에 배포
	Importance             string  `json:"importance" example:"Low" extensions:"x-order=05" enums:"Low, Moderate, High, Critical"`                     // 배포중요도   미 입력시 'Moderate'로 설정
	RequestCPU             float32 `json:"requestCPU" example:"0.5" extensions:"x-order=06"`                                                           // 요청 CPU 단위)1 = 1vCore = 1000millicpu, 범위)0.001 ~, 미 입력시 1
	RequestMEM             float32 `json:"requestMEM" example:"1" extensions:"x-order=07"`                                                             // 요청 MEM 단위)1 = 1G= 1Gi  범위)0.001 ~, 미 입력시 2
	FeatureDriftTracking   bool    `json:"featureDriftTracking" example:"false" extensions:"x-order=08"`                                               // 데이터 드리프트 모니터링 설정, 미 입력시 false
	AccuracyAnalyze        bool    `json:"accuracyAnalyze" example:"false" extensions:"x-order=09"`                                                    // 정확도 모니터링 설정, 미 입력시 false
	AssociationID          string  `json:"associationID" example:"" extensions:"x-order=10"`                                                           // 정확도 측정을 위해 유니크한 요청 ID로 처리할 key명, 요청데이터에 별도 json key항목입력 필요
	AssociationIDInFeature bool    `json:"associationIDInFeature" example:"false" extensions:"x-order=11"`                                             // 요청 피쳐컬럼에 유니크한 ID값이 포함되어 있는경우 설정
}

type CreateDeploymentResponseDTO struct {
	DeploymentID string
}

type ReplaceModelRequestDTO struct {
	//ProjectID      string `json:"projectID" validate:"false" swaggerignore:"true"`                                                           // 프로젝트 ID
	DeploymentID   string `json:"deploymentID" validate:"required" swaggerignore:"true"`                                                    // 베포 ID
	ModelPackageID string `json:"modelPackageID" validate:"required"`                                                                       // 교체 할 모델패키지 ID
	Reason         string `json:"reason" validate:"required" enums:"Accuracy, DataDrift, Errors, ScheduledRefresh, PredictionSpeed, Other"` // 교체 이유
	//ManualApplication bool
}

type ReplaceModelResponseDTO struct {
	Message string
}

type UpdateDeploymentRequestDTO struct {
	//ProjectID            string  `json:"projectID" validate:"false" swaggerignore:"true"`                                       // 프로젝트 ID
	DeploymentID           string   `json:"deploymentID" validate:"required" swaggerignore:"true"`                                 // 베포 ID
	Name                   *string  `json:"name" `                                                                                 // 베포 명
	Description            *string  `json:"description" `                                                                          // 베포 설명
	Importance             *string  `json:"importance" example:"Low" extensions:"x-order=5" enums:"Low, Moderate, High, Critical"` // 배포중요도
	RequestCPU             *float32 `json:"requestCPU" example:"1" extensions:"x-order=6"`                                         // 요청 CPU
	RequestMEM             *float32 `json:"requestMEM" example:"2" extensions:"x-order=7"`                                         // 요청 MEM
	FeatureDriftTracking   *bool    `json:"featureDriftTracking" example:"true" extensions:"x-order=8"`                            // 데이터 드리프트 설정
	AccuracyAnalyze        *bool    `json:"accuracyAnalyze" example:"true" extensions:"x-order=9"`                                 // 정확도 측정 설정
	AssociationID          *string  `json:"associationID" example:"Index" extensions:"x-order=10"`                                 // 정확도 측정을 위해 유니크한 요청 ID로 처리할 key명, 요청데이터에 별도 json key항목입력 필요
	AssociationIDInFeature *bool    `json:"associationIDInFeature" example:"false" extensions:"x-order=11"`                        // 요청 피쳐컬럼에 유니크한 ID값이 포함되어 있는경우 설정
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
