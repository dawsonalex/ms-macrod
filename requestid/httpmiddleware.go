package requestid

import (
	"github.com/google/uuid"
	"net/http"
)

// RequestContextMiddleware applies context values that should exist the entire lifetime of a request.
func RequestContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AddToRequest(r, uuid.New())
		next.ServeHTTP(w, r)
	})
}
