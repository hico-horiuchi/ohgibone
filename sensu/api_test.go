package sensu

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	assert := assert.New(t)

	response, err := DefaultAPI.get("/test")
	assert.Nil(err)
	assert.Equal(response.Body, "")
	assert.Equal(response.StatusCode, 404)
}

func TestPost(t *testing.T) {
	assert := assert.New(t)
	payload := strings.NewReader("/test")

	response, err := DefaultAPI.post("/test", payload)
	assert.Nil(err)
	assert.Equal(response.Body, "")
	assert.Equal(response.StatusCode, 404)
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)

	response, err := DefaultAPI.delete("/test")
	assert.Nil(err)
	assert.Equal(response.Body, "")
	assert.Equal(response.StatusCode, 404)
}

func TestDo(t *testing.T) {
	assert := assert.New(t)

	response, err := DefaultAPI.do("GET", "/test", nil)
	assert.Nil(err)
	assert.Equal(response.Body, "")
	assert.Equal(response.StatusCode, 404)

	_, err = testAPI.do("GET", "/test", nil)
	assert.Contains(err.Error(), "getsockopt: connection refused")
}

func TestNewRequest(t *testing.T) {
	assert := assert.New(t)
	DefaultAPI.User = "admin"
	DefaultAPI.Password = "secret"

	request, err := DefaultAPI.newRequest("GET", "/test", nil)
	assert.Nil(err)
	assert.Equal(request.Method, "GET")
	assert.Equal(request.Host, "localhost:4567")
	assert.Equal(request.URL.String(), "http://localhost:4567/test")
	assert.Equal(request.Header["Authorization"][0], "Basic YWRtaW46c2VjcmV0")
}
