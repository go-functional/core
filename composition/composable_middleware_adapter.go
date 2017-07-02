package composition

import (
	"net/http"
)

// HTTPRoute converts a series of ComposableHandlers into a http.HandlerFunc, for use
// in the standard net/http library
func HTTPRoute(handler ComposableHandler, others ...ComposableHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := &ComposableHandlerResponse{}
		resp = handler(r, *resp)
		for _, otherHandler := range others {
			resp = otherHandler(r, *resp)
		}
		w.WriteHeader(resp.Code)
		w.Write(resp.Body)
	}
}
