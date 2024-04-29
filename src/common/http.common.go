package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// HTTPClient represents an HTTP client.
type HTTPClient struct{}

// NewHTTPClient creates a new instance of HTTPClient.
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{}
}

// SendRequest sends an HTTP request with the specified method, URL, and optional body.
func (c *HTTPClient) SendRequest(method, url string, body interface{}) ([]byte, error) {
	var requestBody []byte
	if body != nil {
		var err error
		requestBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

// Post sends an HTTP POST request to the specified URL with the given body.
func (c *HTTPClient) Post(url string, body interface{}) ([]byte, error) {
	return c.SendRequest(http.MethodPost, url, body)
}

// Get sends an HTTP GET request to the specified URL.
func (c *HTTPClient) Get(url string) ([]byte, error) {
	return c.SendRequest(http.MethodGet, url, nil)
}

// Delete sends an HTTP DELETE request to the specified URL.
func (c *HTTPClient) Delete(url string) ([]byte, error) {
	return c.SendRequest(http.MethodDelete, url, nil)
}

// Patch sends an HTTP PATCH request to the specified URL with the given body.
func (c *HTTPClient) Patch(url string, body interface{}) ([]byte, error) {
	return c.SendRequest(http.MethodPatch, url, body)
}
