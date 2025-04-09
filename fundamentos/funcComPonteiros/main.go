package main

import "fmt"

func mudarValor(n *int, valor int) {
	*n = valor
}

func main() {

	var valor int = 10

	mudarValor(&valor, 2)
	fmt.Println(valor)
}
