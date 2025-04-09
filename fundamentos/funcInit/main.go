package main

import "fmt"

// Funções init

var n int

func init() {
	fmt.Println("Executando a função init")
	n = 10

}

func main() {
	fmt.Println("Executando a função main")
	fmt.Println(n)
}
