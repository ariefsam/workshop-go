package main

import "testing"

func TestInsert(t *testing.T) {
	err := InsertPerson("Arief", "Hidayatulloh", "Bogor")
	if err != nil {
		t.Error("Gagal insert Person. " + err.Error())
	}
}
