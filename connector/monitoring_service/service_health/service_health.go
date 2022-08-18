package service_health

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	monType "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/monitoring_service/service_health/types"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func NewServiceHealthMonitor(config Config, httpClient *http.Client) *ServiceHealthMonitor {
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

	return &ServiceHealthMonitor{httpClient: httpClient, config: config}
}

type ServiceHealthMonitor struct {
	httpClient *http.Client
	config     Config
}

func (c *ServiceHealthMonitor) doRequest(req *http.Request) ([]byte, error) {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return []byte(`{}`), err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if res.StatusCode > 201 {
		return nil, fmt.Errorf("ERROR:[%v] %s; URL: %s", res.StatusCode, body, req.URL)
	}

	return body, err
}

func (c *ServiceHealthMonitor) getRequest(url string, body io.Reader) ([]byte, error) {
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

func (c *ServiceHealthMonitor) postRequest(url string, body []byte) ([]byte, error) {
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

func (c *ServiceHealthMonitor) delRequest(url string, body io.Reader) ([]byte, error) {
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

func (c *ServiceHealthMonitor) putRequest(url string, body []byte) ([]byte, error) {
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

func (c *ServiceHealthMonitor) patchRequest(url string, body []byte) ([]byte, error) {
	fmt.Printf("body: %v\n", string(body))
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
	req.Header.Add("content-type", "application/json")

	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetServiceHealth get service health service
func (c *ServiceHealthMonitor) GetServiceHealth(req *monType.GetServiceHealthRequest) (*monType.GetServiceHealthResponse, error) {
	module := fmt.Sprintf("servicehealth-monitor/servicehealth/%s?model_history_id=%s&type=%s&start_time=%s&end_time=%s", req.InferenceName, req.ModelHistoryID, req.DataType, req.StartTime, req.EndTime)
	env := c.getServiceHealthEnv()
	url := fmt.Sprintf("%s/%s", env.ConnectionInfo, module)

	resp, err := c.getRequest(url, nil)
	if err != nil {
		return nil, err
	}

	respObj := monType.GetServiceHealthResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}
	return &respObj, nil
}

// CreateServiceHealthMonitor create service health service
func (c *ServiceHealthMonitor) CreateServiceHealthMonitor(req *monType.CreateServiceHealthRequest) (*monType.CreateServiceHealthResponse, error) {
	module := "servicehealth-monitor"
	env := c.getServiceHealthEnv()
	url := fmt.Sprintf("%s/%s", env.ConnectionInfo, module)

	resp, err := c.postRequest(url, req.ToJSON())
	if err != nil {
		return nil, err
	}

	respObj := monType.CreateServiceHealthResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}
	return &respObj, nil
}

func (c *ServiceHealthMonitor) EnableMonitor(req *monType.EnableMonitorRequest) (*monType.EnableMonitorResponse, error) {
	module := fmt.Sprintf("servicehealth-monitor/enable-monitor/%s", req.InferenceName)
	env := c.getServiceHealthEnv()
	url := fmt.Sprintf("%s/%s", env.ConnectionInfo, module)

	resp, err := c.patchRequest(url, nil)
	if err != nil {
		return nil, err
	}

	respObj := monType.EnableMonitorResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}
	return &respObj, nil
}

func (c *ServiceHealthMonitor) DisableMonitor(req *monType.DisableMonitorRequest) (*monType.DisableMonitorResponse, error) {
	module := fmt.Sprintf("servicehealth-monitor/disable-monitor/%s", req.InferenceName)
	env := c.getServiceHealthEnv()
	url := fmt.Sprintf("%s/%s", env.ConnectionInfo, module)

	resp, err := c.patchRequest(url, nil)
	if err != nil {
		return nil, err
	}

	respObj := monType.DisableMonitorResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}
	return &respObj, nil
}

func (c *ServiceHealthMonitor) getServiceHealthEnv() *monType.ServiceHealthServerEnv {
	env := new(monType.ServiceHealthServerEnv)
	env.ConnectionInfo = c.config.Endpoint

	return env
}
