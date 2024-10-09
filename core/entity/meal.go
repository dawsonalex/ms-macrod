package entity

import (
	"github.com/google/uuid"
	"slices"
)

type Meal struct {
	id      uuid.UUID
	Name    string
	entries []FoodEntry
}

// Calories return the total calories for a meal
func (m *Meal) Calories() int {
	totalCalories := 0
	for _, entry := range m.entries {
		totalCalories += entry.Calories()
	}

	return totalCalories
}

// Nutrition returns the total Macros and calories for a meal in one call.
func (m *Meal) Nutrition() (Macros, int) {
	macros := Macros{}
	totalCalories := 0
	for _, entry := range m.entries {
		macros.Carbs += entry.SelectedServing().Carbs
		macros.Fats += entry.SelectedServing().Fats
		macros.Proteins += entry.SelectedServing().Proteins
		totalCalories += entry.Calories()
	}

	return macros, totalCalories
}

func (m *Meal) AddFood(food ...FoodEntry) {
	for _, foodItem := range food {
		m.entries = append(m.entries, foodItem)
	}
}

func (m *Meal) GetFood() []FoodEntry {
	return slices.Clone(m.entries)
}

func (m *Meal) ID() uuid.UUID {
	return m.id
}
