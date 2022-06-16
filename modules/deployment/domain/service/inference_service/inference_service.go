package inference_service

import domSvcDto "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/service/inference_service/dto"

// ICovid19Adapter interface
type IInferenceServiceAdapter interface {
	InferenceServiceCreate(req *domSvcDto.InferenceServiceCreateRequest) (*domSvcDto.InferenceServiceCreateResponse, error)
	InferenceServiceDelete(req *domSvcDto.InferenceServiceDeleteRequest) error
	InferenceServiceGet(req *domSvcDto.InferenceServiceGetRequest) (*domSvcDto.InferenceServiceGetResponse, error)
	InferenceServiceActive(id string) error
	InferenceServiceInActive(id string) error
	//Update(req *InferenceServiceModelReplaceRequest) (*InferenceServiceModelReplaceResponse, error)
}
