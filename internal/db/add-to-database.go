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

func DeleteUser(conn *sql.DB, telegramID int64) error {
	query := `DELETE FROM users WHERE telegram_id = ?`

	res, err := conn.Exec(query, telegramID)
	if err != nil {
		log.Printf("Ошибка при удалении пользователя %d: %v", telegramID, err)
		return err
	}

	// Опционально: проверяем, удалили ли мы кого-то
	rows, _ := res.RowsAffected()
	if rows == 0 {
		log.Printf("Пользователь с ID %d не найден в базе для удаления", telegramID)
	} else {
		log.Printf("Пользователь %d успешно удален", telegramID)
	}

	return nil
}
