package controller

import "errors"

var (
	NoResponseWriter = errors.New("no_response_writter")
	NoRequest        = errors.New("no_request")
)
