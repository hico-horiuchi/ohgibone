package sensu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResults(t *testing.T) {
	assert := assert.New(t)

	results, err := DefaultAPI.GetResults()
	assert.Nil(err)
	assert.NotEqual(len(results), 0)

	for _, result := range results {
		if result.Check.Name == "keepalive" {
			continue
		}
		assert.Equal(result.Client, "test")
		assert.Equal(result.Check.Name, "default")
		assert.Equal(result.Check.Command, "/etc/sensu/plugins/default.sh")
		assert.Equal(result.Check.Subscribers, []string{"test"})
		assert.Equal(result.Check.Interval, 1)
		assert.Equal(result.Check.Handlers, []string{"default"})
		assert.Equal(result.Check.Aggregate, true)
		assert.NotEqual(result.Check.Issued, 0)
		assert.NotEqual(result.Check.Executed, 0)
		assert.Equal(result.Check.Output, "Default WARNING\n")
		assert.Equal(result.Check.Status, 1)
		assert.NotEqual(result.Check.Duration, 0.0)
	}
}

func TestGetResultsClient(t *testing.T) {
	assert := assert.New(t)

	results, err := DefaultAPI.GetResultsClient("test")
	assert.Nil(err)
	assert.NotEqual(len(results), 0)

	for _, result := range results {
		if result.Check.Name == "keepalive" {
			continue
		}
		assert.Equal(result.Client, "test")
		assert.Equal(result.Check.Name, "default")
		assert.Equal(result.Check.Command, "/etc/sensu/plugins/default.sh")
		assert.Equal(result.Check.Subscribers, []string{"test"})
		assert.Equal(result.Check.Interval, 1)
		assert.Equal(result.Check.Handlers, []string{"default"})
		assert.Equal(result.Check.Aggregate, true)
		assert.NotEqual(result.Check.Issued, 0)
		assert.NotEqual(result.Check.Executed, 0)
		assert.Equal(result.Check.Output, "Default WARNING\n")
		assert.Equal(result.Check.Status, 1)
		assert.NotEqual(result.Check.Duration, 0.0)
	}

	results, err = DefaultAPI.GetResultsClient("production")
	assert.Equal(err.Error(), "sensu: Not Found")
	assert.Len(results, 0)
}

func TestGetResultsClientCheck(t *testing.T) {
	assert := assert.New(t)

	result, err := DefaultAPI.GetResultsClientCheck("test", "default")
	assert.Nil(err)
	assert.Equal(result.Client, "test")
	assert.Equal(result.Check.Name, "default")
	assert.Equal(result.Check.Command, "/etc/sensu/plugins/default.sh")
	assert.Equal(result.Check.Subscribers, []string{"test"})
	assert.Equal(result.Check.Interval, 1)
	assert.Equal(result.Check.Handlers, []string{"default"})
	assert.Equal(result.Check.Aggregate, true)
	assert.NotEqual(result.Check.Issued, 0)
	assert.NotEqual(result.Check.Executed, 0)
	assert.Equal(result.Check.Output, "Default WARNING\n")
	assert.Equal(result.Check.Status, 1)
	assert.NotEqual(result.Check.Duration, 0.0)

	_, err = DefaultAPI.GetResultsClientCheck("test", "custom")
	assert.Equal(err.Error(), "sensu: Not Found")
}
