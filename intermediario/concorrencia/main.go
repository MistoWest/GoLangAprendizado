package main

import (
	"fmt"
	"time"
)

// Concorrência e paralelismo (Goroutines)

func main() {
	// CONCORRÊNCIA != PARALELISMO

	go escrever("OLá") // Goroutine
	go escrever("Olá 2")
	escrever("Olá 3")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
