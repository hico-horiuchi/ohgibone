package sensu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHealth(t *testing.T) {
	assert := assert.New(t)

	err := DefaultAPI.GetHealth(1, 1)
	assert.Nil(err)
}
