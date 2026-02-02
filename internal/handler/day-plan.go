package handler

import (
	"database/sql"
	"diplomkabot/internal/models"
	"fmt"

	tb "gopkg.in/telebot.v3"
)

type item struct {
	name     string
	weight   float64
	kcal100g float64
	p        float64
	f        float64
	carb     float64
}

func DayPlan(conn *sql.DB, c tb.Context, all models.AllInformation) error {
	var lang string
	var targetKcal float64
	var goal string
	userID := c.Sender().ID

	err := conn.QueryRow("SELECT language, target_kcal, goal FROM users WHERE telegram_id = ?", userID).Scan(&lang, &targetKcal, &goal)
	if err != nil {
		c.Send("Сначала пройдите регистрацию!")
		return err
	}

	mNames := map[string]string{"breakfast": "Завтрак", "second_breakfast": "Второй завтрак", "lunch": "Обед", "snack": "Полдник", "dinner": "Ужин"}
	mTimes := map[string]string{"breakfast": "08:00", "second_breakfast": "11:00", "lunch": "13:30", "snack": "16:00", "dinner": "19:00"}
	if lang == "kz" {
		mNames = map[string]string{"breakfast": "Таңғы ас", "second_breakfast": "Екінші таңғы ас", "lunch": "Түскі ас", "snack": "Бесін ас", "dinner": "Кешкі ас"}
	}
	ratios := map[string]float64{"breakfast": 0.25, "second_breakfast": 0.10, "lunch": 0.35, "snack": 0.10, "dinner": 0.20}
	if goal == "📉 Сбросить вес" || goal == "📉 Салмақ тастау" {
		ratios = map[string]float64{"breakfast": 0.30, "second_breakfast": 0.10, "lunch": 0.40, "snack": 0.05, "dinner": 0.15}
	}

	if c.Callback() != nil {
		category := c.Callback().Data
		kcalForMeal := targetKcal * ratios[category]

		mealText, _, _, _, _ := getMealData(conn, lang, category, kcalForMeal)

		newText := fmt.Sprintf("<b>%s</b> (%s)\n%s└ <i>Энергия: %.0f ккал</i>\n",
			mNames[category], mTimes[category], mealText, kcalForMeal)

		selector := &tb.ReplyMarkup{}
		btn := selector.Data("♻️ Заменить блюдо", "replace_meal", category)
		selector.Inline(selector.Row(btn))

		err := c.Edit(newText, tb.ModeHTML, selector)
		if err != nil {
			fmt.Println("Telegram Error Ignored:", err)
			return c.Respond(&tb.CallbackResponse{Text: "Попробуйте еще раз!"})
		}
		return nil
	}
	c.Send("<b>📅 Ваш план питания на сегодня:</b>", tb.ModeHTML)

	for cat, perc := range ratios {
		kcalForMeal := targetKcal * perc
		mealText, _, _, _, _ := getMealData(conn, lang, cat, kcalForMeal)

		report := fmt.Sprintf("<b>%s</b> (%s)\n%s└ <i>Энергия: %.0f ккал</i>",
			mNames[cat], mTimes[cat], mealText, kcalForMeal)

		selector := &tb.ReplyMarkup{}
		btn := selector.Data("♻️ Заменить блюдо", "replace_meal", cat)
		selector.Inline(selector.Row(btn))

		c.Send(report, tb.ModeHTML, selector)
	}

	return nil
}

// Вспомогательная функция, чтобы не дублировать код
func getMealData(conn *sql.DB, lang, category string, targetKcal float64) (text string, p, f, c float64, err error) {
	var recipeID int
	var recipeName, recipeDesc string

	// Ищем рандомный рецепт
	query := `
        SELECT id, 
               CASE WHEN ? = 'kz' THEN name_kz ELSE name_ru END,
               CASE WHEN ? = 'kz' THEN description_kz ELSE description_ru END
        FROM recipes 
        WHERE category = ? AND is_premium = 0 
        ORDER BY RANDOM() 
        LIMIT 1`

	err = conn.QueryRow(query, lang, lang, category).Scan(&recipeID, &recipeName, &recipeDesc)
	if err != nil {
		return "", 0, 0, 0, err
	}

	//Тянем ингредиенты
	rows, err := conn.Query(`
        SELECT 
            CASE WHEN ? = 'kz' THEN i.name_kz ELSE i.name_ru END,
            ri.weight_g, i.calories, i.proteins, i.fats, i.carbs
        FROM recipe_ingredients ri
        JOIN ingredients i ON ri.ingredient_id = i.id
        WHERE ri.recipe_id = ?`, lang, recipeID)
	if err != nil {
		return "", 0, 0, 0, err
	}
	defer rows.Close()

	var totalBaseKcal float64
	var items []item
	for rows.Next() {
		var it item
		if err := rows.Scan(&it.name, &it.weight, &it.kcal100g, &it.p, &it.f, &it.carb); err != nil {
			return "", 0, 0, 0, err
		}
		totalBaseKcal += (it.kcal100g * it.weight) / 100
		items = append(items, it)
	}

	if totalBaseKcal == 0 {
		return "", 0, 0, 0, fmt.Errorf("zero base kcal")
	}

	//Считаем коэффициент масштабирования
	coefficient := targetKcal / totalBaseKcal
	var ingredientsList string

	for _, it := range items {
		finalWeight := it.weight * coefficient
		ingredientsList += fmt.Sprintf(" - %s: %.0fг\n", it.name, finalWeight)

		// Считаем БЖУ для этого конкретного приема пищи
		p += (it.p * finalWeight) / 100
		f += (it.f * finalWeight) / 100
		c += (it.carb * finalWeight) / 100
	}

	// Формируем блок текста для этого приема пищи
	resText := fmt.Sprintf("<b>%s</b>\n%s", recipeName, ingredientsList)

	// Возвращаем текст, а также P, F, C (белки, жиры, углеводы) как отдельные числа
	return resText, p, f, c, nil
}
