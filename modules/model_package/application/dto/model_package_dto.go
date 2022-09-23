package dto

type CreateModelPackageRequestDTO struct {
	ProjectID             string  `json:"projectID" example:"cbjmlbnr2g4j4bjpq18g" validate:"required" extensions:"x-order=0"`                                                  // 프로젝트 ID
	Name                  string  `json:"name" example:"California house price 01" validate:"required" extensions:"x-order=01"`                                                 // 모델패키지 명
	Description           string  `json:"description" example:"New house price predict" extensions:"x-order=02"`                                                                // 모델패키지 설명
	TargetType            string  `json:"targetType" example:"Regression" validate:"required" extensions:"x-order=03" enums:"Binary, Regression"`                               // 모델 예측타입 Binary(Binary classification), Regression
	ModelName             string  `json:"modelName" example:"house price Regression Example Model" extensions:"x-order=04"`                                                     // 모델명
	ModelVersion          string  `json:"modelVersion" example:"1.0" extensions:"x-order=05"`                                                                                   // 모델 버전
	ModelDescription      string  `json:"modelDescription" example:"Best Accuracy" extensions:"x-order=06"`                                                                     // 모델 설명
	PredictionTargetName  string  `json:"predictionTargetName" example:"median_house_value" extensions:"x-order=07"`                                                            // 데이터 셋 예측타켓 컬럼 명칭
	ModelFrameWork        string  `json:"modelFrameWork" example:"SkLearn" validate:"required" extensions:"x-order=08" enums:"TensorFlow, PyTorch, SkLearn, XGBoost, LightGBM"` // 모델 런타임 프레임워크
	ModelFrameWorkVersion string  `json:"modelFrameWorkVersion" example:"1.1.0" extensions:"x-order=09"`                                                                        // 모델 런타임 프레임워크 버전
	PredictionThreshold   float32 `json:"predictionThreshold" example:"0.5" extensions:"x-order=10"`                                                                            // 이진분류 예측임계값. 모델 예측타입이 'Binary' 일 때 필수
	PositiveClassLabel    string  `json:"positiveClassLabel" example:"True" extensions:"x-order=11"`                                                                            // 이진분류 양성명칭 지정. 모델 예측타입이 'Binary' 일 때 필수
	NegativeClassLabel    string  `json:"negativeClassLabel" example:"False" extensions:"x-order=12"`                                                                           // 이진분류 음성명칭 지정. 모델 예측타입이 'Binary' 일 때 필수
}

type CreateModelPackageResponseDTO struct {
	ModelPackageID string
}

type GetModelPackageRequestDTO struct {
	//	ProjectID      string `json:"projectID" validate:"false" swaggerignore:"true"`         // 프로젝트 ID
	ModelPackageID string `json:"modelPackageID" validate:"required" swaggerignore:"true"` // 모델패키지 ID
}

type GetModelPackageResponseDTO struct {
	ModelPackageID        string  `json:"modelPackageID" extensions:"x-order=01"`
	ProjectID             string  `json:"projectID" example:"cbjmlbnr2g4j4bjpq18g" validate:"required" extensions:"x-order=02"`                                                 // 프로젝트 ID
	Name                  string  `json:"name" example:"California house price 01" validate:"required" extensions:"x-order=03"`                                                 // 모델패키지 명
	Description           string  `json:"description" example:"New house price predict" extensions:"x-order=04"`                                                                // 모델패키지 설명
	TargetType            string  `json:"targetType" example:"Regression" validate:"required" extensions:"x-order=05" enums:"Binary, Regression"`                               // 모델 예측타입 Binary(Binary classification), Regression
	ModelName             string  `json:"modelName" example:"house price Regression Example Model" extensions:"x-order=06"`                                                     // 모델명
	ModelVersion          string  `json:"modelVersion" example:"1.0" extensions:"x-order=07"`                                                                                   // 모델 버전
	ModelDescription      string  `json:"modelDescription" example:"Best Accuracy" extensions:"x-order=08"`                                                                     // 모델 설명
	PredictionTargetName  string  `json:"predictionTargetName" example:"median_house_value" extensions:"x-order=09"`                                                            // 데이터 셋 예측타켓 컬럼 명칭
	ModelFrameWork        string  `json:"modelFrameWork" example:"SkLearn" validate:"required" extensions:"x-order=10" enums:"TensorFlow, PyTorch, SkLearn, XGBoost, LightGBM"` // 모델 런타임 프레임워크
	ModelFrameWorkVersion string  `json:"modelFrameWorkVersion" example:"1.1.0" extensions:"x-order=11"`                                                                        // 모델 런타임 프레임워크 버전
	PredictionThreshold   float32 `json:"predictionThreshold" example:"0.5" extensions:"x-order=12"`                                                                            // 이진분류 예측임계값. 모델 예측타입이 'Binary' 일 때 필수
	PositiveClassLabel    string  `json:"positiveClassLabel" example:"True" extensions:"x-order=13"`                                                                            // 이진분류 양성명칭 지정. 모델 예측타입이 'Binary' 일 때 필수
	NegativeClassLabel    string  `json:"negativeClassLabel" example:"False" extensions:"x-order=14"`                                                                           // 이진분류 음성명칭 지정. 모델 예측타입이 'Binary' 일 때 필수
	ModelFileName         string  `json:"modelFileName" example:"model.pkl" extensions:"x-order=15"`                                                                            // 모델파일 명
	TrainingDatasetName   string  `json:"trainingDatasetName" example:"training.csv" extensions:"x-order=16"`                                                                   //훈련 데이터셋 명
	HoldoutDatasetName    string  `json:"holdoutDatasetName" example:"test.csv" extensions:"x-order=17"`                                                                        // 홀드아웃 데이터셋 명
	Archived              bool    `json:"archived" example:"false" extensions:"x-order=18"`                                                                                     // 아카이브 여부
}

