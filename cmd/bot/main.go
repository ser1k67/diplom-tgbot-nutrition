package main

import (
	"diplomkabot/internal/db"
	"fmt"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	tg "gopkg.in/telebot.v3"
)

func main() {
	//Инициализация сервера
	server := gin.Default()
	//Открытие файла .env для чтения
	err := godotenv.Load("../../configs/config.env")
	if err != nil {
		fmt.Println("Не найден .env файл, пожалуйста загрузите его в папку configs в корне проекта:", err)
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
		fmt.Println("Не получилось подключиться к боту", err)
	}
	println("Сервер запущен на порте: 8282")
	server.Run("8282") //запуск сервера

}
