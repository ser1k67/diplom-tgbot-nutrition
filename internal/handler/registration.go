package handler

import (
	"database/sql"
	"diplomkabot/internal/models"
	"fmt"

	tb "gopkg.in/telebot.v3"
)

func KeyboardButtons(arr []string) []tb.ReplyButton {
	var inlineElements []tb.ReplyButton
	var inline tb.ReplyButton
	for _, v := range arr {
		inline.Text = v
		inlineElements = append(inlineElements, inline)
	}
	return inlineElements
}

var machine = make(map[int64]models.AllInformation)

func Registration(conn *sql.DB, c tb.Context) error {
	var exists bool
	err := conn.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE telegram_id = ?)", c.Sender().ID).Scan(&exists)

	// Если произошла ошибка БД (кроме отсутствия строк), лучше её залогировать
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("Ошибка проверки регистрации:", err)
	}
	// 2. Если запись найдена, блокируем регистрацию
	if exists {
		return c.Send("✅ <b>Вы уже зарегистрированы!</b>\nИспользуйте /profile для просмотра данных.", tb.RemoveKeyboard, tb.ModeHTML)
	}
	machine[c.Sender().ID] = models.AllInformation{
		TelegramID: c.Sender().ID,
		State:      "WAIT_LANGUAGE",
	}

	buttons := KeyboardButtons([]string{"🇰🇿 Қазақ тілі", "🇷🇺 Русский язык"})
	var markups tb.ReplyMarkup
	markups.ReplyKeyboard = append(markups.ReplyKeyboard, buttons)
	markups.ResizeKeyboard = true

	err = c.Send("⚙️ <b>1. Выберите язык / Тілді таңдаңыз</b> ", &markups, tb.ModeHTML)
	if err != nil {
		return fmt.Errorf("не удалось отправить сообщение пользователю: %w", err)
	}

	return nil
}
func MachineState(conn *sql.DB, c tb.Context) error {
	user, ok := machine[c.Sender().ID]
	if !ok {
		return nil
	}
	switch user.State {
	case "WAIT_LANGUAGE":
		return HandleLanguage(c, user)
	case "WAIT_GENDER":
		return HandleGender(c, user)
	case "WAIT_AGE":
		return HandleAge(c, user)
	case "WAIT_WEIGHT":
		return HandleWeight(c, user)
	case "WAIT_HEIGHT":
		return HandleHeight(c, user)
	case "WAIT_ACTIVITY":
		return HandleActivity(c, user)
	case "WAIT_GOALS":
		return HandleGoals(conn, c, user)
	}
	return nil
}
