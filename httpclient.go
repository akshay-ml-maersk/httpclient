package httpclient

import (
	"net/http"
	"time"
)

// Client is a struct representing an HTTP client.
type Client struct {
	client *http.Client
}

// NewClient returns a new instance of the HTTP client.
func NewClient() *Client {
	return &Client{
		client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

// Get sends an HTTP GET request and returns the response.
func (c *Client) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
