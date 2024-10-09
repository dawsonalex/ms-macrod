package entity

// Macros represents macro values for a food source.
// They are useful only when used relative to a serving or quantity.
type Macros struct {
	Carbs    float64 `json:"carbs"`
	Fats     float64 `json:"fats"`
	Proteins float64 `json:"proteins"`
}

func (m Macros) Calories() int {
	return int((m.Carbs * CaloriesPer1gCarbohydrate) +
		(m.Fats * CaloriesPer1gFat) +
		(m.Proteins * CaloriesPer1gProtein))
}
