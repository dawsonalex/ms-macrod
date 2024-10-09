package entity

import (
	"github.com/google/uuid"
)

type FoodListing struct {
	Id       uuid.UUID         `json:"id,omitempty"`
	Name     string            `json:"name"`
	Servings map[string]Macros `json:"servings"`
}

func NewFoodListing(name string) FoodListing {
	return FoodListing{
		Name:     name,
		Servings: map[string]Macros{},
	}
}
