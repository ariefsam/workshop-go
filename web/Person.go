package main

import "log"

type Person struct {
	FirstName string
	LastName  string
}

func (p *Person) Cetak() (name string) {
	name = p.FirstName + " " + p.LastName
	return
}

func (p *Person) GetFirstName() (firstname string) {
	firstname = p.FirstName
	return
}
func (p *Person) CetakTanpaOutput() {
	log.Println(p.Cetak())
}

func (p Person) TanpaPointer() {
	p.FirstName = "Doni"
}

func (p *Person) DenganPointer() {
	p.FirstName = "Doni"
}
