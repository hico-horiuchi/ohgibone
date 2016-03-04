package sensu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAggregates(t *testing.T) {
	assert := assert.New(t)

	aggregates, err := DefaultAPI.GetAggregates(-1, -1)
	assert.Nil(err)
	assert.Len(aggregates, 1)
	assert.Equal(aggregates[0].Check, "default")
	assert.NotEqual(len(aggregates[0].Issued), 0)
}

func TestGetAggregatesCheck(t *testing.T) {
	assert := assert.New(t)

	issues, err := DefaultAPI.GetAggregatesCheck("default", -1)
	assert.Nil(err)
	assert.NotEqual(len(issues), 0)

	issues, err = DefaultAPI.GetAggregatesCheck("custom", -1)
	assert.Empty(issues)
	assert.Equal(err.Error(), "sensu: Not Found")
}

func TestGetAggregatesCheckIssued(t *testing.T) {
	assert := assert.New(t)
	issues, _ := DefaultAPI.GetAggregatesCheck("default", -1)

	aggregate, err := DefaultAPI.GetAggregatesCheckIssued("default", issues[0], "output", true)
	assert.Nil(err)
	assert.Equal(aggregate.Ok, 0)
	assert.Equal(aggregate.Warning, 1)
	assert.Equal(aggregate.Critical, 0)
	assert.Equal(aggregate.Unknown, 0)
	assert.Equal(aggregate.Total, 1)
	assert.Equal(aggregate.Outputs["Default WARNING\n"], 1)
	assert.Equal(aggregate.Results[0].Client, "test")
	assert.Equal(aggregate.Results[0].Output, "Default WARNING\n")
	assert.Equal(aggregate.Results[0].Status, 1)

	_, err = DefaultAPI.GetAggregatesCheckIssued("default", 0, "", false)
	assert.Equal(err.Error(), "sensu: Not Found")
}

func TestDeleteAggregatesCheck(t *testing.T) {
	assert := assert.New(t)

	err := DefaultAPI.DeleteAggregatesCheck("default")
	assert.Nil(err)

	err = DefaultAPI.DeleteAggregatesCheck("custom")
	assert.Equal(err.Error(), "sensu: Not Found")
}
