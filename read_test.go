package main

import (
	"log"
	"testing"
)

func TestRead(t *testing.T) {
	var persons []Person

	persons, err := Read("Arief")
	if err != nil {
		t.Error("Error ketika baca " + err.Error())
	}

	if len(persons) == 0 {
		t.Error("Tidak ada hasil")
	}

	log.Printf("%+v", persons)
}
