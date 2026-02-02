package rations

import (
	"diplomkabot/internal/models"
	"fmt"
	"strconv"
)

// CalculateAll выполняет цепочку всех расчетов
func CalculateAll(info *models.AllInformation) error {
	if err := CalculateBMR(info); err != nil {
		return err
	}
	CalculateBMI(info)
	CalculateTDEE(info)
	CalculatePFC(info)
	return nil
}

func CalculateBMR(info *models.AllInformation) error {
	wF, err := strconv.ParseFloat(info.Weight, 64)
	hF, err := strconv.ParseFloat(info.Height, 64)
	aF, err := strconv.ParseFloat(info.Age, 64)
	if err != nil {
		return err
	}

	if info.Gender == "👨 Мужской" || info.Gender == "👨 Ер адам" {
		info.BMR = 10*wF + 6.25*hF - 5*aF + 5
	} else {
		info.BMR = 10*wF + 6.25*hF - 5*aF - 161
	}
	return nil
}

func CalculateBMI(info *models.AllInformation) {
	wF, _ := strconv.ParseFloat(info.Weight, 64)
	hF, _ := strconv.ParseFloat(info.Height, 64)
	hM := hF / 100
	info.BMI = wF / (hM * hM)
}
func CalculatePFC(all *models.AllInformation) {
	all.Proteins = (all.BMR * 0.3) / 4
	fmt.Println(all.Proteins)
	all.Fats = (all.BMR * 0.3) / 9
	all.Carbs = (all.BMR * 0.4) / 4
}

func CalculateTDEE(info *models.AllInformation) {
	var activityCoeff float64
	switch info.Activity {
	case "🟢 Легкие нагрузки", "🟢 Жеңіл жаттығулар":
		activityCoeff = 1.2
	case "🟡 Умеренные нагрузки", "🟡 Орташа жаттығулар":
		activityCoeff = 1.55
	case "🔴 Сильные нагрузки", "🔴 Ауыр жаттығулар":
		activityCoeff = 1.75
	default:
		activityCoeff = 1.2
	}

	info.TDEE = info.BMR * activityCoeff

	var targetCoeff float64
	switch info.Goal {
	case "📉 Сбросить вес", "📉 Салмақ тастау":
		targetCoeff = 0.85 // дефицит 15%
	case "📈 Набрать массу", "📈 Салмақ қосу":
		targetCoeff = 1.15 // профицит 15%
	default:
		targetCoeff = 1.0
	}
	info.TargetKcal = info.TDEE * targetCoeff
}
