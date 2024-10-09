package entity

const (
	CaloriesPer1gProtein      = 4
	CaloriesPer1gCarbohydrate = 4
	CaloriesPer1gFat          = 9
)

type Serving struct {
	Macros Macros
}

func (s Serving) Calories() int {
	return int((s.Macros.Carbs * CaloriesPer1gCarbohydrate) +
		(s.Macros.Fats * CaloriesPer1gFat) +
		(s.Macros.Proteins * CaloriesPer1gProtein))
}
