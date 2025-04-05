package main

import (
	"fmt"
	"reflect"
)

// Arrays e Slices

func main() {

	// Array de 5 posições do tipo inteiro
	var arr [5]int64
	arr2 := [5]int64{1, 2, 3, 4, 5}

	fmt.Println(arr)
	fmt.Println(arr2)

	array3 := [...]int{1, 2, 3, 4}

	fmt.Println(array3)

	slice := []int{1, 2} // Isso é um slice (Dinâmico)
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf((array3)))

	slice = append(slice, 18)
	fmt.Println(slice)
}
