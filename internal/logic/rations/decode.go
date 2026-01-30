package rations

import (
	"diplomkabot/internal/models"
	"encoding/json"
	"os"
)

// функция для ахуеть декодирования прикинь в нем рационы хранятся с неправильной грамовкой
func Decode(ration string) (models.NutritionPlan, error) {
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
		return models.NutritionPlan{}, err
	}
	defer file.Close()

	// дальше не ебу братан за эти ебаные структуры
	var conf models.NutritionPlan
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		return models.NutritionPlan{}, err
	}

	return conf, err
}
