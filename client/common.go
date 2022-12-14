package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (client *Client) doRequest(method string, relativePath string) (*http.Response, error) {
	fullURL := fmt.Sprintf("%s%s", client.apiUrl, relativePath)
	return DoRequestBearer(method, fullURL, nil, nil, client.token)
}

func (client *Client) doJSONRequest(method string, relativePath string, requestBody interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	fullURL := fmt.Sprintf("%s%s", client.apiUrl, relativePath)
	return DoRequestBearer(method, fullURL, jsonBody, nil, client.token)
}

func (client *Client) doJSONRequestUnauthorized(method string, relativePath string, requestBody interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	fullURL := fmt.Sprintf("%s%s", client.apiUrl, relativePath)
	return DoRequest(method, fullURL, jsonBody, nil)
}
