package sensu

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
)

var testAPI *API = &API{
	Host: "localhost",
	Port: 80,
}

func testServerAndAPI(status int, body string) (*httptest.Server, *API) {
	server := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(status)
				w.Write([]byte(body))
			},
		),
	)

	url := strings.Split(server.URL, ":")
	port, _ := strconv.ParseInt(url[2], 10, 0)

	return server, &API{
		Host: strings.TrimLeft(url[1], "/"),
		Port: int(port),
	}
}
