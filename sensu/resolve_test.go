package sensu

import (
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

	time.Sleep(1 * time.Second)
}
