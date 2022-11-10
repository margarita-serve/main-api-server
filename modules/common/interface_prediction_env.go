package common

type PredictionEnv struct {
	PredictionEnvID string      //
	Name            string      // 배포환경 명
	Description     string      // 설명
	ClusterInfo     ClusterInfo // 클러스터정보
	UseType         string      // 클러스터 등록 타입
	Namespace       string      // 배포할 Kubernetes Namespace
}
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
