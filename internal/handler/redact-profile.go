package handler

import (
	"database/sql"
	"diplomkabot/internal/db"
	"diplomkabot/internal/logic/rations"
	"diplomkabot/internal/models"
	"fmt"

	tb "gopkg.in/telebot.v3"
)

type RedactProfileS struct {
	Id    int64
	State string
	Btn   string
	Value string
}

// хранилище FSM machine
var storage = make(map[int64]RedactProfileS)

// обработка Редакт-информации
func RedactProfile(conn *sql.DB, c tb.Context) error {
	var user models.AllInformation

	// Берем ID того, кто отправил запрос
	senderID := c.Sender().ID
	// повесил в FSM machine стартовое значение
	storage[senderID] = RedactProfileS{
		Id:    senderID,
		State: "WAIT_BTN",
	}

	query := `
		SELECT 
			telegram_id, language, gender, age, weight, 
			height, activity, goal
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
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Send("❌ <b>Анкета не найдена.</b>\nПожалуйста, сначала пройдите регистрацию командой /start", tb.ModeHTML)
		}
		// Логируем другие ошибки БД (например, проблемы с подключением)
		fmt.Println("Database error:", err)
		return err
	}

	information := fmt.Sprintf(
		"👤 <b>Редактировать профиль:</b>\n\n"+
			"1 - <b>Язык:</b> %s\n"+
			"2 - <b>Пол:</b> %s\n"+
			"3 - <b>Возраст:</b> %s лет\n"+
			"4 - <b>Вес:</b> %s кг\n"+
			"5 - <b>Рост:</b> %s см\n"+
			"6 - <b>Активность:</b> %s\n"+
			"7 - <b>Цель:</b> %s\n\n",
		user.Language,
		user.Gender,
		user.Age,
		user.Weight,
		user.Height,
		user.Activity,
		user.Goal,
	)

	markup := RedactProfileKb(c)
	markup.ResizeKeyboard = true
	return c.Send(information, &markup, tb.ModeHTML)
}

// проверка находится ли FSM активным
func IsInRedactProcess(userID int64) bool {
	state, ok := storage[userID]
	return ok && state.State != ""
}

func HandleRedactState(conn *sql.DB, c tb.Context) error {
	senderID := c.Sender().ID
	// add user from db
	var user models.AllInformation

	query := `
		SELECT 
			telegram_id, language, gender, age, weight, 
			height, activity, goal
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
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Send("❌ <b>Анкета не найдена.</b>\nПожалуйста, сначала пройдите регистрацию командой /start", tb.ModeHTML)
		}
		// Логируем другие ошибки БД (например, проблемы с подключением)
		fmt.Println("Database error:", err)
		return err
	}

	// FSM
	session, ok := storage[senderID]
	if !ok {
		return nil
	}

	switch session.State {
	case "WAIT_BTN":
		switch c.Text() {
		case "1":
			session.Btn = c.Text()
			session.State = "WAIT_VALUE"

			// kb
			buttons := KeyboardButtons([]string{"🇰🇿 Қазақ тілі", "🇷🇺 Русский язык"})
			var markups tb.ReplyMarkup
			markups.ReplyKeyboard = append(markups.ReplyKeyboard, buttons)
			markups.ResizeKeyboard = true

			storage[senderID] = session
			return c.Send("⚙️ <b>Выберите язык / Тілді таңдаңыз</b> ", &markups, tb.ModeHTML)
		case "2":
			session.Btn = c.Text()
			session.State = "WAIT_VALUE"

			// kb
			buttons := KeyboardButtons([]string{"👨 Мужской", "👩 Женский"})
			var markups tb.ReplyMarkup
			markups.ReplyKeyboard = append(markups.ReplyKeyboard, buttons)
			markups.ResizeKeyboard = true

			storage[senderID] = session
			return c.Send("👤 <b>Ваш пол</b> ", &markups, tb.ModeHTML)
		case "3":
			session.Btn = c.Text()
			session.State = "WAIT_VALUE"

			storage[senderID] = session
			return c.Send("🎂 <b>Возраст</b>", tb.RemoveKeyboard, tb.ModeHTML)
		case "4":
			session.Btn = c.Text()
			session.State = "WAIT_VALUE"

			storage[senderID] = session
			return c.Send("⚖️ <b>Вес</b>", tb.RemoveKeyboard, tb.ModeHTML)
		case "5":
			session.Btn = c.Text()
			session.State = "WAIT_VALUE"

			storage[senderID] = session
			return c.Send("📏 <b>Рост</b>", tb.RemoveKeyboard, tb.ModeHTML)
		case "6":
			session.Btn = c.Text()
			session.State = "WAIT_VALUE"

			// kb
			buttons := KeyboardButtons([]string{"🟢 Легкие нагрузки", "🟡 Умеренные нагрузки", "🔴 Сильные нагрузки"})
			var markups tb.ReplyMarkup
			markups.ReplyKeyboard = append(markups.ReplyKeyboard, buttons)
			markups.ResizeKeyboard = true

			storage[senderID] = session
			return c.Send("🏃 <b>Интенсивность нагрузок</b>", &markups, tb.ModeHTML)
		case "7":
			session.Btn = c.Text()
			session.State = "WAIT_VALUE"

			// kb
			buttons := KeyboardButtons([]string{"📉 Сбросить вес", "⚖️ Поддержка массы", "📈 Набрать массу"})
			var markups tb.ReplyMarkup
			markups.ReplyKeyboard = append(markups.ReplyKeyboard, buttons)
			markups.ResizeKeyboard = true

			storage[senderID] = session
			return c.Send("🎯 <b>Шаг 7 из 8: Ваша цель</b>", &markups, tb.ModeHTML)
		case "Заполнить заново":
			delete(storage, senderID)
			db.DeleteUser(conn, senderID)
			return Registration(conn, c)
		}
	case "WAIT_VALUE":
		session.Value = c.Text()
		storage[senderID] = session

		switch session.Btn {
		case "1":
			user.Language = session.Value
		case "2":
			user.Gender = session.Value
		case "3":
			user.Age = session.Value
		case "4":
			user.Weight = session.Value
		case "5":
			user.Height = session.Value
		case "6":
			user.Activity = session.Value
		case "7":
			user.Goal = session.Value
		}

		rations.CalculateAll(&user)
		db.AddToDatabase(conn, user)
		delete(storage, senderID)
		return Profile(conn, c)
	}
	return nil
}
