package main

import "fmt"

func main() {

	var value int = 43
	var value2 *int = &value

	*value2 += 1 // o value2 aponta pra um endereço de memória, que por consequẽncia conseguimos mexer nesse valor através do valor de referência

	// value2 = mostra endereço
	// *value2 = mostra o valor do endereço, podendo alterar valores dentro dele
	_ = value2

	fmt.Println(value)
}
