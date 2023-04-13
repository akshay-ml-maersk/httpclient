package pensieve

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type PensieveClient struct {
	baseURL string
}

func NewPensieveClient(baseURL string) *PensieveClient {
	return &PensieveClient{baseURL}
}

func (c *PensieveClient) SendRequest(method string, path string, requestBody interface{}) (interface{}, error) {
	url := c.baseURL + path

	var jsonRequest []byte
	var err error
	if requestBody != nil {
		jsonRequest, err = json.Marshal(requestBody)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %s", err)
		}
	}

	var req *http.Request
	if method == "POST" {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonRequest))
		if err != nil {
			return nil, fmt.Errorf("error creating HTTP request: %s", err)
		}
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return nil, fmt.Errorf("error creating HTTP request: %s", err)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("unexpected response status code: %d", resp.StatusCode)
	}

	var responseBody interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return nil, fmt.Errorf("error decoding response body: %s", err)
	}

	return responseBody, nil
}
