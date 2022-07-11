package dto

import "io"

type CreateModelPackageRequestDTO struct {
	ProjectID             string  `json:"projectID" validate:"false" swaggerignore:"true"`                                                                    // 프로젝트 ID
	Name                  string  `json:"name" example:"ModelPacakge 01" extensions:"x-order=1"`                                                              // 모델패키지 명
	Description           string  `json:"description" example:"New ModelPacakge" extensions:"x-order=2"`                                                      // 모델패키지 설명
	ModelName             string  `json:"modelName" example:"Example Model" extensions:"x-order=3"`                                                           // 모델명
	ModelVersion          string  `json:"modelVersion" example:"1.4" extensions:"x-order=4"`                                                                  // 모델 버전
	ModelDescription      string  `json:"modelDescription" example:"Best Accuracy" extensions:"x-order=5"`                                                    // 모델 설명
	TargetType            string  `json:"targetType" example:"Binary" extensions:"x-order=6" enums:"Binary, Regression"`                                      // 모델 예측타입
	PredictionTargetName  string  `json:"predictionTargetName" example:"Target" extensions:"x-order=7"`                                                       // 예측 타켓명칭
	ModelFrameWork        string  `json:"modelFrameWork" example:"TensorFlow" extensions:"x-order=8" enums:"TensorFlow, PyTorch, SkLearn, XGBoost, LightGBM"` // 모델 프레임워크
	ModelFrameWorkVersion string  `json:"modelFrameWorkVersion" example:"1.14.0" extensions:"x-order=9"`                                                      // 모델 프레임워크 버전
	PredictionThreshold   float32 `json:"predictionThreshold" example:"0.5" extensions:"x-order=10"`                                                          // 이진분류 예측임계값
	PositiveClassLabel    string  `json:"positiveClassLabel" example:"True" extensions:"x-order=11"`                                                          // 이진분류 양성명칭 지정
	NegativeClassLabel    string  `json:"negativeClassLabel" example:"False" extensions:"x-order=12"`                                                         // 이진분류 음성명칭 지정
}

type CreateModelPackageResponseDTO struct {
	ModelPackageID string
}

type GetModelPackageRequestDTO struct {
	ProjectID      string `json:"projectID" validate:"false" swaggerignore:"true"`         // 프로젝트 ID
	ModelPackageID string `json:"modelPackageID" validate:"required" swaggerignore:"true"` // 모델패키지 ID
}

type GetModelPackageResponseDTO struct {
	ModelPackageID        string
	ProjectID             string
	Name                  string
	Description           string
	ModelName             string
	ModelVersion          string
	ModelDescription      string
	TargetType            string
	PredictionTargetName  string
	ModelFrameWork        string
	ModelFrameWorkVersion string
	PredictionThreshold   float32
	PositiveClassLabel    string
	NegativeClassLabel    string
	ModelFileName         string
	TrainingDatasetName   string
	HoldoutDatasetName    string
	Archived              bool
}

type GetModelPackageListRequestDTO struct {
	ProjectID string `json:"projectID" validate:"false" swaggerignore:"true"` // 프로젝트 ID
	Name      string `json:"name" extensions:"x-order=1"`                     // 검색조건: 배포 명
	Limit     int    `json:"limit" extensions:"x-order=2"`                    // 한번에 조회 할 건수
	Page      int    `json:"page" extensions:"x-order=3"`                     // 조회 할 페이지, 첫 조회후 TotalPages 범위 내에서 선택 후 보낸다
	Sort      string `enums:"CreateAsc,CreateDesc" extensions:"x-order=4"`    //정열방식, CreateAsc: 생성시간 내림차순, CraeteDesc: 생성시간 역차순
}

type GetModelPackageListResponseDTO struct {
	Limit      int
	Page       int
	Sort       string
	TotalRows  int64
	TotalPages int
	Rows       interface{}
}

type DeleteModelPackageRequestDTO struct {
	ProjectID      string
	ModelPackageID string
}

type DeleteModelPackageResponseDTO struct {
	Message string
}

