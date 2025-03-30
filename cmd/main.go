package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	const FILE_PATH = "../db/store.db"

	file, err := os.OpenFile(FILE_PATH, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file")
		panic(err)
	}
	file.Close()

	db, err := sql.Open("sqlite3", FILE_PATH)
	if err != nil {
		fmt.Println("open connect")
		panic(err)
	}
	defer db.Close()

	err = createTablesIfNotExist(db)
	if err != nil {
		panic(err)
	}

}

func createTablesIfNotExist(conn *sql.DB) error {
	_, err := conn.Exec(`
		CREATE TABLE users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT UNIQUE NOT NULL,
			name TEXT NOT NULL
		);
	`)
	if err != nil {
		return err
	}

	return nil
}
