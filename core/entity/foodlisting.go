package entity

import (
	"github.com/google/uuid"
)

type FoodListing struct {
	Id       uuid.UUID             `json:"id"`
	Name     string                `json:"name"`
	Servings map[uuid.UUID]Serving `json:"servings"`
}

func NewFoodListing(name string) FoodListing {
	return FoodListing{
		Name:     name,
		Servings: map[uuid.UUID]Serving{},
	}
}

func (f *FoodListing) AddServing(serving Serving) {
	if f.Servings == nil {
		f.Servings = make(map[uuid.UUID]Serving)
	}

	f.Servings[serving.id] = serving
}

func (f *FoodListing) ID() uuid.UUID {
	return f.Id
}
