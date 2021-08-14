package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	input := []string{"Halo", "halo", "Bandung"}

	for index, value := range input {
		fmt.Println("Key: ", index, "value: ", value)
	}

	for {
		fmt.Println("for tanpa batas")
		break
	}
}
