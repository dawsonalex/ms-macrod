package httpserver

import (
	"github.com/dawsonalex/ms-macrod/config"
	"net/http"
)

func New(conf config.C) http.Handler {
	mux := http.NewServeMux()
	addRoutes(
		mux,
		conf,
	)
	var handler http.Handler = mux
	return handler
}
