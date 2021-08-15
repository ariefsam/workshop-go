package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func CreateTablePerson() (err error) {
	db, err := sql.Open("mysql", "root:password@/workshop")
	if err != nil {
		return
	}

	sqlStmt := `
	   CREATE TABLE if not exists person (
		   id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
		   first_name TEXT,
		   last_name TEXT,
		   address TEXT,
		   created_at INT NOT NULL DEFAULT 0,
		   update_at INT NOT NULL DEFAULT 0
	   );
	`

	_, err = db.Exec(sqlStmt)
	return
}
