package main

import "fmt"

// Interfaces genéricas

// Não recomendado
func generica(i interface{}) {
	fmt.Println(i)
}

func main() {
	generica("teste")
}
