package models

import (
	"io/ioutil"
	"net/http"
	"time"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: 10 * time.Second,
	}
}

func HttpGet(address string, header map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}

	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return nil, Errorf("Request failed, code:%v", rsp.StatusCode)
	}

	result, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, Errorf("Read response failed: %v, code:%v", err, rsp.StatusCode)
	}

	return result, nil
}
