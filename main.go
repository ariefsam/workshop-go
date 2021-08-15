package main

import (
	"fmt"
	"log"
)

type Animal struct {
	Name  string
	Sound string
	Age   int
	X     func(int) int
}

func (a Animal) PrintSound() {
	fmt.Println(a.Sound)
}

func main() {

	var data Person
	data.FirstName = "Budi"
	data.LastName = "Susanto"

	cetak := data.Cetak()
	log.Print(cetak)

	// data.CetakTanpaOutput()
	// data.TanpaPointer()
	// log.Println(data.FirstName)

	// data.DenganPointer()
	// log.Println(data.FirstName)
	// log.Println(data.GetFirstName())
	// log.Println(data.Cetak())
	// var pod []Animal
	// pod = []Animal{
	// 	{
	// 		Name:  "Heli",
	// 		Sound: "Woof",
	// 		Age:   4,
	// 		X: func(i int) (o int) {
	// 			return i * 2
	// 		},
	// 	},
	// 	{
	// 		Name:  "Donald",
	// 		Sound: "Wak",
	// 		Age:   3,
	// 	},
	// }
	// for idx, value := range pod {
	// 	fmt.Println("\nindex:", idx)
	// 	value.PrintSound()
	// 	name := value.Name
	// 	fmt.Println("Name:", name)
	// 	if value.X != nil {
	// 		fmt.Println(value.X(3))
	// 	}
	// }
}
