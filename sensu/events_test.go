package sensu

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetEvents(t *testing.T) {
	assert := assert.New(t)

	events, err := DefaultAPI.GetEvents()
	assert.Nil(err)
	assert.Len(events, 1)
	assert.NotEqual(events[0].ID, "")
	assert.Equal(events[0].Client.Name, "test")
	assert.Equal(events[0].Client.Address, "localhost")
	assert.Equal(events[0].Client.Subscriptions, []string{"test"})
	assert.NotEqual(events[0].Client.Timestamp, 0)
	assert.Equal(events[0].Client.Version, "0.22.1")
	assert.Equal(events[0].Check.Name, "default")
	assert.Equal(events[0].Check.Command, "/etc/sensu/plugins/default.sh")
	assert.Equal(events[0].Check.Subscribers, []string{"test"})
	assert.Equal(events[0].Check.Interval, 1)
	assert.Equal(events[0].Check.Handlers, []string{"default"})
	assert.Equal(events[0].Check.Aggregate, true)
	assert.NotEqual(events[0].Check.Issued, 0)
	assert.NotEqual(events[0].Check.Executed, 0)
	assert.Equal(events[0].Check.Output, "Default WARNING\n")
	assert.Equal(events[0].Check.Status, 1)
	assert.NotEqual(events[0].Check.Duration, 0.0)
	assert.NotEqual(len(events[0].Check.History), 0)
	assert.Equal(events[0].Check.History[0], "1")
	assert.NotEqual(events[0].Occurrences, 0)
	assert.Equal(events[0].Action, "create")

	_, err = testAPI.GetEvents()
	assert.Contains(err.Error(), "getsockopt: connection refused")

	server, api := testServerAndAPI(http.StatusInternalServerError, "")
	defer server.Close()
	_, err = api.GetEvents()
	assert.Equal(err.Error(), "sensu: Internal Server Error")

	server, api = testServerAndAPI(http.StatusOK, "")
	defer server.Close()
	_, err = api.GetEvents()
	assert.Equal(err.Error(), "unexpected end of JSON input")
}

func TestGetEventsClient(t *testing.T) {
	assert := assert.New(t)

	events, err := DefaultAPI.GetEventsClient("test")
	assert.Nil(err)
	assert.Len(events, 1)
	assert.NotEqual(events[0].ID, "")
	assert.Equal(events[0].Client.Name, "test")
	assert.Equal(events[0].Client.Address, "localhost")
	assert.Equal(events[0].Client.Subscriptions, []string{"test"})
	assert.NotEqual(events[0].Client.Timestamp, 0)
	assert.Equal(events[0].Client.Version, "0.22.1")
	assert.Equal(events[0].Check.Name, "default")
	assert.Equal(events[0].Check.Command, "/etc/sensu/plugins/default.sh")
	assert.Equal(events[0].Check.Subscribers, []string{"test"})
	assert.Equal(events[0].Check.Interval, 1)
	assert.Equal(events[0].Check.Handlers, []string{"default"})
	assert.Equal(events[0].Check.Aggregate, true)
	assert.NotEqual(events[0].Check.Issued, 0)
	assert.NotEqual(events[0].Check.Executed, 0)
	assert.Equal(events[0].Check.Output, "Default WARNING\n")
	assert.Equal(events[0].Check.Status, 1)
	assert.NotEqual(events[0].Check.Duration, 0.0)
	assert.NotEqual(len(events[0].Check.History), 0)
	assert.Equal(events[0].Check.History[0], "1")
	assert.NotEqual(events[0].Occurrences, 0)
	assert.Equal(events[0].Action, "create")

	events, err = DefaultAPI.GetEventsClient("production")
	assert.Nil(err)
	assert.Empty(events)

	_, err = testAPI.GetEventsClient("test")
	assert.Contains(err.Error(), "getsockopt: connection refused")

	server, api := testServerAndAPI(http.StatusInternalServerError, "")
	defer server.Close()
	_, err = api.GetEventsClient("test")
	assert.Equal(err.Error(), "sensu: Internal Server Error")

	server, api = testServerAndAPI(http.StatusOK, "")
	defer server.Close()
	_, err = api.GetEventsClient("test")
	assert.Equal(err.Error(), "unexpected end of JSON input")
}

func TestGetEventsClientCheck(t *testing.T) {
	assert := assert.New(t)

	event, err := DefaultAPI.GetEventsClientCheck("test", "default")
	assert.Nil(err)
	assert.NotEqual(event.ID, "")
	assert.Equal(event.Client.Name, "test")
	assert.Equal(event.Client.Address, "localhost")
	assert.Equal(event.Client.Subscriptions, []string{"test"})
	assert.NotEqual(event.Client.Timestamp, 0)
	assert.Equal(event.Client.Version, "0.22.1")
	assert.Equal(event.Check.Name, "default")
	assert.Equal(event.Check.Command, "/etc/sensu/plugins/default.sh")
	assert.Equal(event.Check.Subscribers, []string{"test"})
	assert.Equal(event.Check.Interval, 1)
	assert.Equal(event.Check.Handlers, []string{"default"})
	assert.Equal(event.Check.Aggregate, true)
	assert.NotEqual(event.Check.Issued, 0)
	assert.NotEqual(event.Check.Executed, 0)
	assert.Equal(event.Check.Output, "Default WARNING\n")
	assert.Equal(event.Check.Status, 1)
	assert.NotEqual(event.Check.Duration, 0.0)
	assert.NotEqual(len(event.Check.History), 0)
	assert.Equal(event.Check.History[0], "1")
	assert.NotEqual(event.Occurrences, 0)
	assert.Equal(event.Action, "create")

	_, err = DefaultAPI.GetEventsClientCheck("test", "custom")
	assert.Equal(err.Error(), "sensu: Not Found")

	_, err = testAPI.GetEventsClientCheck("test", "default")
	assert.Contains(err.Error(), "getsockopt: connection refused")

	server, api := testServerAndAPI(http.StatusInternalServerError, "")
	defer server.Close()
	_, err = api.GetEventsClientCheck("test", "default")
	assert.Equal(err.Error(), "sensu: Internal Server Error")

	server, api = testServerAndAPI(http.StatusOK, "")
	defer server.Close()
	_, err = api.GetEventsClientCheck("test", "default")
	assert.Equal(err.Error(), "unexpected end of JSON input")
}

func TestDeleteEventsClientCheck(t *testing.T) {
	assert := assert.New(t)

	err := DefaultAPI.DeleteEventsClientCheck("test", "default")
	assert.Nil(err)

	err = DefaultAPI.DeleteEventsClientCheck("test", "custom")
	assert.Equal(err.Error(), "sensu: Not Found")

	err = testAPI.DeleteEventsClientCheck("test", "default")
	assert.Contains(err.Error(), "getsockopt: connection refused")

	server, api := testServerAndAPI(http.StatusInternalServerError, "")
	defer server.Close()
	err = api.DeleteEventsClientCheck("test", "default")
	assert.Equal(err.Error(), "sensu: Internal Server Error")

	time.Sleep(1 * time.Second)
}
