package main

import "fmt"

// Panic e recover
// Panic é uma função que para a execução do programa
// Recover é uma função que recupera o programa
// Panic e recover são usados para tratar erros
// Panic é usado para parar a execução do programa
// Recover é usado para recuperar o programa
// Panic e recover são usados para tratar erros

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 10 || media < 0 {
		panic("A média não pode ser maior que 10 ou menor que 0!")
	}

	return media >= 6
}

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
		fmt.Println("Erro:", r)
	}
}

func main() {
	fmt.Println(alunoEstaAprovado(8, 6))    // Aprovado
	fmt.Println(alunoEstaAprovado(4, 5))    // Reprovado
	fmt.Println(alunoEstaAprovado(1143, 4)) // Panic
}
