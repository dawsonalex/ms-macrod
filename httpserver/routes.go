package httpserver

import (
	"context"
	"encoding/json"
	"github.com/dawsonalex/ms-macrod/build"
	"github.com/dawsonalex/ms-macrod/config"
	"github.com/dawsonalex/ms-macrod/core/entity"
	"github.com/dawsonalex/ms-macrod/core/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func addRoutes(
	mux *http.ServeMux,
	conf config.C,
	logger *log.Logger,
	foodListingService *service.FoodListing,
) {
	mux.Handle("GET /version", handleVersionGet())
	mux.Handle("GET /config", handleConfigGet(conf))

	mux.Handle("POST /foodlisting", handleFoodListingCreate(foodListingService))
	mux.Handle("GET /foodlisting", handleFoodListingSearch(logger, foodListingService))
}

func handleVersionGet() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			e := json.NewEncoder(w)
			err := e.Encode(build.Info())
			if err != nil {
				w.WriteHeader(500)
				_, err = w.Write([]byte(err.Error()))
				if err != nil {
					panic(err)
				}
			}
		},
	)
}

func handleConfigGet(conf config.C) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			e := json.NewEncoder(w)
			err := e.Encode(conf)
			if err != nil {
				w.WriteHeader(500)
				_, err = w.Write([]byte(err.Error()))
				if err != nil {
					panic(err)
				}
			}
		},
	)
}

func handleFoodListingCreate(service *service.FoodListing) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			decoder := json.NewDecoder(r.Body)
			var foodListing entity.FoodListing
			err := decoder.Decode(&foodListing)
			if err != nil {
				panic(err)
			}

			err = service.CreateFood(context.Background(), foodListing)
			if err != nil {
				panic(err)
			}
		},
	)
}

func handleFoodListingSearch(logger *log.Logger, s *service.FoodListing) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			res := []entity.FoodListing{}
			if query := r.URL.Query().Get("q"); len(query) > 0 {
				matches, err := s.Search(context.Background(), query)
				if err != nil {
					w.WriteHeader(500)
					_, err = w.Write([]byte(err.Error()))
					if err != nil {
						panic(err)
					}
				}
				res = matches
			}

			e := json.NewEncoder(w)
			err := e.Encode(res)
			if err != nil {
				w.WriteHeader(500)
				_, err = w.Write([]byte(err.Error()))
				if err != nil {
					panic(err)
				}
			}
		},
	)
}
