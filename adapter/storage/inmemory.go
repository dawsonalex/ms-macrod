package storage

import (
	"context"
	"github.com/dawsonalex/ms-macrod/core/entity"
	"github.com/dawsonalex/ms-macrod/core/port"
	"github.com/google/uuid"
	"sync"
)

// port interface implementation check
var _ port.FoodRepository = &InMemory{}

// InMemory implements all repository ports.
//
// InMemory is safe for concurrent use.
type InMemory struct {
	sync.RWMutex
	foodListings map[uuid.UUID]entity.FoodListing
}

func NewInMemory() *InMemory {
	return &InMemory{
		foodListings: make(map[uuid.UUID]entity.FoodListing),
	}
}

func (i *InMemory) CreateFood(ctx context.Context, food entity.FoodListing) (err error) {
	i.Lock()
	defer i.Unlock()
	i.foodListings[food.Id] = food
	return nil
}

func (i *InMemory) GetFood(ctx context.Context, id uuid.UUID) (entity.FoodListing, error) {
	i.RLock()
	defer i.RUnlock()

	food, ok := i.foodListings[id]
	if !ok {
		return entity.FoodListing{}, port.ErrEntityNoExist{ID: id.String()}
	}
	return food, nil
}

func (i *InMemory) GetAllFood(ctx context.Context, ids ...uuid.UUID) ([]entity.FoodListing, error) {
	i.RLock()
	defer i.RUnlock()

	foods := make([]entity.FoodListing, 0, len(ids))
	for _, id := range ids {
		foods = append(foods, i.foodListings[id])
	}

	return foods, nil
}
