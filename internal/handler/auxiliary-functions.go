package handler

import (
	"database/sql"
	"diplomkabot/internal/db"
	"diplomkabot/internal/logic/rations"
	"diplomkabot/internal/models"
	"fmt"

	tb "gopkg.in/telebot.v3"
)

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

func HandleGender(c tb.Context, user models.AllInformation) error {
	user.Gender = c.Text()
	user.State = "WAIT_AGE"

	txt := models.Phrases[user.Language]["ask_age"]

	err := c.Send(txt, tb.RemoveKeyboard, tb.ModeHTML)
	machine[c.Sender().ID] = user
	return err
}

func HandleAge(c tb.Context, user models.AllInformation) error {
	user.Age = c.Text()
	user.State = "WAIT_WEIGHT"

	txt := models.Phrases[user.Language]["ask_weight"]

	err := c.Send(txt, tb.ModeHTML)
	machine[c.Sender().ID] = user
	return err
}

func HandleWeight(c tb.Context, user models.AllInformation) error {
	user.Weight = c.Text()
	user.State = "WAIT_HEIGHT"

	// Исправлено: теперь спрашиваем РОСТ после веса
	txt := models.Phrases[user.Language]["ask_height"]

	err := c.Send(txt, tb.ModeHTML)
	machine[c.Sender().ID] = user
	return err
}

func HandleHeight(c tb.Context, user models.AllInformation) error {
	user.Height = c.Text()
	user.State = "WAIT_ACTIVITY"

	// Исправлено: теперь спрашиваем АКТИВНОСТЬ после роста
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

func HandleActivity(c tb.Context, user models.AllInformation) error {
	user.Activity = c.Text()
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

func HandleGoals(conn *sql.DB, c tb.Context, user models.AllInformation) error {
	user.Goal = c.Text()
	user.State = "COMPLETED"
	rations.CalculateAll(&user)
	thanks := ""
	if user.Language == "kz" {
		thanks = "✅ <b>Тіркелу аяқталды!</b>\nМәліметтер сақталды."
	} else {
		thanks = "✅ <b>Регистрация завершена!</b>\nДанные сохранены."
	}
	markup := MainMenu(c)
	err := c.Send(thanks, &markup, tb.ModeHTML)

	// Сохраняем в мапу и БД
	machine[c.Sender().ID] = user
	err = db.AddToDatabase(conn, user)
	if err != nil {
		fmt.Println("Ошибка вставки в БД:", err)
		return err
	}
	c.Send(thanks)
	return Profile(conn, c)
}
