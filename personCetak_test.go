package main

import "testing"

func TestPersonCetak(t *testing.T) {
	var data Person
	data.FirstName = "Budi"
	data.LastName = "Susanto"

	var cetakNama string
	cetakNama = data.Cetak()

	if cetakNama != "Budi Susanto" {
		t.Errorf("Cetak nama gagal.")
	}
}
