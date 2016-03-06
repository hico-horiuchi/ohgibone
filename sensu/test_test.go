package sensu

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestServerAndAPI(t *testing.T) {
	assert := assert.New(t)

	server, api := testServerAndAPI(http.StatusOK, "")
	defer server.Close()
	assert.Equal(server.URL, "http://"+api.Host+":"+strconv.FormatInt(int64(api.Port), 10))
}
