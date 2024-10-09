package entity

// TODO: Refactor this to match the new entity structure
//func TestMeal(t *testing.T) {
//	pastaListing := NewFoodListing("pasta")
//	pastaListing.AddServing(Serving{
//		id:     uuid.New(),
//		size:   "100g",
//		Macros: NewMacros(35, 0, 5),
//	})
//
//	chickenListing := NewFoodListing("chicken")
//	chickenListing.AddServing(Serving{
//		id:     uuid.New(),
//		size:   "100g",
//		Macros: NewMacros(0, 4, 40),
//	})
//
//	meal := Meal{
//		id:   uuid.New(),
//		Name: "Lunch",
//	}
//
//	meal.AddFood(FoodEntry{
//		FoodListing:       pastaListing,
//		id:                uuid.New(),
//		SelectedServingId: pastaListing.Servings()[0].id,
//		Quantity:          1,
//	})
//
//	meal.AddFood(FoodEntry{
//		FoodListing:       chickenListing,
//		id:                uuid.New(),
//		SelectedServingId: chickenListing.Servings()[0].id,
//		Quantity:          1,
//	})
//
//	// Just hard coding the calorie count for this test.
//	expectedCals := 4*CaloriesPer1gFat + 40*CaloriesPer1gProtein + 35*CaloriesPer1gCarbohydrate + 5*CaloriesPer1gProtein
//	if _, cals := meal.Nutrition(); cals != expectedCals {
//		t.Errorf("Expected Nutrition calories to be %d, got %d", expectedCals, cals)
//	}
//}
