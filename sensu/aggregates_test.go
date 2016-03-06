package sensu

import (
	"net/http"
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

	_, err = testAPI.GetAggregates(-1, -1)
	assert.Contains(err.Error(), "getsockopt: connection refused")

	server, api := testServerAndAPI(http.StatusInternalServerError, "")
	defer server.Close()
	_, err = api.GetAggregates(-1, -1)
	assert.Equal(err.Error(), "sensu: Internal Server Error")

	server, api = testServerAndAPI(http.StatusOK, "")
	defer server.Close()
	_, err = api.GetAggregates(-1, -1)
	assert.Equal(err.Error(), "unexpected end of JSON input")
}

func TestGetAggregatesCheck(t *testing.T) {
	assert := assert.New(t)

	issues, err := DefaultAPI.GetAggregatesCheck("default", -1)
	assert.Nil(err)
	assert.NotEqual(len(issues), 0)

	issues, err = DefaultAPI.GetAggregatesCheck("custom", -1)
	assert.Empty(issues)
	assert.Equal(err.Error(), "sensu: Not Found")

	_, err = testAPI.GetAggregatesCheck("default", -1)
	assert.Contains(err.Error(), "getsockopt: connection refused")

	server, api := testServerAndAPI(http.StatusInternalServerError, "")
	defer server.Close()
	_, err = api.GetAggregatesCheck("default", -1)
	assert.Equal(err.Error(), "sensu: Internal Server Error")

	server, api = testServerAndAPI(http.StatusOK, "")
	defer server.Close()
	_, err = api.GetAggregatesCheck("default", -1)
	assert.Equal(err.Error(), "unexpected end of JSON input")
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

	_, err = testAPI.GetAggregatesCheckIssued("default", 0, "", false)
	assert.Contains(err.Error(), "getsockopt: connection refused")

	server, api := testServerAndAPI(http.StatusInternalServerError, "")
	defer server.Close()
	_, err = api.GetAggregatesCheckIssued("default", 0, "", false)
	assert.Equal(err.Error(), "sensu: Internal Server Error")

	server, api = testServerAndAPI(http.StatusOK, "")
	defer server.Close()
	_, err = api.GetAggregatesCheckIssued("default", 0, "", false)
	assert.Equal(err.Error(), "unexpected end of JSON input")
}

func TestDeleteAggregatesCheck(t *testing.T) {
	assert := assert.New(t)

	err := DefaultAPI.DeleteAggregatesCheck("default")
	assert.Nil(err)

	err = DefaultAPI.DeleteAggregatesCheck("custom")
	assert.Equal(err.Error(), "sensu: Not Found")

	err = testAPI.DeleteAggregatesCheck("default")
	assert.Contains(err.Error(), "getsockopt: connection refused")

	server, api := testServerAndAPI(http.StatusInternalServerError, "")
	defer server.Close()
	err = api.DeleteAggregatesCheck("default")
	assert.Equal(err.Error(), "sensu: Internal Server Error")
}
