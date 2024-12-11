package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}

// 	_, err = db.Exec(`
// 	INSERT OR IGNORE INTO users (id, balance) VALUES (1, 10000.00)
// `)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

	_, err = db.Exec(`
	create table if not exists users (
		id integer primary key autoincrement,
		balance real not null
	)
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		symbol TEXT NOT NULL,
		quantity REAL NOT NULL,
		price REAL NOT NULL,
		type TEXT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
`)
	if err != nil {
		log.Fatal(err)
	}

	return db

}
