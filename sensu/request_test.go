package sensu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostRequest(t *testing.T) {
	assert := assert.New(t)

	err := DefaultAPI.PostRequest("default", []string{})
	assert.Nil(err)

	err = DefaultAPI.PostRequest("custom", []string{})
	assert.Equal(err.Error(), "sensu: Not Found")
}
