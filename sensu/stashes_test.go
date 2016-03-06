package sensu

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type exampleStruct struct {
	StashStruct
	Content contentStruct `json:"content"`
}

type contentStruct struct {
	Message string `json:"message"`
}

func TestGetStashes(t *testing.T) {
	assert := assert.New(t)
	stashes := []exampleStruct{}

	err := DefaultAPI.GetStashes(&stashes, -1, -1)
	assert.Nil(err)

	err = testAPI.GetStashes(&stashes, -1, -1)
	assert.Contains(err.Error(), "getsockopt: connection refused")

	server, api := testServerAndAPI(http.StatusInternalServerError, "")
	defer server.Close()
	err = api.GetStashes(&stashes, -1, -1)
	assert.Equal(err.Error(), "sensu: Internal Server Error")

	server, api = testServerAndAPI(http.StatusOK, "")
	defer server.Close()
	err = api.GetStashes(&stashes, -1, -1)
	assert.Equal(err.Error(), "unexpected end of JSON input")
}

func TestPostStashes(t *testing.T) {
	assert := assert.New(t)
	stash := exampleStruct{
		StashStruct: StashStruct{
			Expire: -1,
			Path:   "example/stash",
		},
		Content: contentStruct{
			Message: "example",
		},
	}

	err := DefaultAPI.PostStashes(&stash)
	assert.Nil(err)

	err = testAPI.PostStashes(&stash)
	assert.Contains(err.Error(), "getsockopt: connection refused")

	server, api := testServerAndAPI(http.StatusInternalServerError, "")
	defer server.Close()
	err = api.PostStashes(&stash)
	assert.Equal(err.Error(), "sensu: Internal Server Error")
}

func TestPostStashesPath(t *testing.T) {
	assert := assert.New(t)
	content := contentStruct{
		Message: "example",
	}

	err := DefaultAPI.PostStashesPath("example/stash", &content)
	assert.Nil(err)

	err = testAPI.PostStashesPath("example/stash", &content)
	assert.Contains(err.Error(), "getsockopt: connection refused")

	server, api := testServerAndAPI(http.StatusInternalServerError, "")
	defer server.Close()
	err = api.PostStashesPath("example/stash", &content)
	assert.Equal(err.Error(), "sensu: Internal Server Error")
}

func TestGetStashesPath(t *testing.T) {
	assert := assert.New(t)
	content := contentStruct{}

	err := DefaultAPI.GetStashesPath("example/stash", &content)
	assert.Nil(err)
	assert.Equal(content.Message, "example")

	err = DefaultAPI.GetStashesPath("test/stash", &content)
	assert.Equal(err.Error(), "sensu: Not Found")

	err = testAPI.GetStashesPath("example/stash", &content)
	assert.Contains(err.Error(), "getsockopt: connection refused")

	server, api := testServerAndAPI(http.StatusInternalServerError, "")
	defer server.Close()
	err = api.GetStashesPath("example/stash", &content)
	assert.Equal(err.Error(), "sensu: Internal Server Error")

	server, api = testServerAndAPI(http.StatusOK, "")
	defer server.Close()
	err = api.GetStashesPath("example/stash", &content)
	assert.Equal(err.Error(), "unexpected end of JSON input")
}

func TestDeleteStashesPath(t *testing.T) {
	assert := assert.New(t)

	err := DefaultAPI.DeleteStashesPath("example/stash")
	assert.Nil(err)

	err = DefaultAPI.DeleteStashesPath("test/stash")
	assert.Equal(err.Error(), "sensu: Not Found")

	err = testAPI.DeleteStashesPath("example/stash")
	assert.Contains(err.Error(), "getsockopt: connection refused")

	server, api := testServerAndAPI(http.StatusInternalServerError, "")
	defer server.Close()
	err = api.DeleteStashesPath("example/stash")
	assert.Equal(err.Error(), "sensu: Internal Server Error")
}
