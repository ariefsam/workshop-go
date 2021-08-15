package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is page about")
}

func Server() {
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/about", AboutHandler)
	http.HandleFunc("/template", TemplateHandler)
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

func main() {
	Server()
}

func Read(firstName string) (persons []Person, err error) {
	db, err := sql.Open("sqlite3", "../name.db")
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
