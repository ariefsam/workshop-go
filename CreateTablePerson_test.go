package main

import "testing"

func TestCreateTablePerson(t *testing.T) {
	err := CreateTablePerson()
	if err != nil {
		t.Error(err.Error())
	}
}
