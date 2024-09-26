// Package port provides the repository and service interfaces that core services adhere to.
package port

import (
	"context"
	"github.com/dawsonalex/ms-macrod/core/entity"
	"github.com/google/uuid"
)

type FoodRepository interface {
	CreateFood(ctx context.Context, food entity.FoodListing) error
	GetFood(ctx context.Context, id uuid.UUID) (entity.FoodListing, error)
	GetAllFood(ctx context.Context, id ...uuid.UUID) ([]entity.FoodListing, error)
}

type FoodService interface {
	CreateFood(ctx context.Context, food entity.FoodListing) error
	GetFood(ctx context.Context, id uuid.UUID) (entity.FoodListing, error)
	Search(ctx context.Context, query string) ([]entity.FoodListing, error)
}
