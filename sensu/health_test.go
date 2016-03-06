package sensu

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHealth(t *testing.T) {
	assert := assert.New(t)

	err := DefaultAPI.GetHealth(1, 1)
	assert.Nil(err)

	err = testAPI.GetHealth(1, 1)
	assert.Contains(err.Error(), "getsockopt: connection refused")

	server, api := testServerAndAPI(http.StatusInternalServerError, "")
	defer server.Close()
	err = api.GetHealth(1, 1)
	assert.Equal(err.Error(), "sensu: Internal Server Error")
}
