package handler

import (
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

func Registration(c tb.Context) error {
	machine[c.Sender().ID] = models.AllInformation{State: "WAIT_LANGUAGE"}
	buttons := KeyboardButtons([]string{"🇰🇿 Қазақ тілі", "🇷🇺 Русский язык"})
	var markups tb.ReplyMarkup
	markups.ReplyKeyboard = append(markups.ReplyKeyboard, buttons)
	markups.ResizeKeyboard = true
	var err error
	err = c.Send("⚙️ <b>1. Выберите язык / Тілді таңдаңыз</b> ", &markups, tb.ModeHTML)
	if err != nil {
		return fmt.Errorf("Не удалось отправить сообщение пользователю")
	}
	return nil
}

func MachineState(c tb.Context) error {
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
	case "WAIT_ACTIVITY":
		return HandleActivity(c, user)
	case "WAIT_GOALS":
		return HandleGoals(c, user)
	}
	return nil
}
