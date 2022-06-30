package kserve_cntr

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	infsType "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/inference_service/kserve_cntr/types"
)

// NewCovid19 new Covid19 Client SDK
func NewInferenceService(config Config, httpClient *http.Client) *InferenceService {

	if httpClient == nil {
		netTransport := &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		}
		httpClient = &http.Client{Timeout: time.Second * 600, Transport: netTransport}
	}

	//return &InferenceService{Server: config.Server, httpClient: httpClient}
	return &InferenceService{httpClient: httpClient}
}

// Covid19 Client SDK
type InferenceService struct {
	//Server     string
	httpClient *http.Client
}

func (c *InferenceService) doRequest(req *http.Request) ([]byte, error) {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return []byte(`{}`), err
	}
	defer res.Body.Close()

	if res.StatusCode > 201 {
		return nil, fmt.Errorf("ERROR:[%v] %s; URL: %s", res.StatusCode, res.Status, req.URL)
	}

	body, err := ioutil.ReadAll(res.Body)
	return body, err
}

func (c *InferenceService) getRequest(url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest("GET", url, body)
	req.Header.Add("Content-type", "application/json")

	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *InferenceService) postRequest(url string, body []byte) ([]byte, error) {
	fmt.Printf("body: %v\n", string(body))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Add("Content-type", "application/json")

	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *InferenceService) delRequest(url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest("DELETE", url, body)
	req.Header.Add("Content-type", "application/json")

	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *InferenceService) putRequest(url string, body []byte) ([]byte, error) {
	fmt.Printf("body: %v\n", string(body))
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	req.Header.Add("Content-type", "application/json")

	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetInferenceService get GetInferenceService
func (c *InferenceService) GetInferenceService(req *infsType.GetInferenceServiceRequest) (*infsType.GetInferenceServiceResponse, error) {
	url := fmt.Sprintf("inference-service/%s?namespace=%s", strings.ToLower(req.Inferencename), strings.ToLower(req.Namespace))

	resp, err := c.getRequest(url, nil)
	if err != nil {
		return nil, err
	}

	respObj := infsType.GetInferenceServiceResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}

	return &respObj, nil
}

// CreateInferenceService Post CreateInferenceService
func (c *InferenceService) CreateInferenceService(req *infsType.CreateInferenceServiceRequest) (*infsType.CreateInferenceServiceResponse, error) {
	module := "inference-service"
	url := fmt.Sprintf("%s/%s", req.InferenceServer, module)

	resp, err := c.postRequest(url, req.ToJSON())
	if err != nil {
		return nil, err
	}

	respObj := infsType.CreateInferenceServiceResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}

	return &respObj, nil
}

// GetInferenceService get GetInferenceService
func (c *InferenceService) DeleteInferenceService(req *infsType.DeleteInferenceServiceRequest) (*infsType.DeleteInferenceServiceResponse, error) {

	module := fmt.Sprintf("inference-service/%s?namespace=%s", strings.ToLower(req.Inferencename), strings.ToLower(req.Namespace))
	url := fmt.Sprintf("%s/%s", req.InferenceServer, module)

	resp, err := c.delRequest(url, nil)
	if err != nil {
		return nil, err
	}

	respObj := infsType.DeleteInferenceServiceResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}

	return &respObj, nil
}

// UpdateInferenceService Post UpdateInferenceService
func (c *InferenceService) UpdateInferenceService(req *infsType.UpdateInferenceServiceRequest) (*infsType.UpdateInferenceServiceResponse, error) {
	module := "inference-service"
	url := fmt.Sprintf("%s/%s", req.InferenceServer, module)

	resp, err := c.putRequest(url, req.ToJSON())
	if err != nil {
		return nil, err
	}

	respObj := infsType.UpdateInferenceServiceResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}

	return &respObj, nil
}
