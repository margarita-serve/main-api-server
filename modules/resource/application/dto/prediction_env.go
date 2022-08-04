package dto

type CreatePredictionEnvRequestDTO struct {
	Name          string `json:"name" example:"PredictionEnv 01" extensions:"x-order=1"`              // 배포환경 명
	Description   string `json:"description" example:"New PredictionEnv" extensions:"x-order=2"`      // 설명
	ClusterInfoID string `json:"clusterInfoID" example:"12345678901234567890" extensions:"x-order=3"` // 클러스터정보 ID
	UseType       string `json:"useType" enums:"Default, Custom, PackageTest" extensions:"x-order=4"` // 클러스터 등록 타입
	Namespace     string `json:"namespace" example:"default" extensions:"x-order=5"`                  // 배포할 Kubernetes Namespace
}

type CreatePredictionEnvResponseDTO struct {
	PredictionEnvID string
}

type GetPredictionEnvRequestDTO struct {
	PredictionEnvID string `json:"predictionEnvID" validate:"false" swaggerignore:"true"` // 프로젝트 ID
}

type GetPredictionEnvResponseDTO struct {
	PredictionEnvID string
	Name            string      `json:"name"  extensions:"x-order=1"`          // 배포환경 명
	Description     string      `json:"description"  extensions:"x-order=2"`   // 설명
	ClusterInfo     ClusterInfo `json:"clusterInfoID"  extensions:"x-order=3"` // 클러스터정보 ID
	UseType         string      `json:"useType"  extensions:"x-order=4"`       // 클러스터 등록 타입
	Namespace       string      `json:"namespace" extensions:"x-order=5"`      // 배포할 Kubernetes Namespace
	Projects        []interface{}
}

type InternalGetPredictionEnvRequestDTO struct {
	PredictionEnvID string `json:"predictionEnvID" validate:"false" swaggerignore:"true"` // 프로젝트 ID
}

type InternalGetPredictionEnvResponseDTO struct {
	PredictionEnvID string
	Name            string      `json:"name"  extensions:"x-order=1"`         // 배포환경 명
	Description     string      `json:"description" extensions:"x-order=2"`   // 설명
	ClusterInfo     ClusterInfo `json:"clusterInfoID" extensions:"x-order=3"` // 클러스터정보 ID
	UseType         string      `json:"useType"  extensions:"x-order=4"`      // 클러스터 등록 타입
	Namespace       string      `json:"namespace"  extensions:"x-order=5"`    // 배포할 Kubernetes Namespace
}

type GetPredictionEnvListRequestDTO struct {
	Name  string `json:"name" extensions:"x-order=1"`                  // 검색조건: 프로젝트 명
	Limit int    `json:"limit" extensions:"x-order=2"`                 // 한번에 조회 할 건수
	Page  int    `json:"page" extensions:"x-order=3"`                  // 조회 할 페이지, 첫 조회후 TotalPages 범위 내에서 선택 후 보낸다
	Sort  string `enums:"CreateAsc,CreateDesc" extensions:"x-order=4"` //정열방식, CreateAsc: 생성시간 내림차순, CraeteDesc: 생성시간 역차순
}

type GetPredictionEnvListResponseDTO struct {
	Limit      int
	Page       int
	Sort       string
	TotalRows  int64
	TotalPages int
	Rows       interface{}
}

type DeletePredictionEnvRequestDTO struct {
	PredictionEnvID string
}

type DeletePredictionEnvResponseDTO struct {
	Message string
}

type UpdatePredictionEnvRequestDTO struct {
	PredictionEnvID string `json:"predictionEnvID" validate:"false" swaggerignore:"true"`               // 프로젝트 ID
	Name            string `json:"name" example:"PredictionEnv 01" extensions:"x-order=1"`              // 배포환경 명
	Description     string `json:"description" example:"New PredictionEnv" extensions:"x-order=2"`      // 설명
	ClusterInfoID   string `json:"clusterInfoID" example:"12345678901234567890" extensions:"x-order=3"` // 클러스터정보 ID
	UseType         string `json:"useType" enums:"Default, Custom, PackageTest" extensions:"x-order=4"` // 클러스터 등록 타입
	Namespace       string `json:"namespace" example:"default" extensions:"x-order=5"`                  // 배포할 Kubernetes Namespace
}

type UpdatePredictionEnvResponseDTO struct {
	Message string
}

type RegProjectRequestDTO struct {
	PredictionEnvID string `json:"predictionEnvID" validate:"false" swaggerignore:"true"`           // 배포환경 ID
	ProjectID       string `json:"projectID" example:"12345678901234567890" extensions:"x-order=1"` // 프로젝트 ID
}

type RegProjectResponseDTO struct {
	Message string
}

type GetProjectsPredictionEnvsRequestDTO struct {
	Name  string `json:"name" extensions:"x-order=1"`                  // 검색조건: 프로젝트 명
	Limit int    `json:"limit" extensions:"x-order=2"`                 // 한번에 조회 할 건수
	Page  int    `json:"page" extensions:"x-order=3"`                  // 조회 할 페이지, 첫 조회후 TotalPages 범위 내에서 선택 후 보낸다
	Sort  string `enums:"CreateAsc,CreateDesc" extensions:"x-order=4"` //정열방식, CreateAsc: 생성시간 내림차순, CraeteDesc: 생성시간 역차순

	ProjectID string `json:"projectID" example:"12345678901234567890" extensions:"x-order=1"` // 프로젝트 ID
}

type GetProjectsPredictionEnvsResponseDTO struct {
	Limit      int
	Page       int
	Sort       string
	TotalRows  int64
	TotalPages int
	Rows       interface{}
}

type PredictionEnvs struct {
	PredictionEnvID string
	Name            string      `json:"name" example:"PredictionEnv 01" extensions:"x-order=1"`              // 배포환경 명
	Description     string      `json:"description" example:"New PredictionEnv" extensions:"x-order=2"`      // 설명
	ClusterInfo     ClusterInfo `json:"clusterInfoID" example:"12345678901234567890" extensions:"x-order=3"` // 클러스터정보 ID
	UseType         string      `json:"useType" enums:"Default, Custom, PackageTest" extensions:"x-order=4"` // 클러스터 등록 타입
	Namespace       string      `json:"namespace" example:"default" extensions:"x-order=5"`                  // 배포할 Kubernetes Namespace
}
