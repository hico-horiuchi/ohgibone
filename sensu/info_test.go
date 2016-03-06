package sensu

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInfo(t *testing.T) {
	assert := assert.New(t)

	info, err := DefaultAPI.GetInfo()
	assert.Nil(err)
	assert.Equal(info.Sensu.Version, "0.22.1")
	assert.Equal(info.Transport.Keepalives.Messages, 0)
	assert.Equal(info.Transport.Keepalives.Consumers, 1)
	assert.Equal(info.Transport.Results.Messages, 0)
	assert.Equal(info.Transport.Results.Consumers, 1)
	assert.Equal(info.Transport.Connected, true)
	assert.Equal(info.Redis.Connected, true)

	_, err = testAPI.GetInfo()
	assert.Contains(err.Error(), "getsockopt: connection refused")

	server, api := testServerAndAPI(http.StatusInternalServerError, "")
	defer server.Close()
	_, err = api.GetInfo()
	assert.Equal(err.Error(), "sensu: Internal Server Error")

	server, api = testServerAndAPI(http.StatusOK, "")
	defer server.Close()
	_, err = api.GetInfo()
	assert.Equal(err.Error(), "unexpected end of JSON input")
}
