package handler

import (
	"database/sql"
	"fmt"

	tb "gopkg.in/telebot.v3"
)

func SearchIngredients(conn *sql.DB, c tb.Context) error {
	userID := c.Sender().ID
	var lang string

	// 1. Достаем язык из БД
	err := conn.QueryRow("SELECT language FROM users WHERE telegram_id = ?", userID).Scan(&lang)
	if err != nil {
		lang = "ru" // Страховка, если что-то пошло не так
	}

	queryText := "%" + c.Text() + "%"

	// 2. Ищем рецепты, подставляя язык в CASE
	rows, err := conn.Query(`
		SELECT DISTINCT r.id, 
		       CASE WHEN ? = 'kz' THEN r.name_kz ELSE r.name_ru END
		FROM recipes r
		JOIN recipe_ingredients ri ON r.id = ri.recipe_id
		JOIN ingredients i ON ri.ingredient_id = i.id
		WHERE i.name_ru LIKE ? OR i.name_kz LIKE ?
		LIMIT 10`, lang, queryText, queryText)

	if err != nil {
		fmt.Println("Search Error:", err)
		return c.Send("Ошибка при поиске.")
	}
	defer rows.Close()

	// 3. Локализация заголовков
	header := "🔍 <b>Вот что я нашел:</b>\n\n"
	notFoundMsg := "Ничего не найдено по вашему запросу 😔"
	if lang == "kz" {
		header = "🔍 <b>Табылған мәзірлер:</b>\n\n"
		notFoundMsg = "Сіздің сұранысыңыз бойынша ештеңе табылмады 😔"
	}

	var results string
	found := false
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			continue
		}
		results += fmt.Sprintf("🍴 <b>%s</b>\n", name)
		found = true
	}

	if !found {
		return c.Send(notFoundMsg)
	}

	return c.Send(header+results, tb.ModeHTML)
}
