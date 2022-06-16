package schema

type PostDeploysRequest struct {
	ProjectID                  string `json:"projectID" validate:"required" example:"ID-1234" extensions:"x-order=0 x-nullable=false"`          // 프로젝트 ID
	ModelPackageID             string `json:"modelPackageID" validate:"required" example:"ID-1234" extensions:"x-order=0 x-nullable=false"`     // 모델패키지 ID
	Name                       string `json:"name" validate:"required" example:"This Is a Test Deploy" extensions:"x-order=1 x-nullable=false"` // 배포 명
	Description                string `json:"description" example:"deploy best model" extensions:"x-order=2"`                                   // 배포 설명
	PredictionEnvID            string `json:"predictionEnvID" example:"k8s-inference-staging" extensions:"x-order=2"`                           //예측 환경
	AssociationID              string `json:"associationID" example:"real_target" extensions:"x-order=2"`                                       //예측 요청 데이터 중 업로드 할 실측값 데이터와 맵핑할 수 있는 연결ID
	AssociationIDInRequests    bool   `json:"associationIDInRequests" example:"false" extensions:"x-order=2"`                                   //연결ID가 예측요청 피쳐에 속해 있을 경우 true, 피쳐 외부 json데이터에 있을 경우 false
	EnableAccurancyMonitoring  bool   `json:"enableAccurancyMonitoring" example:"true" extensions:"x-order=2"`                                  //활성화할 시 실측 데이터가 업로드 될 때마다 연결ID로 맵핑해 예측값과 실측값을 비교하고 정확도를 측정한다(실측데이터 업로드 필수)
	EnableFeatureDriftTracking bool   `json:"enableFeatureDriftTracking" example:"true" extensions:"x-order=2"`                                 //활성화할 시 예측요청을 분석하여 피쳐의 분포도 변화를 추적한다 (훈련데이터 업로드 필수)
}

type PostDeploysResponse struct {
	Message string `json:"message"` // 상세 메세지
}

type GetRuntimeFrameWork struct {
	Type string `json:"type" example:"['Tensorflow','PyTorch','Scikit-learn','XGBoost','PMML','Spark''Lightgbm','PaddleServer','Triton']"` //런타임 프레임워크 ['Tensorflow','PyTorch','Scikit-learn','XGBoost','PMML','Spark''Lightgbm','PaddleServer','Triton']
}

type PatchDeploysRequest struct {
	Name        string `json:"name" example:"Model Package Example" extensions:"x-order=0"`              // 모델 패키지 명
	Description string `json:"description" example:"Description of this package" extensions:"x-order=1"` // 모델 패키지 설명
}

type GetAllDeployResponse struct {
	Count      int    `json:"count" extensions:"x-order=0"`      //Optional. Number of items returned on this page.
	Next       string `json:"next" extensions:"x-order=1"`       //Required. URL pointing to the next page (if null, there is no next page). nullable: True format: uri
	Previous   string `json:"previous" extensions:"x-order=2"`   //Required. URL pointing to the previous page (if null, there is no previous page). nullable: True format: uri
	TotalCount int    `json:"totalCount" extensions:"x-order=3"` // Required. The total number of items across all pages.

}

type GetDeployResponse struct {
	ID              string `json:"ID" validate:"required" example:"ID-1234" extensions:"x-order=0 x-nullable=false"`                 // 배포 ID
	Name            string `json:"name" validate:"required" example:"This Is a Test Deploy" extensions:"x-order=1 x-nullable=false"` // 배포 명
	Description     string `json:"description" example:"deploy best model" extensions:"x-order=2"`                                   // 배포 설명
	Project         GetProject
	ModelPackageID  string          `json:"modelPackageID" validate:"required" example:"ID-1234" extensions:"x-order=2 x-nullable=false"` // 모델패키지 ID
	Status          string          //enum: [‘active’, ‘inactive’, ‘archived’]
	State           string          //enum: [‘ready’, ‘stopped’, ‘replacingModel’, ‘errored’]
	CreatedAt       string          //format: date-time
	ApprovalStatus  string          //enum: [‘PENDING’, ‘APPROVED’]
	Model           GetModel        //모델 정보
	ModelPackage    GetModelPackage //모델 정보
	ServiceHealth   ServiceHealth
	AccurancyHealth AccurancyHealth

	PredictionEnvID        string `json:"predictionEnvID" example:"k8s-inference-staging" extensions:"x-order=2"` //예측 환경
	AssociationID          string `json:"associationID" example:"real_target" extensions:"x-order=2"`             //예측 요청 데이터 중 업로드 할 실측값 데이터와 맵핑할 수 있는 연결ID
	AssociationIDInFeature bool   `json:"associationIDInFeature" example:"false" extensions:"x-order=2"`          //연결ID가 예측요청 피쳐에 속해 있을 경우 true, 피쳐 외부 json데이터에 있을 경우 false
	AccurancyMonitoring    bool   `json:"accurancyMonitoring" example:"true" extensions:"x-order=2"`              //활성화할 시 실측 데이터가 업로드 될 때마다 연결ID로 맵핑해 예측값과 실측값을 비교하고 정확도를 측정한다(실측데이터 업로드 필수)
	FeatureDriftTracking   bool   `json:"featureDriftTracking" example:"true" extensions:"x-order=2"`             //활성화할 시 예측요청을 분석하여 피쳐의 분포도 변화를 추적한다 (훈련데이터 업로드 필수)
}

type GetProject struct {
	ProjectID   string `json:"projectID" validate:"required" example:"ID-1234" extensions:"x-order=1 x-nullable=false"` // 프로젝트 ID
	ProjectName string //프로젝트 명
}

type GetModel struct {
	ID          string `json:"id" example:"Model Example"`   //모델ID
	Name        string `json:"name" example:"Model Example"` //모델명
	Version     string `json:"version" example:"V1"`         //버전
	Description string `json:"description"`                  //설명
	URL         string `json:"URL"`                          //스토리지 URL
}

type GetModelPackage struct {
	ID   string `json:"id" example:"Model Example"`   //모델패키지 ID
	Name string `json:"name" example:"Model Example"` //모델패키지 명
}

type ServiceHealth struct {
	Status    string `json:"status" example:"passing"`                //Required. Service health status. enum: [‘passing’, ‘warning’, ‘unknown’, ‘unavailable’, ‘failing’]
	StartDate string `json:"startDate" example:"2022-01-02 03:04:06"` //Required. Start date of service health period. nullable: True format: date-time
	EndDate   string `json:"endDate" example:"2022-01-02 03:04:06"`   //Required. End date of service health period. nullable: True format: date-time
}

type AccurancyHealth struct {
	Status    string `json:"status" example:"passing"`                //Required. Service health status. enum: [‘passing’, ‘warning’, ‘unknown’, ‘unavailable’, ‘failing’]
	Message   string `json:"message" example:""`                      //Required. A message providing more detail on the status.
	StartDate string `json:"startDate" example:"2022-01-02 03:04:06"` //Required. Start date of service health period. nullable: True format: date-time
	EndDate   string `json:"endDate" example:"2022-01-02 03:04:06"`   //Required. End date of service health period. nullable: True format: date-time
}
