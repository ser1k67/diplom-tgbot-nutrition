package handler

import (
	tb "gopkg.in/telebot.v3"
)

// Распознавание
type Rus struct {
}
type Kaz struct {
}
type AllInformation struct {
	Gender string
	Age    string
	Weight string
	Active string
	Goal   string
}

// Обработка регистрации пользователя
func Registration(bot *tb.Bot, c tb.Context) error {
	var err error
	//Выбор языка

	return err
}
