package main

import "fmt"

func soma(i, y int8) int8 { // ou i int8, y int8
	return i + y
}

func Mat(n1, n2 int8) (int8, int8) { // Dois retornos
	soma := n1 + n2
	sub := n1 - n2

	return soma, sub // é possível mais de um retorno
}

func main() {
	fmt.Printf("%d ", soma(1, 5))

	// função em variável
	var f = func(txt string) string {
		return txt
	}

	fmt.Println(f("ok"))

	soma, sub := Mat(1, 2)
	_ = sub // Caso eu não queira usar

	// OU

	somas, _ := Mat(3, 4) // caso eu só queira o primeiro resultado, usar _

	fmt.Println(soma)
	fmt.Println(somas)

}
