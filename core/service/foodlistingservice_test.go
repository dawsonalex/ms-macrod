package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dawsonalex/ms-macrod/adapter/storage"
	"github.com/dawsonalex/ms-macrod/core/entity"
	"log/slog"
	"testing"
)

func setupService(t *testing.T) *FoodListing {
	repo := storage.NewInMemory()

	// setup and return a FoodListing service with discard logger.
	foodListingService, err := NewFoodListing(slog.Default(), repo)
	if err != nil {
		t.Fatal(fmt.Errorf("[setupService()] error creating food listing service: %w", err))
	}
	return foodListingService
}

func TestFoodListing_Search(t *testing.T) {
	foodListingService := setupService(t)

	err := foodListingService.CreateFood(context.Background(), entity.FoodListing{
		Name: "Apples",
	})

	matches, err := foodListingService.Search(context.Background(), "apples")
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) < 1 {
		t.Error("expected at least one match, but didn't get any")
	}

	jsonMatches, err := json.MarshalIndent(matches, "", " ")
	fmt.Printf("got matches: %s", string(jsonMatches))
}
