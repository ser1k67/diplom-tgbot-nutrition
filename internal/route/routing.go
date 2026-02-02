package route

import (
	"database/sql"
	"diplomkabot/internal/handler"

	tb "gopkg.in/telebot.v3"
)

// Функция обработки команд от пользователя
func Routers(conn *sql.DB, bot *tb.Bot) {
	//Роут на регистрацию нового пользователя
	bot.Handle("/start", func(c tb.Context) error {
		return handler.Registration(conn, c)
	})
	bot.Handle("/profile", func(ctx tb.Context) error {
		return handler.Profile(conn, ctx)
	})
	bot.Handle(tb.OnText, func(ctx tb.Context) error {
		return handler.MachineState(conn, ctx)
	})
}
