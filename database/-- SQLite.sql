-- SQLite
CREATE TABLE users (
    telegram_id INTEGER PRIMARY KEY, -- Оставляем INTEGER, так как это ID
    language    TEXT,
    gender      TEXT,
    age         TEXT,
    weight      TEXT,
    activity    TEXT,
    goal        TEXT,
    height      TEXT,
    
    -- Результаты расчетов (тоже TEXT)
    bmi         TEXT,
    bmr         TEXT,
    tdee        TEXT,
    target_kcal TEXT,
    
    -- БЖУ
    proteins    TEXT,
    fats        TEXT,
    carbs       TEXT,
    
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP
);