package schema

type GetListView struct {
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
