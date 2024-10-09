package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/dawsonalex/ms-macrod/core/entity"
	"github.com/dawsonalex/ms-macrod/core/port"
	"github.com/google/uuid"
	"log/slog"
)

type FoodListing struct {
	repo   port.FoodRepository
	index  bleve.Index
	Logger *slog.Logger
}

func NewFoodListing(logger *slog.Logger, repo port.FoodRepository) (*FoodListing, error) {
	index, err := createOrOpeBleveIndex()
	if err != nil {
		return nil, err
	}

	serviceLogger := logger.With("service", "foodlisting")
	return &FoodListing{
		repo:   repo,
		index:  index,
		Logger: serviceLogger,
	}, nil
}

func createOrOpeBleveIndex() (bleve.Index, error) {
	mapping := bleve.NewIndexMapping()
	return bleve.NewMemOnly(mapping)
}

func (f *FoodListing) CreateFood(ctx context.Context, food entity.FoodListing) error {
	if food.Id == uuid.Nil {
		food.Id = uuid.New()
	}

	err := f.repo.CreateFood(ctx, food)
	if err != nil {
		return err
	}

	return f.index.Index(food.Id.String(), food.Name)
}

func (f *FoodListing) GetFood(ctx context.Context, id uuid.UUID) (entity.FoodListing, error) {
	return f.repo.GetFood(ctx, id)
}

func (f *FoodListing) Search(ctx context.Context, query string) ([]entity.FoodListing, error) {
	bleveQuery := bleve.NewQueryStringQuery(fmt.Sprintf("%s*", query))
	searchRequest := bleve.NewSearchRequest(bleveQuery)
	result, err := f.index.Search(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("query (%s) search err: %s", query, err.Error())
	}

	foods := []entity.FoodListing{}
	missingIds := make([]string, 0)
	for _, hit := range result.Hits {
		id, err := uuid.Parse(hit.ID)
		if err != nil {
			// TODO: something else should happen here.
			f.Logger.Error("error parsing indexed foodlisting item", err)
		}

		food, err := f.repo.GetFood(ctx, id)
		errNoExist := &port.ErrEntityNoExist{}
		if errors.As(err, errNoExist) {
			missingIds = append(missingIds, id.String())
			continue
		}

		foods = append(foods, food)
	}

	if len(missingIds) > 0 {
		go func() {
			f.Logger.Info("purging %d ids from index", len(missingIds))

			for _, id := range missingIds {
				_ = f.index.Delete(id)
			}
		}()
	}

	return foods, nil
}
