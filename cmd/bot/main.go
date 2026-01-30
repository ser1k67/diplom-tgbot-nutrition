package main

import (
	"diplomkabot/internal/db"
	"diplomkabot/internal/route"

	"os"

	"github.com/joho/godotenv"
	tg "gopkg.in/telebot.v3"
)

func main() {
	//Открытие файла .env для чтения
	err := godotenv.Load("../../configs/config.env")
	if err != nil {
		println("Не найден .env файл, пожалуйста загрузите его в папку configs в корне проекта:", err)
		return
	}
	//Открываем подключени к базе данных
	conn := db.Connection()
	defer conn.Close() //Закрытие подключения к базе данных, после завершения программы

	//Инициализация бота
	bot, err := tg.NewBot(tg.Settings{
		Token: os.Getenv("TELEGRAM_BOT_TOKEN"),
	})
	if err != nil {
		println("Не получилось подключиться к боту", err)
	}
	route.Routers(conn, bot)
	println("Бот запущен")
	bot.Start()
}
