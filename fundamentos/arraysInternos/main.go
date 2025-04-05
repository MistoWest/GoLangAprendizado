package main

import "fmt"

func main() {

	// Arrays internos
	slice := make([]float32, 10, 11)

	fmt.Println(slice)
	slice = append(slice, 5)
	fmt.Println(slice)

	// Quando vê que seu slice irá estourar o tamanho, ele cria outro array e dobra o tamanho
	slice = append(slice, 5)
	fmt.Println(slice)
	fmt.Printf("%d %d \n", len(slice), cap(slice))

}
