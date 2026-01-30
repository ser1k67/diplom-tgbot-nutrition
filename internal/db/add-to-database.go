package db

import (
	"database/sql"
	"diplomkabot/internal/models"
	"log"
)

// Добавление профиля в базу данных
func AddToDatabase(conn *sql.DB, profile models.AllInformation) error {
	query := `
	INSERT OR REPLACE INTO users (
		telegram_id, language, gender, age, weight, activity, goal, height, 
		bmi, bmr, tdee, target_kcal, proteins, fats, carbs
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := conn.Exec(query,
		profile.TelegramID,
		profile.Language,
		profile.Gender,
		profile.Age,
		profile.Weight,
		profile.Activity,
		profile.Goal,
		profile.Height,
		profile.BMI,
		profile.BMR,
		profile.TDEE,
		profile.TargetKcal,
		profile.Proteins,
		profile.Fats,
		profile.Carbs,
	)

	if err != nil {
		log.Printf("Ошибка вставки в БД: %v", err)
		return err
	}

	return nil
}
