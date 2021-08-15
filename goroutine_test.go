package main

import (
	"fmt"
	"testing"
	"time"
)

func CetakNumber(i int) {
	fmt.Println("Mencetak ", i)
}

func TestGoroutine(t *testing.T) {
	for i := 0; i <= 10; i++ {
		go CetakNumber(i)
	}
	time.Sleep(1 * time.Second)
}
