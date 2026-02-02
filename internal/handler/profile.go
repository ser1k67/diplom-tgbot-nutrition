package handler

import (
	"database/sql"
	"diplomkabot/internal/models"
	"fmt"

	tb "gopkg.in/telebot.v3"
)

func Profile(conn *sql.DB, c tb.Context) error {
	var user models.AllInformation

	// Берем ID того, кто отправил запрос
	senderID := c.Sender().ID

	// Важно: Порядок полей в SELECT должен СТРОГО совпадать с порядком в .Scan
	query := `
		SELECT 
			telegram_id, language, gender, age, weight, 
			height, activity, goal, bmi, bmr, 
			tdee, target_kcal 
		FROM users 
		WHERE telegram_id = ?`

	err := conn.QueryRow(query, senderID).Scan(
		&user.TelegramID,
		&user.Language,
		&user.Gender,
		&user.Age,
		&user.Weight,
		&user.Height,
		&user.Activity,
		&user.Goal,
		&user.BMI,
		&user.BMR,
		&user.TDEE,
		&user.TargetKcal,
	)

	// 2. Если пользователя нет в базе — прерываем выполнение
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Send("❌ <b>Анкета не найдена.</b>\nПожалуйста, сначала пройдите регистрацию командой /start", tb.ModeHTML)
		}
		// Логируем другие ошибки БД (например, проблемы с подключением)
		fmt.Println("Database error:", err)
		return err
	}

	// 3. Формируем сообщение в виде списка с точками
	// <code> делает текст моноширинным (удобно для ID)
	information := fmt.Sprintf(
		"👤 <b>Ваш профиль:</b>\n\n"+
			"• <b>ID:</b> <code>%d</code>\n"+
			"• <b>Язык:</b> %s\n"+
			"• <b>Пол:</b> %s\n"+
			"• <b>Возраст:</b> %s лет\n"+
			"• <b>Вес:</b> %s кг\n"+
			"• <b>Рост:</b> %s см\n"+
			"• <b>Активность:</b> %s\n"+
			"• <b>Цель:</b> %s\n\n"+
			"📊 <b>Результаты расчетов:</b>\n\n"+
			"• <b>ИМТ (BMI):</b> %.1f\n"+
			"• <b>Базовый метаболизм:</b> %.0f ккал\n"+
			"• <b>Расход с учетом нагрузок:</b> %.0f ккал\n\n"+
			"🔥 <b>Ваша норма для цели:</b> <u>%.0f ккал/день</u>",
		user.TelegramID,
		user.Language,
		user.Gender,
		user.Age,
		user.Weight,
		user.Height,
		user.Activity,
		user.Goal,
		user.BMI,
		user.BMR,
		user.TDEE,
		user.TargetKcal,
	)
	markup := MainMenu(c)
	markup.ResizeKeyboard = true
	return c.Send(information, &markup, tb.ModeHTML)
}
