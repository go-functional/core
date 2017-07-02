package composition

import (
	"net/http"
)

// ComposableHandlerResponse is a response to be returned by a ComposableHandlers
type ComposableHandlerResponse struct {
	Code int
	Body []byte
}

// ComposableHandler is a handler for HTTP requests. It is intended to operate on a given
// request and an in-progress response
type ComposableHandler func(
	r *http.Request,
	cur ComposableHandlerResponse,
) *ComposableHandlerResponse

// HTTPOKHandler returns a handler that simply sets the response to a 200 OK, regardless of
// the request
func HTTPOKHandler(r *http.Request, cur ComposableHandlerResponse) *ComposableHandlerResponse {
	cur.Code = http.StatusOK
	return &cur
}

// StaticBodyHandler returns a handler that sets the body to body, regardless of the request
func StaticBodyHandler(body []byte) ComposableHandler {
	return func(r *http.Request, cur ComposableHandlerResponse) *ComposableHandlerResponse {
		cur.Body = body
		return &cur
	}
}
