package dto

type ClusterInfo struct {
	ClusterInfoID    string
	Name             string // 클러스터 명
	InferenceSvcInfo InferenceSvcInfo
}

type InferenceSvcInfo struct {
	InfereceSvcAPISvrEndPoint   string `json:"infereceSvcAPISvrEndPoint" example:"http://192.168.88.161:30070" extensions:"x-order=2"`   // Kserve API Server 접속주소
	InfereceSvcHostName         string `json:"infereceSvcHostName" example:"example.com" extensions:"x-order=3"`                         // Kserve/Knative에 설정된 호스트명(설치시 기본 example.com)
	InferenceSvcIngressEndPoint string `json:"inferenceSvcIngressEndPoint" example:"http://192.168.88.161:31000" extensions:"x-order=4"` // Kserve Istio Ingress Host IP/Port 설정(Prediction send 요청시 사용)
}

type CreateClusterInfoRequestDTO struct {
	Name                        string `json:"name" example:"ClusterInfo 01" extensions:"x-order=1"`                                     // 클러스터 명
	InfereceSvcAPISvrEndPoint   string `json:"infereceSvcAPISvrEndPoint" example:"http://192.168.88.161:30070" extensions:"x-order=2"`   // Kserve API Server 접속주소
	InfereceSvcHostName         string `json:"infereceSvcHostName" example:"example.com" extensions:"x-order=3"`                         // Kserve/Knative에 설정된 호스트명(설치시 기본 example.com)
	InferenceSvcIngressEndPoint string `json:"inferenceSvcIngressEndPoint" example:"http://192.168.88.161:31000" extensions:"x-order=4"` // Kserve Istio Ingress Host IP/Port 설정(Prediction send 요청시 사용)
}

type CreateClusterInfoResponseDTO struct {
	ClusterInfoID string
}

type GetClusterInfoRequestDTO struct {
	ClusterInfoID string `json:"ClusterInfoID" validate:"false" swaggerignore:"true"` // 프로젝트 ID
}

type GetClusterInfoResponseDTO struct {
	ClusterInfo ClusterInfo
}

type InternalGetClusterInfoRequestDTO struct {
	ClusterInfoID string `json:"ClusterInfoID" validate:"false" swaggerignore:"true"` // 프로젝트 ID
}

type InternalGetClusterInfoResponseDTO struct {
	ClusterInfo ClusterInfo
}

type GetClusterInfoListRequestDTO struct {
	Name  string `json:"name" extensions:"x-order=1"`                              // 검색조건: 프로젝트 명
	Limit int    `json:"limit" extensions:"x-order=2"`                             // 한번에 조회 할 건수
	Page  int    `json:"page" extensions:"x-order=3"`                              // 조회 할 페이지, 첫 조회후 TotalPages 범위 내에서 선택 후 보낸다
	Sort  string `json:"sort" enums:"CreateAsc,CreateDesc" extensions:"x-order=4"` //정열방식, CreateAsc: 생성시간 내림차순, CraeteDesc: 생성시간 역차순
}

type GetClusterInfoListResponseDTO struct {
	Limit      int
	Page       int
	Sort       string
	TotalRows  int64
	TotalPages int
	Rows       interface{}
}

type DeleteClusterInfoRequestDTO struct {
	ClusterInfoID string
}

type DeleteClusterInfoResponseDTO struct {
	Message string
}

type UpdateClusterInfoRequestDTO struct {
	ClusterInfoID               string `json:"ClusterInfoID" validate:"false" swaggerignore:"true"` // 프로젝트 ID
	Name                        string // 클러스터 명
	InfereceSvcAPISvrEndPoint   string `json:"infereceSvcAPISvrEndPoint" example:"http://192.168.88.161:30070" extensions:"x-order=2"`   // Kserve API Server 접속주소
	InfereceSvcHostName         string `json:"infereceSvcHostName" example:"example.com" extensions:"x-order=3"`                         // Kserve/Knative에 설정된 호스트명(설치시 기본 example.com)
	InferenceSvcIngressEndPoint string `json:"inferenceSvcIngressEndPoint" example:"http://192.168.88.161:31000" extensions:"x-order=4"` // Kserve Istio Ingress Host IP/Port 설정(Prediction send 요청시 사용)
}

type UpdateClusterInfoResponseDTO struct {
	Message string
}
