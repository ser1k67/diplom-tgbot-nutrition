-- SQLite
-- 1. Уникальные продукты (названия на двух языках)
CREATE TABLE IF NOT EXISTS ingredients (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name_ru TEXT NOT NULL, 
    name_kz TEXT NOT NULL,
    calories REAL, 
    proteins REAL,
    fats REAL,
    carbs REAL
);

-- 2. Рецепты (названия и инструкции на двух языках)
CREATE TABLE IF NOT EXISTS recipes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name_ru TEXT NOT NULL,
    name_kz TEXT NOT NULL,
    description_ru TEXT,
    description_kz TEXT,
    is_premium INTEGER DEFAULT 0, 
    category TEXT -- можно оставить на английском (breakfast/lunch/dinner) для логики кода
);

-- 3. Состав рецепта (связи остаются прежними, так как это ID)
CREATE TABLE IF NOT EXISTS recipe_ingredients (
    recipe_id INTEGER,
    ingredient_id INTEGER,
    weight_g REAL, 
    PRIMARY KEY (recipe_id, ingredient_id),
    FOREIGN KEY (recipe_id) REFERENCES recipes(id),
    FOREIGN KEY (ingredient_id) REFERENCES ingredients(id)
);