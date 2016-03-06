package sensu

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClientsHistory(t *testing.T) {
	assert := assert.New(t)

	histories, err := DefaultAPI.GetClientsHistory("test")
	assert.Nil(err)
	assert.NotEqual(len(histories), 0)

	for _, history := range histories {
		if history.Check == "keepalive" {
			continue
		}
		assert.Equal(history.Check, "default")
		assert.NotEqual(len(history.History), 0)
		assert.Equal(history.History[0], 1)
		assert.NotEqual(history.LastExecution, 0)
		assert.Equal(history.LastStatus, 1)
	}

	histories, err = DefaultAPI.GetClientsHistory("production")
	assert.Nil(err)
	assert.Empty(histories)

	_, err = testAPI.GetClientsHistory("test")
	assert.Contains(err.Error(), "getsockopt: connection refused")

	server, api := testServerAndAPI(http.StatusInternalServerError, "")
	defer server.Close()
	_, err = api.GetClientsHistory("test")
	assert.Equal(err.Error(), "sensu: Internal Server Error")

	server, api = testServerAndAPI(http.StatusOK, "")
	defer server.Close()
	_, err = api.GetClientsHistory("test")
	assert.Equal(err.Error(), "unexpected end of JSON input")
}
