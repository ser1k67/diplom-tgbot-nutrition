package route

import (
	"diplomkabot/internal/handler"

	tb "gopkg.in/telebot.v3"
)

// Функция обработки команд от пользователя
func Routers(bot *tb.Bot) {
	//Роут на регистрацию нового пользователя
	bot.Handle("/start", handler.Registration)
	bot.Handle(tb.OnText, handler.MachineState)
}
