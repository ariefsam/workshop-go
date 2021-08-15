package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Read(firstName string) (persons []Person, err error) {
	db, err := sql.Open("sqlite3", "name.db")
	if err != nil {
		return
	}
	stmt := "SELECT first_name, last_name FROM person WHERE first_name=?"
	rows, err := db.Query(stmt, firstName)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p Person
		err = rows.Scan(&p.FirstName, &p.LastName)
		if err != nil {
			return
		}
		persons = append(persons, p)
	}

	return
}
