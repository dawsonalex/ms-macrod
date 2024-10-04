package httpserver

import (
	"github.com/dawsonalex/ms-macrod/config"
	"github.com/dawsonalex/ms-macrod/core/service"
	"github.com/dawsonalex/ms-macrod/requestid"
	"github.com/samber/slog-http"
	"log/slog"
	"net/http"
)

func New(logger *slog.Logger, conf config.C, foodListingService *service.FoodListing) http.Handler {
	mux := http.NewServeMux()
	addRoutes(
		mux,
		conf,
		logger,
		foodListingService,
	)
	var handler http.Handler = mux
	handler = sloghttp.Recovery(handler)
	handler = requestid.RequestContextMiddleware(handler)
	handler = sloghttp.NewWithConfig(logger, setSlogHttpConf(conf))(handler)
	return handler
}

// setSlogHttpConf sets the required globals for sloghttp and returns a built config
// from conf
func setSlogHttpConf(conf config.C) sloghttp.Config {
	sloghttp.RequestIDHeaderKey = conf.Log.HttpCorrelationIDHeaderKey
	sloghttp.RequestIDKey = conf.Log.HttpCorrelationIDKey

	return sloghttp.Config{
		WithUserAgent:      conf.Log.IncludeUserAgent,
		WithRequestBody:    conf.Log.IncludeRequestBody,
		WithRequestHeader:  conf.Log.IncludeRequestHeader,
		WithResponseBody:   conf.Log.IncludeResponseBody,
		WithResponseHeader: conf.Log.IncludeResponseHeader,
		WithSpanID:         true,
		WithTraceID:        true,
	}
}
