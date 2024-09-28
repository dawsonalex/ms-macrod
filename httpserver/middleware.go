package httpserver

import (
	"github.com/dawsonalex/ms-macrod/config"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func logRequestMiddleware(conf config.C, logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path, query := splitUriQuery(r.RequestURI)

		fields := log.Fields{
			"Method": r.Method,
			"Route":  path,
		}

		if query != "" {
			fields["Query"] = query
		}
		logger.WithFields(fields).Infof("%s %s", r.Method, r.RequestURI)

		next.ServeHTTP(w, r)
	})
}
