package sensu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInfo(t *testing.T) {
	assert := assert.New(t)

	info, err := DefaultAPI.GetInfo()
	assert.Nil(err)
	assert.Equal(info.Sensu.Version, "0.22.0")
	assert.Equal(info.Transport.Keepalives.Messages, 0)
	assert.Equal(info.Transport.Keepalives.Consumers, 1)
	assert.Equal(info.Transport.Results.Messages, 0)
	assert.Equal(info.Transport.Results.Consumers, 1)
	assert.Equal(info.Transport.Connected, true)
	assert.Equal(info.Redis.Connected, true)
}
