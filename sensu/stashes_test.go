package sensu

import (
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

	err := DefaultAPI.PostStashes(stash)
	assert.Nil(err)
}

func TestPostStashesPath(t *testing.T) {
	assert := assert.New(t)
	content := contentStruct{
		Message: "example",
	}

	err := DefaultAPI.PostStashesPath("example/stash", content)
	assert.Nil(err)
}

func TestGetStashesPath(t *testing.T) {
	assert := assert.New(t)
	content := contentStruct{}

	err := DefaultAPI.GetStashesPath("example/stash", &content)
	assert.Nil(err)
	assert.Equal(content.Message, "example")

	err = DefaultAPI.GetStashesPath("test/stash", &content)
	assert.Equal(err.Error(), "sensu: Not Found")
}

func TestDeleteStashesPath(t *testing.T) {
	assert := assert.New(t)

	err := DefaultAPI.DeleteStashesPath("example/stash")
	assert.Nil(err)

	err = DefaultAPI.DeleteStashesPath("test/stash")
	assert.Equal(err.Error(), "sensu: Not Found")
}
