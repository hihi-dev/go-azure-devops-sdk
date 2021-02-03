package ado

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	auth    string
	baseUrl string
	headers map[string]string
}

func CreateClient(username, password, baseUrl string) *Client {
	return &Client{base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password))), baseUrl, map[string]string{}}
}

func CreateClientWithToken(authToken, baseUrl string) *Client {
	return &Client{authToken, baseUrl, map[string]string{}}
}

// Perform an action on the API against this path
func (c *Client) doRequest(method string, path string, body io.Reader) (*http.Response, error) {
	c.headers["Accept"] = "application/json"
	c.headers["Authorization"] = "Basic " + c.auth
	url := c.baseUrl + path
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, body)
	for k, v := range c.headers {
		req.Header.Add(k, v)
	}
	return client.Do(req)
}
