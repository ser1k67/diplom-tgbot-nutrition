package rations

import (
	"encoding/json"
	"os"
)

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

// функция для ахуеть декодирования прикинь в нем рационы хранятся с неправильной грамовкой
func Decode(ration string) (NutritionPlan, error) {
	var file *os.File
	var err error

	// форматирование для выбора пути для чтения
	if ration == "стандартная" {
		file, err = os.Open("./standart.json")
	} else if ration == "премиумная" {
		file, err = os.Open("./premium.json")
	}

	// чтение .json
	if err != nil {
		return NutritionPlan{}, err
	}
	defer file.Close()

	// дальше не ебу братан за эти ебаные структуры
	var conf NutritionPlan
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		return NutritionPlan{}, err
	}

	return conf, err
}
