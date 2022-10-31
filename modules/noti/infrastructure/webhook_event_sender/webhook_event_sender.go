package webHookEventSender

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

func NewWebHookEventSendService(h *handler.Handler) (*WebHookEventSender, error) {

	netTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Timeout: time.Second * 10, Transport: netTransport}

	return &WebHookEventSender{httpClient: httpClient}, nil
}

// Covid19 Client SDK
type WebHookEventSender struct {
	httpClient *http.Client
}

func (c *WebHookEventSender) doRequest(req *http.Request) ([]byte, error) {
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

func (c *WebHookEventSender) SendWebHookEvent(url string, method string, header string, bodyStr string) ([]byte, error) {
	fmt.Printf("body: %v\n", string(bodyStr))

	bodyBytes := []byte(bodyStr)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}

	procHeadersCR := strings.Split(strings.ReplaceAll(header, "\\n", "\n"), "\n")
	for _, str := range procHeadersCR {
		fmt.Println(str)
		slice := strings.Split(str, ":")
		req.Header.Add(strings.Trim(slice[0], " "), strings.Trim(slice[1], " "))
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	result := resp

	return result, nil
}
