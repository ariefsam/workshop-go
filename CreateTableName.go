package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTableName() (err error) {
	db, err := sql.Open("sqlite3", "name.db")
	if err != nil {
		return
	}

	sqlStmt := `
	   CREATE TABLE if not exists person (
		   id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		   first_name TEXT NOT NULL DEFAULT "",
		   last_name TEXT NOT NULL DEFAULT "",
		   address TEXT NOT NULL DEFAULT "",
		   created_at NUMERIC NOT NULL DEFAULT CURRENT_TIMESTAMP,
		   update_at NUMERIC NOT NULL DEFAULT 0
	   );
	`

	_, err = db.Exec(sqlStmt)

	return
}
