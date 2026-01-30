package rations

import "diplomkabot/internal/models"

// Меняет грамовку и БЖУ
func Ration(file *models.NutritionPlan, calories int) {
	allMeals := [][]models.Ingredient{
		file.Meals.Breakfast,
		file.Meals.Snack1,
		file.Meals.Lunch,
		file.Meals.Snack2,
		file.Meals.Dinner,
	}

	for i, value := range allMeals {
		for i2, _ := range value {
			allMeals[i][i2].WeightG = allMeals[i][i2].WeightG * calories / 2000
		}
	}

	file.BaseMacros.CarbsG = file.BaseMacros.CarbsG * calories / 2000
	file.BaseMacros.FatG = file.BaseMacros.FatG * calories / 2000
	file.BaseMacros.ProteinG = file.BaseMacros.ProteinG * calories / 2000
}
