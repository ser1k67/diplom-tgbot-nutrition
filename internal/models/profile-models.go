package models

type AllInformation struct {
	ID       int64  // Telegram ID (из c.Sender().ID)
	State    string // Состояние FSM
	Language string // Код языка (ru/kz)

	Gender   string // Пол
	Age      string // Возраст (удобнее сразу в int)
	Weight   string // Вес
	Activity string // Тип активности (ключ из словаря)
	Goal     string // Цель (ключ из словаря)

	// Результаты расчетов (то, что летит в REAL поля базы)
	BMI        float64 // ИМТ
	BMR        float64 // Базовый метаболизм
	TDEE       float64 // Общие расходы
	TargetKcal float64 // Целевые калории

	// Макронутриенты (БЖУ)
	Proteins float64
	Fats     float64
	Carbs    float64
}
