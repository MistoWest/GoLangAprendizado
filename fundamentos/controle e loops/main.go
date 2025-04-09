package main

// Estrutura de controle
import (
	"fmt"
)

func main() {
	numero := 10

	// If Else
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("O número é maior que 5")
	} else {
		fmt.Println("O número é menor que 5")
	}

	// Switch
	switch numero {
	case 1:
		fmt.Println("O número é 1")
	case 2:
		fmt.Println("O número é 2")
	default:
		fmt.Println("O número não é 1 ou 2")
	}

	// Loops
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Loops com range
	for i, v := range []int{1, 2, 3} {
		fmt.Println(i, v) // i =
	}

	// Loops com while
	for numero > 0 {
		fmt.Println(numero)
		numero--
	}

	// Loops com break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
	}

	// Loops com continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
	}
}
