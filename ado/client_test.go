package ado

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateClient_SendingAuthHeader(t *testing.T) {
	c := CreateClient("username", "password", "http://localhost")
	// Do a request to trigger generating the Authorisation header
	c.doRequest("", "", nil)
	assert.Contains(t, c.headers, "Authorization")
	assert.Equal(t, c.headers["Authorization"], "Basic dXNlcm5hbWU6cGFzc3dvcmQ=")
}
