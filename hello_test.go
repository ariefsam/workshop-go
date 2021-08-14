package main

import "testing"

func TestHello(t *testing.T) {
	var actual, expected string

	expected = "Hello World!"
	actual = Hello()

	if actual != expected {
		t.Error("Fungsi tidak mengembalikan Hello World!")
	}
}
