package kserve_cntr

import (
	"testing"

	infsType "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/inference_service/kserve_cntr/types"
)

func newInferenceService(t *testing.T) *InferenceService {

	config := Config{Server: "http://192.168.88.161:30070"}
	return NewInferenceService(config, nil)
}

func TestGetInfereneceService(t *testing.T) {

	c := newInferenceService(t)
	resp, err := c.GetInferenceService(
		&infsType.GetInferenceServiceRequest{
			Namespace:     "default",
			Inferencename: "mpg-sample-test",
		})
	if err != nil {
		t.Error(err)
	}

	if resp != nil && resp.Message == "Success" {
		t.Logf("RESPONSE.Inferencename: %#v", resp.Inferencename)
	}
}

func TestInfereneceServiceNotFound(t *testing.T) {

	c := newInferenceService(t)
	resp, err := c.GetInferenceService(
		&infsType.GetInferenceServiceRequest{
			Namespace:     "default",
			Inferencename: "mpg",
		})
	if err != nil {
		t.Logf("TestERROR: %s", err)
	}

	if resp != nil && resp.Message != "Success" {
		t.Logf("RESPONSE.Message: %#v", resp.Message)
	}
}

func TestCreateInferenceService(t *testing.T) {
	c := newInferenceService(t)
	resp, err := c.CreateInferenceService(
		&infsType.CreateInferenceServiceRequest{
			Namespace:     "default",
			Inferencename: "mpg-sample-full-13",
			Predictor: &infsType.Predictor{Modelspec: &infsType.Modelspec{
				Modelframwwork: "tensorflow",
				Storageuri:     "s3://testmodel/mpg2",
				RuntimeVersion: "1.14.0",
			},
				Logger: "all"},
		})
	if err != nil {
		t.Logf("TestERROR: %s", err)
	}

	if resp != nil {
		t.Logf("RESPONSE.Message: %#v", resp.Message)
	}
}

func TestDeleteInfereneceService(t *testing.T) {

	c := newInferenceService(t)
	resp, err := c.DeleteInferenceService(
		&infsType.DeleteInferenceServiceRequest{
			Namespace:     "default",
			Inferencename: "mpg-sample-rrr",
		})
	if err != nil {
		t.Error(err)
	}

	if resp != nil && resp.Message == "Success" {
		t.Logf("RESPONSE.Inferencename: %#v", resp.Inferencename)
	}
}
