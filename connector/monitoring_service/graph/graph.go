package graph

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	monType "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/monitoring_service/graph/types"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func NewGraphMonitor(config Config, httpClient *http.Client) *GraphMonitor {
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

	return &GraphMonitor{httpClient: httpClient, config: config}
}

type GraphMonitor struct {
	httpClient *http.Client
	config     Config
}

func (c *GraphMonitor) doRequest(req *http.Request) ([]byte, error) {
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

func (c *GraphMonitor) getRequest(url string, body io.Reader) ([]byte, error) {
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

func (c *GraphMonitor) postRequest(url string, body []byte) ([]byte, error) {
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

func (c *GraphMonitor) delRequest(url string, body io.Reader) ([]byte, error) {
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

func (c *GraphMonitor) putRequest(url string, body []byte) ([]byte, error) {
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

func (c *GraphMonitor) patchRequest(url string, body []byte) ([]byte, error) {
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

func (c *GraphMonitor) GetFeatureDetailGraph(req *monType.GetDetailGraphRequest) (*monType.GetDetailGraphResponse, error) {
	module := fmt.Sprintf("detail/%s?start_time=%s&end_time=%s&model_history_id=%s", req.InferenceName, req.StartTime, req.EndTime, req.ModelHistoryID)
	env := c.getGraphEnv()
	url := fmt.Sprintf("%s/%s", env.ConnectionInfo, module)

	resp, err := c.getRequest(url, nil)
	if err != nil {
		return nil, err
	}

	respObj := monType.GetDetailGraphResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}
	return &respObj, nil
}

func (c *GraphMonitor) GetDataDriftGraph(req *monType.GetDriftGraphRequest) (*monType.GetDriftGraphResponse, error) {
	module := fmt.Sprintf("drift/%s?start_time=%s&end_time=%s&model_history_id=%s&drift_threshold=%f&importance_threshold=%f",
		req.InferenceName, req.StartTime, req.EndTime, req.ModelHistoryID, req.DriftThreshold, req.ImportanceThreshold)
	env := c.getGraphEnv()
	url := fmt.Sprintf("%s/%s", env.ConnectionInfo, module)

	resp, err := c.getRequest(url, nil)
	if err != nil {
		return nil, err
	}

	respObj := monType.GetDriftGraphResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}
	return &respObj, nil
}

func (c *GraphMonitor) getGraphEnv() *monType.GraphServerEnv {
	env := new(monType.GraphServerEnv)
	env.ConnectionInfo = c.config.Endpoint

	return env
}
