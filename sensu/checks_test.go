package sensu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetChecks(t *testing.T) {
	assert := assert.New(t)

	checks, err := DefaultAPI.GetChecks()
	assert.Nil(err)
	assert.Len(checks, 1)
	assert.Equal(checks[0].Name, "default")
	assert.Equal(checks[0].Command, "/etc/sensu/plugins/default.sh")
	assert.Equal(checks[0].Subscribers, []string{"test"})
	assert.Equal(checks[0].Interval, 1)
	assert.Equal(checks[0].Handlers, []string{"default"})
	assert.Equal(checks[0].Aggregate, true)
}

func TestGetChecksCheck(t *testing.T) {
	assert := assert.New(t)

	check, err := DefaultAPI.GetChecksCheck("default")
	assert.Nil(err)
	assert.Equal(check.Name, "default")
	assert.Equal(check.Command, "/etc/sensu/plugins/default.sh")
	assert.Equal(check.Subscribers, []string{"test"})
	assert.Equal(check.Interval, 1)
	assert.Equal(check.Handlers, []string{"default"})
	assert.Equal(check.Aggregate, true)

	_, err = DefaultAPI.GetChecksCheck("custom")
	assert.Equal(err.Error(), "sensu: Not Found")
}