type ArchiveModelPackageRequestDTO struct {
	ProjectID      string `json:"projectID" validate:"false" swaggerignore:"true"`         // 프로젝트 ID
	ModelPackageID string `json:"modelPackageID" validate:"required" swaggerignore:"true"` // 모델패키지 ID
}

type ArchiveModelPackageResponseDTO struct {
	Message string
}

type InternalGetModelPackageRequestDTO struct {
	ModelPackageID string
}

type InternalGetModelPackageResponseDTO struct {
	ModelPackageID        string
	ProjectID             string
	Name                  string
	Description           string
	ModelName             string
	ModelVersion          string
	ModelDescription      string
	TargetType            string
	PredictionTargetName  string
	ModelFrameWork        string
	ModelFrameWorkVersion string
	PredictionThreshold   float32
	PositiveClassLabel    string
	NegativeClassLabel    string
	ModelFilePath         string
	TrainingDatasetPath   string
	HoldoutDatasetPath    string
	Archived              bool
}

type UploadModelRequestDTO struct {
	ProjectID      string
	ModelPackageID string
	File           io.Reader
	FileName       string
}

type UploadModelResponseDTO struct {
	Message string
}

type UploadTrainingDatasetRequestDTO struct {
	ProjectID      string `json:"projectID" validate:"false" swaggerignore:"true"`         // 프로젝트 ID
	ModelPackageID string `json:"modelPackageID" validate:"required" swaggerignore:"true"` // 모델패키지 ID
	File           io.Reader
	FileName       string
}

type UploadTrainingDatasetResponseDTO struct {
	Message string
}

type UploadHoldoutDatasetRequestDTO struct {
	ProjectID      string
	ModelPackageID string
	File           io.Reader
	FileName       string
}

type UploadHoldoutDatasetResponseDTO struct {
	Message string
}

type UpdateModelPackageRequestDTO struct {
	ProjectID             string  `json:"projectID" validate:"false" swaggerignore:"true"`
	ModelPackageID        string  `json:"modelPackageID" validate:"required" swaggerignore:"true"`                                                            // 프로젝트 ID
	Name                  string  `json:"name" example:"ModelPacakge 01" extensions:"x-order=1"`                                                              // 모델패키지 명
	Description           string  `json:"description" example:"New ModelPacakge" extensions:"x-order=2"`                                                      // 모델패키지 설명
	ModelName             string  `json:"modelName" example:"Example Model" extensions:"x-order=3"`                                                           // 모델명
	ModelVersion          string  `json:"modelVersion" example:"1.4" extensions:"x-order=4"`                                                                  // 모델 버전
	ModelDescription      string  `json:"modelDescription" example:"Best Accuracy" extensions:"x-order=5"`                                                    // 모델 설명
	TargetType            string  `json:"targetType" example:"Binary" extensions:"x-order=6" enums:"Binary, Regression"`                                      // 모델 예측타입
	PredictionTargetName  string  `json:"predictionTargetName" example:"Target" extensions:"x-order=7"`                                                       // 예측 타켓명칭
	ModelFrameWork        string  `json:"modelFrameWork" example:"TensorFlow" extensions:"x-order=8" enums:"TensorFlow, PyTorch, SkLearn, XGBoost, LightGBM"` // 모델 프레임워크
	ModelFrameWorkVersion string  `json:"modelFrameWorkVersion" example:"1.14.0" extensions:"x-order=9"`                                                      // 모델 프레임워크 버전
	PredictionThreshold   float32 `json:"predictionThreshold" example:"0.5" extensions:"x-order=10"`                                                          // 이진분류 예측임계값
	PositiveClassLabel    string  `json:"positiveClassLabel" example:"True" extensions:"x-order=11"`                                                          // 이진분류 양성명칭 지정
	NegativeClassLabel    string  `json:"negativeClassLabel" example:"False" extensions:"x-order=12"`
	// 이진분류 음성명칭 지정
}

type UpdateModelPackageResponseDTO struct {
	Message string
}
