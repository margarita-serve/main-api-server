package kserve

import (
	"fmt"

	conInfSvcKserve "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/inference_service/kserve_cntr"
	domSvcDto "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/service/inference_service/dto"
)

//
func NewInfSvcKserveAdapter() (*InfSvcKserveAdapter, error) {
	var err error

	adp := new(InfSvcKserveAdapter)
	// adp.handler = h

	//cfg, err := h.GetConfig()
	// if err != nil {
	// 	return nil, nil, err

	// }

	// info := infC19Adp.Covid19AdapterInfo{
	// 	Code:   cfg.Connectors.Covid19.Covid19goid.Code,
	// 	Name:   cfg.Connectors.Covid19.Covid19goid.Name,
	// 	Server: cfg.Connectors.Covid19.Covid19goid.Server,
	// 	Enable: cfg.Connectors.Covid19.Covid19goid.Enable,
	// }
	// adp.SetInfo(info)

	//config := conInfSvcKserve.Config{Server: "http://192.168.88.161:30070"}
	config := conInfSvcKserve.Config{}
	adp.connector = conInfSvcKserve.NewInferenceService(config, nil)

	return adp, err
}

type InfSvcKserveAdapter struct {
	//infInfSvcAdp.BaseInferenceServiceAdapter
	// handler   *handler.Handler
	connector *conInfSvcKserve.InferenceService
}

func (a *InfSvcKserveAdapter) InferenceServiceGet(req *domSvcDto.InferenceServiceGetRequest) (*domSvcDto.InferenceServiceGetResponse, error) {
	resp := new(domSvcDto.InferenceServiceGetResponse)
	connReq, err := MapGetReq(req)
	if err != nil {
		return nil, err
	}

	connResp, err := a.connector.GetInferenceService(connReq)
	if err != nil {
		return nil, err
	}

	resp, err = MapGetRes(connResp)
	if err != nil {
		return nil, err
	}

	// resp.Provider = a.GetInfo().Code
	// resp.Information = fmt.Sprintf("[Enable: %v] %s [%s]", a.GetInfo().Enable, a.GetInfo().Name, a.GetInfo().Server)

	return resp, nil
}

func (a *InfSvcKserveAdapter) InferenceServiceCreate(req *domSvcDto.InferenceServiceCreateRequest) (*domSvcDto.InferenceServiceCreateResponse, error) {
	resp := new(domSvcDto.InferenceServiceCreateResponse)

	connReq, err := MapCreateReq(req)
	if err != nil {
		return nil, err
	}

	connResp, err := a.connector.CreateInferenceService(connReq)
	if err != nil {
		return nil, err
	}

	resp, err = MapCreateRes(connResp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *InfSvcKserveAdapter) InferenceServiceReplaceModel(req *domSvcDto.InferenceServiceReplaceModelRequest) (*domSvcDto.InferenceServiceReplaceModelResponse, error) {
	resp := new(domSvcDto.InferenceServiceReplaceModelResponse)

	connReq, err := MapReplaceModelReq(req)
	if err != nil {
		return nil, err
	}

	connResp, err := a.connector.UpdateInferenceService(connReq)
	if err != nil {
		return nil, err
	}

	resp, err = MapReplaceModelRes(connResp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *InfSvcKserveAdapter) InferenceServiceDelete(req *domSvcDto.InferenceServiceDeleteRequest) (*domSvcDto.InferenceServiceDeleteResponse, error) {
	resp := new(domSvcDto.InferenceServiceDeleteResponse)
	connReq, err := MapDeleteReq(req)
	if err != nil {
		return nil, err
	}

	connResp, err := a.connector.DeleteInferenceService(connReq)
	if err != nil {
		return nil, err
	}

	resp, err = MapDeleteRes(connResp)
	if err != nil {
		fmt.Printf("resp: %v\n", resp)
		return nil, err
	}

	return nil, err
}

func (a *InfSvcKserveAdapter) InferenceServiceActive(req *domSvcDto.InferenceServiceActiveRequest) (*domSvcDto.InferenceServiceActiveResponse, error) {
	resp := new(domSvcDto.InferenceServiceActiveResponse)
	connReq, err := MapActiveReq(req)
	if err != nil {
		return nil, err
	}

	connResp, err := a.connector.UpdateInferenceService(connReq)
	if err != nil {
		return nil, err
	}

	resp, err = MapActiveRes(connResp)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *InfSvcKserveAdapter) InferenceServiceInActive(req *domSvcDto.InferenceServiceInActiveRequest) (*domSvcDto.InferenceServiceInActiveResponse, error) {

	resp := new(domSvcDto.InferenceServiceInActiveResponse)
	connReq, err := MapInActiveReq(req)
	if err != nil {
		return nil, err
	}

	connResp, err := a.connector.UpdateInferenceService(connReq)
	if err != nil {
		return nil, err
	}

	resp, err = MapInActiveRes(connResp)
	if err != nil {
		return nil, err
	}

	return resp, err
}
