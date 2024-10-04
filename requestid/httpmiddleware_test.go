package requestid

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestContextMiddleware(t *testing.T) {
	checkContextFunc := func() http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			value, ok := FromContext(r.Context())
			if !ok || (ok && value == "") {
				t.Errorf("missing context value for %s", requestIdKey)
			}
		})
	}

	req := httptest.NewRequest("GET", "/test", nil)

	handler := RequestContextMiddleware(checkContextFunc())
	handler.ServeHTTP(httptest.NewRecorder(), req)
}
