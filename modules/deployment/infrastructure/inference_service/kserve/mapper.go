package kserve

import (
	"fmt"
	"strings"

	conInfSvcKserve "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/inference_service/kserve_cntr/types"
	domSvcDto "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/service/inference_service/dto"
)

// MapDisplayCurrentDataByCountryReq mapping DisplayCurrentDataByCountryReq
func MapGetReq(req *domSvcDto.InferenceServiceGetRequest) (*conInfSvcKserve.GetInferenceServiceRequest, error) {

	reqCon := new(conInfSvcKserve.GetInferenceServiceRequest)
	reqCon.InferenceServer = req.ConnectionInfo
	reqCon.Inferencename = req.DeploymentID
	reqCon.Namespace = req.Namespace

	return reqCon, nil
}

func MapGetRes(res *conInfSvcKserve.GetInferenceServiceResponse) (*domSvcDto.InferenceServiceGetResponse, error) {

	resDom := new(domSvcDto.InferenceServiceGetResponse)
	resDom.DeploymentID = res.Inferencename
	resDom.Message = res.Message

	return resDom, nil
}

func MapCreateReq(req *domSvcDto.InferenceServiceCreateRequest) (*conInfSvcKserve.CreateInferenceServiceRequest, error) {

	reqCon := new(conInfSvcKserve.CreateInferenceServiceRequest)
	reqCon.InferenceServer = req.ConnectionInfo
	reqCon.Inferencename = req.DeploymentID
	reqCon.Namespace = req.Namespace
	reqCon.Predictor = &conInfSvcKserve.Predictor{Modelspec: &conInfSvcKserve.Modelspec{
		Modelframwwork: strings.ToLower(req.ModelFrameWork),
		Storageuri:     req.ModelURL,
		RuntimeVersion: req.ModelFrameWorkVersion,
	},
		Logger:      "all",
		MinReplicas: 1,
		Resource: &conInfSvcKserve.Resource{
			Requests: &conInfSvcKserve.ResourceType{
				Cpu:    fmt.Sprintf("%.1f", req.RequestCPU),
				Memory: fmt.Sprintf("%.1fGi", req.RequestMEM),
			},
			Limits: &conInfSvcKserve.ResourceType{
				Cpu:    fmt.Sprintf("%.1f", req.LimitCPU),
				Memory: fmt.Sprintf("%.1fGi", req.LimitMEM),
			},
		},
	}

	return reqCon, nil
}

func MapCreateRes(res *conInfSvcKserve.CreateInferenceServiceResponse) (*domSvcDto.InferenceServiceCreateResponse, error) {

	resDom := new(domSvcDto.InferenceServiceCreateResponse)
	resDom.DeploymentID = res.Inferencename
	resDom.ModelHistoryID = res.Revision
	resDom.Message = res.Message

	return resDom, nil
}

func MapActiveReq(req *domSvcDto.InferenceServiceActiveRequest) (*conInfSvcKserve.UpdateInferenceServiceRequest, error) {

	reqCon := new(conInfSvcKserve.UpdateInferenceServiceRequest)
	reqCon.InferenceServer = req.ConnectionInfo
	reqCon.Inferencename = req.DeploymentID
	reqCon.Namespace = req.Namespace
	reqCon.Predictor = &conInfSvcKserve.Predictor{Modelspec: &conInfSvcKserve.Modelspec{
		Modelframwwork: strings.ToLower(req.ModelFrameWork),
		Storageuri:     req.ModelURL,
		RuntimeVersion: req.ModelFrameWorkVersion,
	},
		Logger:      "all",
		MinReplicas: 1,
		Resource: &conInfSvcKserve.Resource{
			Requests: &conInfSvcKserve.ResourceType{
				Cpu:    fmt.Sprintf("%.1f", req.RequestCPU),
				Memory: fmt.Sprintf("%.1fGi", req.RequestMEM),
			},
			Limits: &conInfSvcKserve.ResourceType{
				Cpu:    fmt.Sprintf("%.1f", req.LimitCPU),
				Memory: fmt.Sprintf("%.1fGi", req.LimitMEM),
			},
		},
	}

	return reqCon, nil
}

func MapActiveRes(res *conInfSvcKserve.UpdateInferenceServiceResponse) (*domSvcDto.InferenceServiceActiveResponse, error) {

	resDom := new(domSvcDto.InferenceServiceActiveResponse)
	resDom.DeploymentID = res.Inferencename
	resDom.ModelHistoryID = res.Revision

	return resDom, nil
}

