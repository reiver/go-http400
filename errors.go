package http400

import (
	"github.com/reiver/go-erorr"
)

const (
	ErrNilHTTPResponseWriter = erorr.Error("http400: nil http-response-writer")
)

const (
	errNilHTTPResponseWriterHeader = erorr.Error("http400: nil http-response-writer header")
)
