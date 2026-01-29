package handler

import (
	"diplomkabot/internal/models"
	"fmt"

	tb "gopkg.in/telebot.v3"
)

// Функция для определения языка
func HandleLanguage(c tb.Context, user models.AllInformation) error {
	if c.Text() == "🇷🇺 Русский язык" {
		user.Language = "ru"
	} else {
		user.Language = "kz"
	}
	user.State = "WAIT_GENDER"

	txt := models.Phrases[user.Language]["ask_gender"]
	b1 := models.Phrases[user.Language]["btn_male"]
	b2 := models.Phrases[user.Language]["btn_female"]

	buttons := KeyboardButtons([]string{b1, b2})
	var markups tb.ReplyMarkup
	markups.ReplyKeyboard = append(markups.ReplyKeyboard, buttons)
	markups.ResizeKeyboard = true

	err := c.Send(txt, &markups, tb.ModeHTML)
	machine[c.Sender().ID] = user
	return err
}

// Функция для взятия пола
func HandleGender(c tb.Context, user models.AllInformation) error {
	user.Gender = c.Text()
	user.State = "WAIT_AGE"

	txt := models.Phrases[user.Language]["ask_age"]

	err := c.Send(txt, tb.RemoveKeyboard, tb.ModeHTML)
	machine[c.Sender().ID] = user
	return err
}

// Функция для взятия возраста
func HandleAge(c tb.Context, user models.AllInformation) error {
	user.Age = c.Text()
	user.State = "WAIT_WEIGHT"

	txt := models.Phrases[user.Language]["ask_weight"]

	err := c.Send(txt, tb.ModeHTML)
	machine[c.Sender().ID] = user
	return err
}

// Функция для взятия веса
func HandleWeight(c tb.Context, user models.AllInformation) error {
	user.Weight = c.Text() // Исправил на Weight
	user.State = "WAIT_ACTIVITY"

	txt := models.Phrases[user.Language]["ask_activity"]
	b1 := models.Phrases[user.Language]["btn_light"]
	b2 := models.Phrases[user.Language]["btn_moderate"]
	b3 := models.Phrases[user.Language]["btn_heavy"]

	buttons := KeyboardButtons([]string{b1, b2, b3})
	var markups tb.ReplyMarkup
	markups.ReplyKeyboard = append(markups.ReplyKeyboard, buttons)
	markups.ResizeKeyboard = true

	err := c.Send(txt, &markups, tb.ModeHTML)
	machine[c.Sender().ID] = user
	return err
}

// Функция для взятия физической активности
func HandleActivity(c tb.Context, user models.AllInformation) error {
	user.Activity = c.Text() // Исправил на Active
	user.State = "WAIT_GOALS"

	txt := models.Phrases[user.Language]["ask_goals"]
	b1 := models.Phrases[user.Language]["btn_lose"]
	b2 := models.Phrases[user.Language]["btn_keep"]
	b3 := models.Phrases[user.Language]["btn_gain"]

	buttons := KeyboardButtons([]string{b1, b2, b3})
	var markups tb.ReplyMarkup
	markups.ReplyKeyboard = append(markups.ReplyKeyboard, buttons)
	markups.ResizeKeyboard = true

	err := c.Send(txt, &markups, tb.ModeHTML)
	machine[c.Sender().ID] = user
	return err
}

// Функция для взятия целей
func HandleGoals(c tb.Context, user models.AllInformation) error {
	user.Goal = c.Text() // Исправил на Goal
	user.State = "COMPLETED"

	// Финальное сообщение (можно добавить в словарь как "finish")
	thanks := ""
	if user.Language == "kz" {
		thanks = "✅ <b>Тіркелу аяқталды!</b>\nМәліметтер сақталды."
	} else {
		thanks = "✅ <b>Регистрация завершена!</b>\nДанные сохранены."
	}

	err := c.Send(thanks, tb.RemoveKeyboard, tb.ModeHTML)

	// Тут можно либо удалить юзера из machine, либо оставить для работы
	machine[c.Sender().ID] = user

	fmt.Printf("User %d registered: %+v\n", c.Sender().ID, user)
	return err
}
