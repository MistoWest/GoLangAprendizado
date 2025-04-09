package main

import "fmt"

// Funções closure
// Funções closure são funções que podem acessar variáveis de fora do seu escopo

// Exemplo:

func closure() func() {
	texto := "Dentro da função closure"

	return func() {
		fmt.Println(texto)
	}
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)
	numero := closure()
	numero()
}
