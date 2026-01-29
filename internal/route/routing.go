package route

import (
	"diplomkabot/internal/handler"

	tb "gopkg.in/telebot.v3"
)

// Функция обработки команд от пользователя
func Routers(bot *tb.Bot) {
	//Роут на регистрацию
	bot.Handle("/start", func(conx tb.Context) error {
		err := handler.Registration(bot, conx)
		return err
	})
}
