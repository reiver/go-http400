package http400

import (
	"net/http"
)

func BadRequest(responseWriter http.ResponseWriter, request *http.Request) {
	ServeString(responseWriter, DefaultStatusText)
}
