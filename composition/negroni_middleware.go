package composition

import (
	"net/http"
)

// NegroniMiddleware is the middleware definition taken from
// https://github.com/urfave/negroni#handlers
type NegroniMiddleware interface {
	ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}
