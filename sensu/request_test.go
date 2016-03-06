package sensu

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostRequest(t *testing.T) {
	assert := assert.New(t)

	err := DefaultAPI.PostRequest("default", []string{})
	assert.Nil(err)

	err = DefaultAPI.PostRequest("custom", []string{})
	assert.Equal(err.Error(), "sensu: Not Found")

	err = testAPI.PostRequest("default", []string{})
	assert.Contains(err.Error(), "getsockopt: connection refused")

	server, api := testServerAndAPI(http.StatusInternalServerError, "")
	defer server.Close()
	err = api.PostRequest("default", []string{})
	assert.Equal(err.Error(), "sensu: Internal Server Error")
}
