package predictionSender

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func NewPredictionSendService() (*PredictionSender, error) {

	netTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Timeout: time.Second * 600, Transport: netTransport}

	return &PredictionSender{httpClient: httpClient}, nil
}

// Covid19 Client SDK
type PredictionSender struct {
	httpClient *http.Client
}

func (c *PredictionSender) doRequest(req *http.Request) ([]byte, error) {
	fmt.Printf("req: %v\n", req)
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

func (c *PredictionSender) SendPrediction(url string, host string, data []byte) ([]byte, error) {
	fmt.Printf("body: %v\n", string(data))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-type", "application/json")
	req.Host = host

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	result := resp

	return result, nil
}
