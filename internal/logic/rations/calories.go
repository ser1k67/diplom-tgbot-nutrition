package rations

// Это калории в состоянии покоя.
func BMR(weight, height float64, age int, gender string) float64 {
	if gender == "male" {
		return 10*weight + 6.25*height - 5*float64(age) + 5
	}
	return 10*weight + 6.25*height - 5*float64(age) - 161
}

// коэффициент активности
func ActivityCoefficient(activity string) float64 {
	switch activity {
	case "легкая":
		return 1.2
	case "средняя":
		return 1.55
	case "сильная":
		return 1.75
	default:
		return 0
	}
}

// коэффициент для похудения, поддержания, набора
func TargetCoefficient(target string) float64 {
	switch target {
	case "похудение":
		return 0.8
	case "поддержание":
		return 1
	case "набор":
		return 1.15
	default:
		return 0
	}
}

// Это калории для поддержания текущего веса.
// tdee := BMR * Activity

// калорий для конечной цели
// calories := tdee * TargetCoefficient
