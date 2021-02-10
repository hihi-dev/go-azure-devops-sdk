package ado

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	auth    string
	org     string
	headers map[string]string
}

// Seems that the DevOps API doesn't have a constant BaseUrl, so this exists :(
const BaseUrl = "https://dev.azure.com/"
const BaseUrlLegacy = "https://vsrm.dev.azure.com/"

func CreateClient(username, pat, org string) *Client {
	return &Client{base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, pat))), org, map[string]string{}}
}

func CreateClientWithEncodedToken(authToken, org string) *Client {
	return &Client{authToken, org, map[string]string{}}
}

// Perform an action on the API against this path
func (c *Client) doRequest(baseUrl, method string, path string, body io.Reader) (*http.Response, error) {
	c.headers["Accept"] = "application/json"
	c.headers["Authorization"] = "Basic " + c.auth
	url := baseUrl + path
	log.Println("Requesting from the following URL")
	log.Println(url)
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, body)
	for k, v := range c.headers {
		req.Header.Add(k, v)
	}
	return client.Do(req)
}

func (c *Client) doRequestForBody(baseUrl, method string, path string, body io.Reader) ([]byte, error) {
	resp, err := c.doRequest(baseUrl, method, path, body)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}
