package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/dawsonalex/ms-macrod/core/entity"
	"github.com/dawsonalex/ms-macrod/core/port"
	"github.com/google/uuid"
)

const foodListingIndexPath = "macrod.foodListing"

type BatchErr struct {
	error
	ids []string
}

type FoodListing struct {
	repo  port.FoodRepository
	index bleve.Index
}

func NewFoodListing(repo port.FoodRepository) (*FoodListing, error) {
	index, err := createOrOpeBleveIndex()
	if err != nil {
		return nil, err
	}

	return &FoodListing{
		repo:  repo,
		index: index,
	}, nil
}

func createOrOpeBleveIndex() (bleve.Index, error) {
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(foodListingIndexPath, mapping)

	// If the index exists we can just open it.
	if err != nil && errors.Is(err, bleve.ErrorIndexPathExists) {

		// If bleve.Open throws an error we have some other issue, return it.
		index, err = bleve.Open(foodListingIndexPath)
		if err != nil {
			return nil, err
		}
	}

	return index, nil
}

func newFoodListing(repo port.FoodRepository, index bleve.Index) (*FoodListing, error) {
	return &FoodListing{
		repo:  repo,
		index: index,
	}, nil
}

func (f *FoodListing) CreateFood(ctx context.Context, food entity.FoodListing) error {
	err := f.repo.CreateFood(ctx, food)
	if err != nil {
		return err
	}

	return f.index.Index(food.ID().String(), food.Name)
}

func (f *FoodListing) GetFood(ctx context.Context, id uuid.UUID) (entity.FoodListing, error) {
	return f.repo.GetFood(ctx, id)
}

func (f *FoodListing) Search(ctx context.Context, query string) ([]entity.FoodListing, error) {
	bleveQuery := bleve.NewQueryStringQuery(query)
	searchRequest := bleve.NewSearchRequest(bleveQuery)
	result, err := f.index.Search(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("query (%s) search err: %s", query, err.Error())
	}

	foods := []entity.FoodListing{}
	batchErr := BatchErr{
		error: fmt.Errorf("error getting some matched foods"),
		ids:   []string{},
	}
	for _, hit := range result.Hits {
		id, err := uuid.Parse(hit.ID)
		if err != nil {

		}
		food, err := f.repo.GetFood(ctx, id)
		if err != nil {
			batchErr.ids = append(batchErr.ids, id.String())
			continue
		}

		foods = append(foods, food)
	}

	if len(batchErr.ids) > 0 {
		return foods, batchErr
	}

	return foods, nil
}