type GetModelPackageListRequestDTO struct {
	//	ProjectID string `json:"projectID" validate:"false" swaggerignore:"true"` // 프로젝트 ID
	Name  string `json:"name" extensions:"x-order=1"`                  // 검색조건: 배포 명
	Limit int    `json:"limit" extensions:"x-order=2"`                 // 한번에 조회 할 건수
	Page  int    `json:"page" extensions:"x-order=3"`                  // 조회 할 페이지, 첫 조회후 TotalPages 범위 내에서 선택 후 보낸다
	Sort  string `enums:"CreateAsc,CreateDesc" extensions:"x-order=4"` //정열방식, CreateAsc: 생성시간 내림차순, CraeteDesc: 생성시간 역차순
}

type GetModelPackageListResponseDTO struct {
	Limit      int
	Page       int
	Sort       string
	TotalRows  int64
	TotalPages int
	Rows       interface{}
}

type GetModelPackageListByProjectRequestDTO struct {
	ProjectID string `json:"projectID" validate:"true"` // 프로젝트 ID
}

type GetModelPackageListByProjectResponseDTO struct {
	Rows interface{}
}

type DeleteModelPackageRequestDTO struct {
	//	ProjectID      string
	ModelPackageID string
}

type DeleteModelPackageResponseDTO struct {
	Message string
}

type ArchiveModelPackageRequestDTO struct {
	//	ProjectID      string `json:"projectID" validate:"false" swaggerignore:"true"`         // 프로젝트 ID
	ModelPackageID string `json:"modelPackageID" validate:"required" swaggerignore:"true"` // 모델패키지 ID
}

type ArchiveModelPackageResponseDTO struct {
	Message string
}

type UploadModelRequestDTO struct {
	//	ProjectID      string
	ModelPackageID string
	File           interface{}
	FileName       string
}

type UploadModelResponseDTO struct {
	Message string
}

type UploadTrainingDatasetRequestDTO struct {
	//	ProjectID      string `json:"projectID" validate:"false" swaggerignore:"true"`         // 프로젝트 ID
	ModelPackageID string `json:"modelPackageID" validate:"required" swaggerignore:"true"` // 모델패키지 ID
	File           interface{}
	FileName       string
}

type UploadTrainingDatasetResponseDTO struct {
	Message string
}

type UploadHoldoutDatasetRequestDTO struct {
	//	ProjectID      string
	ModelPackageID string
	File           interface{}
	FileName       string
}

type UploadHoldoutDatasetResponseDTO struct {
	Message string
}

type UpdateModelPackageRequestDTO struct {
	//	ProjectID             string  `json:"projectID" validate:"false" swaggerignore:"true"`
	ModelPackageID        string   `json:"modelPackageID" validate:"required" swaggerignore:"true"`                                                            // 프로젝트 ID
	Name                  *string  `json:"name" example:"ModelPacakge 01" extensions:"x-order=1"`                                                              // 모델패키지 명
	Description           *string  `json:"description" example:"New ModelPacakge" extensions:"x-order=2"`                                                      // 모델패키지 설명
	ModelName             *string  `json:"modelName" example:"Example Model" extensions:"x-order=3"`                                                           // 모델명
	ModelVersion          *string  `json:"modelVersion" example:"1.4" extensions:"x-order=4"`                                                                  // 모델 버전
	ModelDescription      *string  `json:"modelDescription" example:"Best Accuracy" extensions:"x-order=5"`                                                    // 모델 설명
	TargetType            *string  `json:"targetType" example:"Binary" extensions:"x-order=6" enums:"Binary, Regression"`                                      // 모델 예측타입
	PredictionTargetName  *string  `json:"predictionTargetName" example:"Target" extensions:"x-order=7"`                                                       // 예측 타켓명칭
	ModelFrameWork        *string  `json:"modelFrameWork" example:"TensorFlow" extensions:"x-order=8" enums:"TensorFlow, PyTorch, SkLearn, XGBoost, LightGBM"` // 모델 프레임워크
	ModelFrameWorkVersion *string  `json:"modelFrameWorkVersion" example:"1.14.0" extensions:"x-order=9"`                                                      // 모델 프레임워크 버전
	PredictionThreshold   *float32 `json:"predictionThreshold" example:"0.5" extensions:"x-order=10"`                                                          // 이진분류 예측임계값
	PositiveClassLabel    *string  `json:"positiveClassLabel" example:"True" extensions:"x-order=11"`                                                          // 이진분류 양성명칭 지정
	NegativeClassLabel    *string  `json:"negativeClassLabel" example:"False" extensions:"x-order=12"`
	// 이진분류 음성명칭 지정
}

type UpdateModelPackageResponseDTO struct {
	Message string
}

type AddDeployCountRequestDTO struct {
	//	ProjectID      string `json:"projectID" validate:"false" swaggerignore:"true"`         // 프로젝트 ID
	ModelPackageID string `json:"modelPackageID" validate:"required" swaggerignore:"true"` // 모델패키지 ID
}
