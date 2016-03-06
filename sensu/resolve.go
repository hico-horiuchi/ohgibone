package sensu

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type ResolveStruct struct {
	Client string `json:"client"`
	Check  string `json:"check"`
}

// Resolves an event.
func (api API) PostResolve(client string, check string) error {
	resolve := ResolveStruct{
		Client: client,
		Check:  check,
	}

	body, err := json.Marshal(resolve)
	if err != nil {
		return err
	}
	payload := strings.NewReader(string(body))

	response, err := api.post("/resolve", payload)
	if err != nil {
		return err
	} else if response.StatusCode != http.StatusAccepted {
		return errors.New("sensu: " + statusCodeToString(response.StatusCode))
	}

	return nil
}