func MapInActiveReq(req *domSvcDto.InferenceServiceInActiveRequest) (*conInfSvcKserve.UpdateInferenceServiceRequest, error) {

	reqCon := new(conInfSvcKserve.UpdateInferenceServiceRequest)
	reqCon.InferenceServer = req.ConnectionInfo
	reqCon.Inferencename = req.DeploymentID
	reqCon.Namespace = req.Namespace
	reqCon.Predictor = &conInfSvcKserve.Predictor{Modelspec: &conInfSvcKserve.Modelspec{
		Modelframwwork: strings.ToLower(req.ModelFrameWork),
		Storageuri:     req.ModelURL,
		RuntimeVersion: req.ModelFrameWorkVersion,
	},
		Logger:      "all",
		MinReplicas: 0,
		MaxReplicas: 0,
		Resource: &conInfSvcKserve.Resource{
			Requests: &conInfSvcKserve.ResourceType{
				Cpu:    fmt.Sprintf("%.1f", req.RequestCPU),
				Memory: fmt.Sprintf("%.1fGi", req.RequestMEM),
			},
			Limits: &conInfSvcKserve.ResourceType{
				Cpu:    fmt.Sprintf("%.1f", req.LimitCPU),
				Memory: fmt.Sprintf("%.1fGi", req.LimitMEM),
			},
		},
	}

	return reqCon, nil
}

func MapInActiveRes(res *conInfSvcKserve.UpdateInferenceServiceResponse) (*domSvcDto.InferenceServiceInActiveResponse, error) {

	resDom := new(domSvcDto.InferenceServiceInActiveResponse)
	resDom.DeploymentID = res.Inferencename
	resDom.ModelHistoryID = res.Revision

	return resDom, nil
}

func MapReplaceModelReq(req *domSvcDto.InferenceServiceReplaceModelRequest) (*conInfSvcKserve.UpdateInferenceServiceRequest, error) {

	reqCon := new(conInfSvcKserve.UpdateInferenceServiceRequest)
	reqCon.InferenceServer = req.ConnectionInfo
	reqCon.Inferencename = req.DeploymentID
	reqCon.Namespace = req.Namespace
	reqCon.Predictor = &conInfSvcKserve.Predictor{Modelspec: &conInfSvcKserve.Modelspec{
		Modelframwwork: strings.ToLower(req.ModelFrameWork),
		Storageuri:     req.ModelURL,
		RuntimeVersion: req.ModelFrameWorkVersion,
	},
		Logger:      "all",
		MinReplicas: 1,
		Resource: &conInfSvcKserve.Resource{
			Requests: &conInfSvcKserve.ResourceType{
				Cpu:    fmt.Sprintf("%.1f", req.RequestCPU),
				Memory: fmt.Sprintf("%.1fGi", req.RequestMEM),
			},
			Limits: &conInfSvcKserve.ResourceType{
				Cpu:    fmt.Sprintf("%.1f", req.LimitCPU),
				Memory: fmt.Sprintf("%.1fGi", req.LimitMEM),
			},
		},
	}

	return reqCon, nil
}

func MapReplaceModelRes(res *conInfSvcKserve.UpdateInferenceServiceResponse) (*domSvcDto.InferenceServiceReplaceModelResponse, error) {

	resDom := new(domSvcDto.InferenceServiceReplaceModelResponse)
	resDom.DeploymentID = res.Inferencename
	resDom.ModelHistoryID = res.Revision

	return resDom, nil
}

func MapDeleteReq(req *domSvcDto.InferenceServiceDeleteRequest) (*conInfSvcKserve.DeleteInferenceServiceRequest, error) {

	reqCon := new(conInfSvcKserve.DeleteInferenceServiceRequest)
	reqCon.InferenceServer = req.ConnectionInfo
	reqCon.Inferencename = req.DeploymentID
	reqCon.Namespace = req.Namespace

	return reqCon, nil
}

func MapDeleteRes(res *conInfSvcKserve.DeleteInferenceServiceResponse) (*domSvcDto.InferenceServiceDeleteResponse, error) {

	resDom := new(domSvcDto.InferenceServiceDeleteResponse)
	resDom.DeploymentID = res.Inferencename
	resDom.Message = res.Message

	return resDom, nil
}
