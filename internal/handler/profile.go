package handler

import (
	"fmt"

	tb "gopkg.in/telebot.v3"
)

func Profile(c tb.Context) error {
	user, ok := machine[c.Sender().ID]
	if !ok {
		return c.Send("Сначала заполните анкету")
	}

	// Формируем красивую строку
	information := fmt.Sprintf(
		"👤 **Ваш профиль:**\n\n"+
			"ID: `%d`\n"+
			"Язык: %s\n"+
			"Пол: %s\n"+
			"Возраст: %s лет\n"+
			"Вес: %s кг\n"+
			"Рост: %s см\n"+
			"Активность: %s\n"+
			"Цель: %s\n\n"+
			"**Результаты расчетов:**\n\n"+
			"ИМТ: %.1f\n"+
			"Базовый метаболизм (BMR): %.0f ккал\n"+
			"Норма для поддержания (TDEE): %.0f ккал\n\n"+
			"**Ваша норма для цели:** %.0f ккал/день",
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

	return c.Send(information, tb.ModeMarkdown)
}
