package inference_service

import domSvcDto "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/service/inference_service/dto"

// ICovid19Adapter interface
type IInferenceServiceAdapter interface {
	InferenceServiceCreate(req *domSvcDto.InferenceServiceCreateRequest) (*domSvcDto.InferenceServiceCreateResponse, error)
	InferenceServiceReplaceModel(req *domSvcDto.InferenceServiceReplaceModelRequest) (*domSvcDto.InferenceServiceReplaceModelResponse, error)
	InferenceServiceDelete(req *domSvcDto.InferenceServiceDeleteRequest) (*domSvcDto.InferenceServiceDeleteResponse, error)
	InferenceServiceGet(req *domSvcDto.InferenceServiceGetRequest) (*domSvcDto.InferenceServiceGetResponse, error)
	InferenceServiceActive(req *domSvcDto.InferenceServiceActiveRequest) (*domSvcDto.InferenceServiceActiveResponse, error)
	InferenceServiceInActive(req *domSvcDto.InferenceServiceInActiveRequest) (*domSvcDto.InferenceServiceInActiveResponse, error)
	//Update(req *InferenceServiceModelReplaceRequest) (*InferenceServiceModelReplaceResponse, error)
}
