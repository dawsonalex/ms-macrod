package httpserver

import (
	"github.com/dawsonalex/ms-macrod/config"
	"github.com/dawsonalex/ms-macrod/core/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func New(logger *log.Logger, conf config.C, foodListingService *service.FoodListing) http.Handler {
	mux := http.NewServeMux()
	addRoutes(
		mux,
		conf,
		logger,
		foodListingService,
	)
	var handler http.Handler = mux
	return handler
}
