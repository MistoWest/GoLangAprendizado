package main

import "fmt"

// Funções com retorno

func soma(a int, b int) (s int, sub int) {
	s = a + b
	sub = a + b
	return
}

// Funçoes variática
func somaVariatica(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

// Funções recursivas são funções que chamam a si mesmas
// Exemplo: função que calcula o fatorial de um número
// fatorial(5) = 5 * 4 * 3 * 2 * 1 = 120
// fatorial(4) = 4 * 3 * 2 * 1 = 24
// fatorial(0) = 1
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println(soma(1, 2))

	soma, sub := soma(1, 2)
	fmt.Println(soma, sub)

	fmt.Println(somaVariatica(1, 2, 3, 4, 5))

	// funções anonimas
	func() {
		fmt.Println("Hello, World!")
	}()

	funcA := func() {
		fmt.Println("Hello, World!")
	}

	funcA()
}
