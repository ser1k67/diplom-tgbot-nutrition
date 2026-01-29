package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func Connection() *sql.DB {
	conn, err := sql.Open("sqlite", "../../database")
	if err != nil {
		fmt.Println("Не удалось выполнить подключение к базе данных:", err)
		return nil
	}
	return conn
}
