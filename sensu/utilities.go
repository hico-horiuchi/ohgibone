package sensu

import "net/http"

func statusCodeToString(status int) string {
	var str string

	switch status {
	case http.StatusOK:
		str = "OK"
	case http.StatusCreated:
		str = "Created"
	case http.StatusAccepted:
		str = "Accepted"
	case http.StatusNoContent:
		str = "No Content"
	case http.StatusBadRequest:
		str = "Bad Request"
	case http.StatusNotFound:
		str = "Not Found"
	case http.StatusInternalServerError:
		str = "Internal Server Error"
	case http.StatusServiceUnavailable:
		str = "Service Unavailable"
	}

	return str
}
