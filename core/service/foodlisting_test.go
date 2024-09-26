package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/dawsonalex/ms-macrod/adapter/storage/inmemory"
	"github.com/dawsonalex/ms-macrod/core/entity"
	"github.com/google/uuid"
	"testing"
)

func TestFoodListing_Search(t *testing.T) {
	repo := inmemory.NewRepository()

	index, err := bleve.NewMemOnly(bleve.NewIndexMapping())
	if err != nil {
		t.Fatal(err)
	}

	// Creating the FoodListing here rather than calling
	// NewFoodListing so we can use an in-memory index.
	foodListingService := FoodListing{
		repo:  repo,
		index: index,
	}

	err = foodListingService.CreateFood(context.Background(), entity.FoodListing{
		Id:   uuid.New(),
		Name: "Apples",
	})

	matches, err := foodListingService.Search(context.Background(), "apples")
	if err != nil {
		var batchErr BatchErr
		if errors.As(err, &batchErr) {
			t.Fatalf("%s: %v", batchErr.error, batchErr.ids)
		}
	}
	if len(matches) < 1 {
		t.Error("expected at least one match, but didn't get any")
	}

	jsonMatches, err := json.MarshalIndent(matches, "", " ")
	fmt.Printf("got matches: %s", string(jsonMatches))
}
