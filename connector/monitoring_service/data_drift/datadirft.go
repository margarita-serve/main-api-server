package data_drift

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	monType "git.k3.acornsoft.io/msit-auto-ml/koreserv/connector/monitoring_service/data_drift/types"
)

func NewDriftMonitor(config Config, httpClient *http.Client) *DriftMonitor {
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

	return &DriftMonitor{httpClient: httpClient, config: config}
}

type DriftMonitor struct {
	httpClient *http.Client
	config     Config
}

func (c *DriftMonitor) doRequest(req *http.Request) ([]byte, error) {
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

func (c *DriftMonitor) getRequest(url string, body io.Reader) ([]byte, error) {
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

func (c *DriftMonitor) postRequest(url string, body []byte) ([]byte, error) {
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

func (c *DriftMonitor) delRequest(url string, body io.Reader) ([]byte, error) {
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

func (c *DriftMonitor) putRequest(url string, body []byte) ([]byte, error) {
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

func (c *DriftMonitor) patchRequest(url string, body []byte) ([]byte, error) {
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

// get datadrift service
func (c *DriftMonitor) GetFeatureDrift(req *monType.GetFeatureDriftRequest) (*monType.GetFeatureDriftResponse, error) {
	module := fmt.Sprintf("drift-monitor/feature-drift/%s?start_time=%s&end_time=%s&model_history_id=%s", req.InferenceName, req.StartTime, req.EndTime, req.ModelHistoryID)
	env := c.getDriftEnv()
	url := fmt.Sprintf("%s/%s", env.ConnectionInfo, module)

	resp, err := c.getRequest(url, nil)
	if err != nil {
		return nil, err
	}

	respObj := monType.GetFeatureDriftResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}
	return &respObj, nil
}

func (c *DriftMonitor) GetFeatureDetail(req *monType.GetFeatureDriftRequest) (*monType.GetFeatureDriftResponse, error) {
	module := fmt.Sprintf("drift-monitor/feature-detail/%s?start_time=%s&end_time=%s&model_history_id=%s", req.InferenceName, req.StartTime, req.EndTime, req.ModelHistoryID)
	env := c.getDriftEnv()
	url := fmt.Sprintf("%s/%s", env.ConnectionInfo, module)

	resp, err := c.getRequest(url, nil)
	if err != nil {
		return nil, err
	}

	respObj := monType.GetFeatureDriftResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}
	return &respObj, nil
}

// create datadrift service
func (c *DriftMonitor) CreateDriftMonitor(req *monType.CreateDataDriftRequest) (*monType.CreateDataDriftResponse, error) {
	module := fmt.Sprintf("drift-monitor")
	env := c.getDriftEnv()
	url := fmt.Sprintf("%s/%s", env.ConnectionInfo, module)

	resp, err := c.postRequest(url, req.ToJSON())
	if err != nil {
		return nil, err
	}

	respObj := monType.CreateDataDriftResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}
	return &respObj, nil
}

// patch datadrift service
func (c *DriftMonitor) PatchDriftMonitor(req *monType.PatchDriftMonitorSettingRequest) (*monType.PatchDriftMonitorSettingResponse, error) {
	module := fmt.Sprintf("drift-monitor/%s", req.InferenceName)
	env := c.getDriftEnv()
	url := fmt.Sprintf("%s/%s", env.ConnectionInfo, module)

	reqDto := monType.PatchDriftMonitorSettingRequestDTO{
		DriftThreshold:      req.DriftThreshold,
		ImportanceThreshold: req.ImportanceThreshold,
		MonitorRange:        req.MonitorRange,
		LowImpAtRiskCount:   req.LowImpAtRiskCount,
		LowImpFailingCount:  req.LowImpFailingCount,
		HighImpAtRiskCount:  req.HighImpAtRiskCount,
		HighImpFailingCount: req.HighImpFailingCount,
	}
	resp, err := c.patchRequest(url, reqDto.ToJSON())
	if err != nil {
		return nil, err
	}

	respObj := monType.PatchDriftMonitorSettingResponse{}
	err = json.Unmarshal(resp, &respObj)
	if err != nil {
		return nil, err
	}

	return &respObj, nil
}

func (c *DriftMonitor) EnableMonitor(req *monType.EnableMonitorRequest) (*monType.EnableMonitorResponse, error) {
	module := fmt.Sprintf("drift-monitor/enable-monitor/%s", req.InferenceName)
	env := c.getDriftEnv()
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

func (c *DriftMonitor) DisableMonitor(req *monType.DisableMonitorRequest) (*monType.DisableMonitorResponse, error) {
	module := fmt.Sprintf("drift-monitor/disable-monitor/%s", req.InferenceName)
	env := c.getDriftEnv()
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

func (c *DriftMonitor) getDriftEnv() *monType.DriftServerEnv {
	env := new(monType.DriftServerEnv)
	env.ConnectionInfo = c.config.Endpoint

	return env
}
