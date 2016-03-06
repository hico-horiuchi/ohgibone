package sensu

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPostResolve(t *testing.T) {
	assert := assert.New(t)

	err := DefaultAPI.PostResolve("test", "default")
	assert.Nil(err)

	err = DefaultAPI.PostResolve("test", "custom")
	assert.Equal(err.Error(), "sensu: Not Found")

	err = testAPI.PostResolve("test", "default")
	assert.Contains(err.Error(), "getsockopt: connection refused")

	server, api := testServerAndAPI(http.StatusInternalServerError, "")
	defer server.Close()
	err = api.PostResolve("test", "default")
	assert.Equal(err.Error(), "sensu: Internal Server Error")

	time.Sleep(1 * time.Second)
}
