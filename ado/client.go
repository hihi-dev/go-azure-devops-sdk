package ado

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	auth    string
	baseUrl string
	headers map[string]string
}

func CreateClient(username, pat, org string) *Client {
	return CreateSelfHostedClient(username, pat, createBaseUrl(org))
}

func CreateClientWithEncodedToken(authToken, org string) *Client {
	return CreateSelfHostedClientWithEncodedToken(authToken, createBaseUrl(org))
}

func CreateSelfHostedClient(username, pat, url string) *Client {
	return &Client{base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, pat))), url, map[string]string{}}
}

func CreateSelfHostedClientWithEncodedToken(authToken, url string) *Client {
	return &Client{authToken, url, map[string]string{}}
}

func createBaseUrl(org string) string {
	return "https://vsrm.dev.azure.com/"+org
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

func (c *Client) doRequestForBody(method string, path string, body io.Reader) ([]byte, error) {
	resp, err := c.doRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}
