package handler

import tb "gopkg.in/telebot.v3"

//Функция для добавления кнопок выбора после регистрации
func MainMenu(c tb.Context) tb.ReplyMarkup {
	var keyboardButtons tb.ReplyMarkup
	row1 := KeyboardButtons([]string{"👤 Мой профиль", "🥗 План питания"})
	row2 := KeyboardButtons([]string{"Экспорт данных", "⚙️ Редактировать профиль"})
	row3 := KeyboardButtons([]string{"🔍 Поиск рецептов"})
	keyboardButtons.ReplyKeyboard = append(keyboardButtons.ReplyKeyboard, row1, row2, row3)
	return keyboardButtons
}
