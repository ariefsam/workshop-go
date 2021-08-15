package main

import "database/sql"

func InsertPerson(firstName, lastName, address string) (err error) {
	db, err := sql.Open("sqlite3", "name.db")
	if err != nil {
		return
	}
	db.SetMaxOpenConns(1)

	sqlStmt := `
	INSERT INTO person (first_name, last_name, address) VALUES (?,?, ?);
	`
	_, err = db.Exec(sqlStmt, firstName, lastName, address)
	return
}
