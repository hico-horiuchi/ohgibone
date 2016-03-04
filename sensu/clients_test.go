package sensu

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetClients(t *testing.T) {
	assert := assert.New(t)

	clients, err := DefaultAPI.GetClients(-1, -1)
	assert.Nil(err)
	assert.Len(clients, 1)
	assert.Equal(clients[0].Name, "test")
	assert.Equal(clients[0].Address, "localhost")
	assert.Equal(clients[0].Subscriptions, []string{"test"})
	assert.NotEqual(clients[0].Timestamp, 0)
	assert.Equal(clients[0].Version, "0.22.0")
}

func TestGetClientsClient(t *testing.T) {
	assert := assert.New(t)

	client, err := DefaultAPI.GetClientsClient("test")
	assert.Nil(err)
	assert.Equal(client.Name, "test")
	assert.Equal(client.Address, "localhost")
	assert.Equal(client.Subscriptions, []string{"test"})
	assert.NotEqual(client.Timestamp, 0)
	assert.Equal(client.Version, "0.22.0")

	_, err = DefaultAPI.GetClientsClient("custom")
	assert.Equal(err.Error(), "sensu: Not Found")
}

func TestPostClients(t *testing.T) {
	assert := assert.New(t)

	err := DefaultAPI.PostClients("router", "192.168.0.1", []string{"network"})
	assert.Nil(err)

	client, err := DefaultAPI.GetClientsClient("router")
	assert.Nil(err)
	assert.Equal(client.Name, "router")
	assert.Equal(client.Address, "192.168.0.1")
	assert.Equal(client.Subscriptions, []string{"network"})
	assert.NotEqual(client.Timestamp, 0)
	assert.Equal(client.Version, "0.22.0")
}

func TestDeleteClientsClient(t *testing.T) {
	assert := assert.New(t)

	err := DefaultAPI.DeleteClientsClient("router")
	assert.Nil(err)

	time.Sleep(10 * time.Millisecond)

	_, err = DefaultAPI.GetClientsClient("router")
	assert.Equal(err.Error(), "sensu: Not Found")

	err = DefaultAPI.DeleteClientsClient("router")
	assert.Equal(err.Error(), "sensu: Not Found")
}
