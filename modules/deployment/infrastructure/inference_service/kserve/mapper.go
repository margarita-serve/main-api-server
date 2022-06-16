package kserve

import (
	"strings"

	conInfSvcKserve "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/inference_service/kserve_cntr/types"
	domSchema "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/service/inference_service/dto"
)

// MapDisplayCurrentDataByCountryReq mapping DisplayCurrentDataByCountryReq
func MapGetReq(req *domSchema.InferenceServiceGetRequest) (*conInfSvcKserve.GetInferenceServiceRequest, error) {

	reqCon := new(conInfSvcKserve.GetInferenceServiceRequest)
	reqCon.InferenceServer = req.ConnectionInfo
	reqCon.Inferencename = req.Inferencename
	reqCon.Namespace = req.Namespace

	return reqCon, nil
}

func MapGetRes(res *conInfSvcKserve.GetInferenceServiceResponse) (*domSchema.InferenceServiceGetResponse, error) {

	resDom := new(domSchema.InferenceServiceGetResponse)
	resDom.Inferencename = res.Inferencename
	resDom.Message = res.Message

	return resDom, nil
}

func MapCreateReq(req *domSchema.InferenceServiceCreateRequest) (*conInfSvcKserve.CreateInferenceServiceRequest, error) {

	reqCon := new(conInfSvcKserve.CreateInferenceServiceRequest)
	reqCon.InferenceServer = req.ConnectionInfo
	reqCon.Inferencename = req.Inferencename
	reqCon.Namespace = req.Namespace
	reqCon.Predictor = &conInfSvcKserve.Predictor{Modelspec: &conInfSvcKserve.Modelspec{
		Modelframwwork: strings.ToLower(req.ModelFrameWork),
		Storageuri:     req.ModelURL,
		RuntimeVersion: req.ModelFrameWorkVersion,
	},
		Logger: "all"}

	return reqCon, nil
}

func MapCreateRes(res *conInfSvcKserve.CreateInferenceServiceResponse) (*domSchema.InferenceServiceCreateResponse, error) {

	resDom := new(domSchema.InferenceServiceCreateResponse)
	resDom.Inferencename = res.Inferencename
	resDom.Message = res.Message

	return resDom, nil
}

func MapDeleteReq(req *domSchema.InferenceServiceDeleteRequest) (*conInfSvcKserve.DeleteInferenceServiceRequest, error) {

	reqCon := new(conInfSvcKserve.DeleteInferenceServiceRequest)
	reqCon.InferenceServer = req.ConnectionInfo
	reqCon.Inferencename = req.Inferencename
	reqCon.Namespace = req.Namespace

	return reqCon, nil
}

func MapDeleteRes(res *conInfSvcKserve.DeleteInferenceServiceResponse) (*domSchema.InferenceServiceDeleteResponse, error) {

	resDom := new(domSchema.InferenceServiceDeleteResponse)
	resDom.Inferencename = res.Inferencename
	resDom.Message = res.Message

	return resDom, nil
}
