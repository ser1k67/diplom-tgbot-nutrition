-- SQLite
CREATE TABLE IF NOT EXISTS users (
    telegram_id INTEGER PRIMARY KEY,
    language    TEXT,
    gender      TEXT,
    age         TEXT,
    weight      TEXT,
    activity    TEXT,
    goal        TEXT,
    height      TEXT,
    bmi         REAL,
    bmr         REAL,
    tdee        REAL,
    target_kcal REAL,
    proteins    REAL, -- Поменяй на REAL
    fats        REAL, -- Поменяй на REAL
    carbs       REAL, -- Поменяй на REAL
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP
);