package accuracy

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	monType "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/monitoring_service/accuracy/types"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func NewAccuracyMonitor(config Config, httpClient *http.Client) *AccuracyMonitor {
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

	return &AccuracyMonitor{httpClient: httpClient, config: config}
}

type AccuracyMonitor struct {
	httpClient *http.Client
	config     Config
}

func (c *AccuracyMonitor) doRequest(req *http.Request) ([]byte, error) {
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

func (c *AccuracyMonitor) getRequest(url string, body io.Reader) ([]byte, error) {
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

func (c *AccuracyMonitor) postRequest(url string, body []byte) ([]byte, error) {
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

func (c *AccuracyMonitor) delRequest(url string, body io.Reader) ([]byte, error) {
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

func (c *AccuracyMonitor) putRequest(url string, body []byte) ([]byte, error) {
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

func (c *AccuracyMonitor) patchRequest(url string, body []byte) ([]byte, error) {
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

// get accuracy service
func (c *AccuracyMonitor) GetAccuracy(req *monType.GetAccuracyRequest) (*monType.GetAccuracyResponse, error) {
	module := fmt.Sprintf("accuracy-monitor/accuracy/%s?model_history_id=%s&type=%s&start_time=%s&end_time=%s", req.InferenceName, req.ModelHistoryID, req.DataType, req.StartTime, req.EndTime)
	env := c.getAccuracyEnv()
	url := fmt.Sprintf("%s/%s", env.ConnectionInfo, module)

	resp, err := c.getRequest(url, nil)
	if err != nil {
		return nil, err
	}

	respObj := monType.GetAccuracyResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}
	return &respObj, nil
}

// create accuracy service
func (c *AccuracyMonitor) CreateAccuracyMonitor(req *monType.CreateAccuracyRequest) (*monType.CreateAccuracyResponse, error) {
	module := "accuracy-monitor"
	env := c.getAccuracyEnv()
	//url := fmt.Sprintf("%s/%s", "http://192.168.88.151:30072", module)
	url := fmt.Sprintf("%s/%s", env.ConnectionInfo, module)

	resp, err := c.postRequest(url, req.ToJSON())
	if err != nil {
		return nil, err
	}

	respObj := monType.CreateAccuracyResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}
	return &respObj, nil
}

// patch accuracy service
func (c *AccuracyMonitor) PatchAccuracyMonitor(req *monType.PatchAccuracySettingRequest) (*monType.PatchAccuracySettingResponse, error) {
	module := fmt.Sprintf("accuracy-monitor/%s", req.InferenceName)
	env := c.getAccuracyEnv()
	url := fmt.Sprintf("%s/%s", env.ConnectionInfo, module)

	reqDto := monType.PatchAccuracySettingRequestDTO{
		DriftMetrics:     req.DriftMetrics,
		DriftMeasurement: req.DriftMeasurement,
		AtriskValue:      req.AtriskValue,
		FailingValue:     req.FailingValue,
	}
	resp, err := c.patchRequest(url, reqDto.ToJSON())
	if err != nil {
		return nil, err
	}

	respObj := monType.PatchAccuracySettingResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}

	return &respObj, nil
}

// post actual service
func (c *AccuracyMonitor) PostActual(req *monType.ActualRequest) (*monType.ActualResponse, error) {
	module := fmt.Sprintf("accuracy-monitor/actual")
	env := c.getAccuracyEnv()
	url := fmt.Sprintf("%s/%s", env.ConnectionInfo, module)

	resp, err := c.postRequest(url, req.ToJSON())
	if err != nil {
		return nil, err
	}

	respObj := monType.ActualResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}

	return &respObj, nil
}

func (c *AccuracyMonitor) EnableMonitor(req *monType.EnableMonitorRequest) (*monType.EnableMonitorResponse, error) {
	module := fmt.Sprintf("accuracy-monitor/enable-monitor/%s", req.InferenceName)
	env := c.getAccuracyEnv()
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

func (c *AccuracyMonitor) DisableMonitor(req *monType.DisableMonitorRequest) (*monType.DisableMonitorResponse, error) {
	module := fmt.Sprintf("accuracy-monitor/disable-monitor/%s", req.InferenceName)
	env := c.getAccuracyEnv()
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

func (c *AccuracyMonitor) UpdateAssociationID(req *monType.UpdateAssociationIDRequest) (*monType.UpdateAssociationIDResponse, error) {
	module := fmt.Sprintf("accuracy-monitor/%s/association-id", req.InferenceName)
	env := c.getAccuracyEnv()
	url := fmt.Sprintf("%s/%s", env.ConnectionInfo, module)

	reqDto := monType.UpdateAssociationIDRequestDTO{
		AssociationID: req.AssociationID,
	}

	resp, err := c.patchRequest(url, reqDto.ToJSON())
	if err != nil {
		return nil, err
	}

	respObj := monType.UpdateAssociationIDResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}
	return &respObj, nil
}

func (c *AccuracyMonitor) getAccuracyEnv() *monType.AccuracyServerEnv {
	env := new(monType.AccuracyServerEnv)
	env.ConnectionInfo = c.config.Endpoint

	return env
}
