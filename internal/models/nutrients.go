package models

// Продукт внутри приема пищи
type Ingredient struct {
	Name    string `json:"name"`
	WeightG int    `json:"weight_g"`
}

// БЖУ
type PFC struct {
	ProteinG int `json:"protein_g"`
	FatG     int `json:"fat_g"`
	CarbsG   int `json:"carbs_g"`
}

// Группы приемов пищи
type Meals struct {
	Breakfast []Ingredient `json:"breakfast"`
	Snack1    []Ingredient `json:"snack1"`
	Lunch     []Ingredient `json:"lunch"`
	Snack2    []Ingredient `json:"snack2"`
	Dinner    []Ingredient `json:"dinner"`
}

// Главный объект
type NutritionPlan struct {
	BaseCalories int    `json:"baseCalories"`
	DietType     string `json:"dietType"`
	BaseMacros   PFC    `json:"BaseMacros"`
	Meals        Meals  `json:"meals"`
}
