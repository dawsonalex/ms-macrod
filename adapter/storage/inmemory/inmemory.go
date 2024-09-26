package inmemory

import (
	"context"
	"fmt"
	"github.com/dawsonalex/ms-macrod/core/entity"
	"github.com/google/uuid"
	"sync"
)

// TODO: context needs using here.

// inMemoryImpl implements all repository ports.
type inMemoryImpl struct {
	sync.RWMutex
	foodListings map[uuid.UUID]entity.FoodListing
}

func NewRepository() *inMemoryImpl {
	return &inMemoryImpl{
		foodListings: make(map[uuid.UUID]entity.FoodListing),
	}
}

func (i *inMemoryImpl) CreateFood(ctx context.Context, food entity.FoodListing) (err error) {
	i.Lock()
	defer i.Unlock()
	i.foodListings[food.ID()] = food
	return nil
}

func (i *inMemoryImpl) GetFood(ctx context.Context, id uuid.UUID) (entity.FoodListing, error) {
	i.RLock()
	defer i.RUnlock()

	food, ok := i.foodListings[id]
	if !ok {
		return entity.FoodListing{}, fmt.Errorf("food %s not found", id.String())
	}
	return food, nil
}

func (i *inMemoryImpl) GetAllFood(ctx context.Context, ids ...uuid.UUID) ([]entity.FoodListing, error) {
	i.RLock()
	defer i.RUnlock()

	foods := make([]entity.FoodListing, 0, len(ids))
	for _, id := range ids {
		foods = append(foods, i.foodListings[id])
	}

	return foods, nil
}
