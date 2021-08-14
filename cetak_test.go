package main

import "testing"

func TestCetak(t *testing.T) {
	var addr Alamat
	addr.AlamatLengkap = "Jalan Kemang Timur No 90"
	addr.Provinsi = "DKI Jakarta"
	addr.Kota = "Jakarta Selatan"
	addr.Kecamatan = "Mampang Prapatan"
	addr.Kelurahan = "Bangka"
	addr.KodePos = "12730"

	expected := `Jalan Kemang Timur No 90
Kelurahan Bangka
Kecamatan Mampang Prapatan
Kota Jakarta Selatan
Provinsi DKI Jakarta`
	actual := Cetak(addr)

	if actual != expected {
		t.Error("Want \n"+expected, ", got \n"+actual)
	}

}
