package route

import (
	"database/sql"
	"diplomkabot/internal/handler"
	"diplomkabot/internal/models"
	"fmt"

	tb "gopkg.in/telebot.v3"
)

// Функция обработки команд от пользователя
func Routers(conn *sql.DB, bot *tb.Bot) {
	// 1. Команды
	bot.Handle("/start", func(c tb.Context) error {
		return handler.Registration(conn, c)
	})

	bot.Handle("/profile", func(ctx tb.Context) error {
		return handler.Profile(conn, ctx)
	})

	// 2. ОБРАБОТКА КНОПОК ЗАМЕНЫ (ИНЛАЙН)
	// Создаем шаблон разметки, чтобы привязаться к ID кнопки "replace_meal"
	selector := &tb.ReplyMarkup{}
	btnReplace := selector.Data("", "replace_meal")

	// Этот хендлер ловит нажатия на "♻️ Заменить блюдо"
	bot.Handle(&btnReplace, func(c tb.Context) error {
		err := handler.DayPlan(conn, c, models.AllInformation{})
		if err != nil {
			// Если текст не изменился (выпал тот же рандом), просто уведомляем
			if err.Error() == "telegram: Bad Request: message is not modified (400)" {
				return c.Respond(&tb.CallbackResponse{Text: "Выпал тот же рецепт, попробуйте еще раз!"})
			}
			fmt.Println("Ошибка обновления плана:", err)
		}
		return c.Respond() // Убирает крутилку (часики) на кнопке
	})

	// 3. ТЕКСТОВЫЕ СООБЩЕНИЯ И КНОПКИ МЕНЮ
	bot.Handle(tb.OnText, func(c tb.Context) error {
		userID := c.Sender().ID

		// Проверка состояния регистрации
		if handler.IsRegistering(userID) {
			return handler.MachineState(conn, c)
		}

		text := c.Text()
		switch text {
		case "👤 Мой профиль":
			return handler.Profile(conn, c)

		case "🥗 План питания":
			return handler.DayPlan(conn, c, models.AllInformation{})

		case "🔍 Поиск рецептов":
			// ВКЛЮЧАЕМ режим поиска
			handler.SearchMode[userID] = true
			return c.Send("Введите название продукта (например: Курица):")

		case "⚙️ Редактировать профиль":
			return handler.Registration(conn, c)

		default:
			// Если юзер в режиме поиска
			if handler.SearchMode[userID] {
				delete(handler.SearchMode, userID)
				return handler.SearchIngredients(conn, c)
			}

			// Если ничего не подошло — показываем главное меню
			menu := handler.MainMenu(c)
			return c.Send("Выберите действие:", &menu)
		}
	})
}
