package sensu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusCodeToString(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(statusCodeToString(200), "OK")
	assert.Equal(statusCodeToString(201), "Created")
	assert.Equal(statusCodeToString(202), "Accepted")
	assert.Equal(statusCodeToString(204), "No Content")
	assert.Equal(statusCodeToString(400), "Bad Request")
	assert.Equal(statusCodeToString(404), "Not Found")
	assert.Equal(statusCodeToString(500), "Internal Server Error")
	assert.Equal(statusCodeToString(503), "Service Unavailable")
}
