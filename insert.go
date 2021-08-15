package main

import (
	"database/sql"
	"log"
)

func InsertPerson(firstName, lastName, address string) (err error) {
	//db, err := sql.Open("mysql", "root:password@/workshop")
	db, err := sql.Open("sqlite3", "name.db")
	if err != nil {
		return
	}
	db.SetMaxOpenConns(1)

	sqlStmt := `
	INSERT INTO person (first_name, last_name, address) VALUES (?,?, ?);
	`
	result, err := db.Exec(sqlStmt, firstName, lastName, address)
	lastID, _ := result.LastInsertId()
	rowsAffected, _ := result.RowsAffected()
	log.Println("insert id: ", lastID, "\naffected rows:", rowsAffected)
	return
}
