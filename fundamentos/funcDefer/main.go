package main

// Defer é uma função que será executada após a função main terminar
// Ele é útil para fechar conexões, arquivos, etc
// Ele é executado na ordem inversa da declaração
// Ele é executado mesmo que ocorra um panic
// Ele é executado mesmo que ocorra um erro
// Ele é executado mesmo que ocorra um return
// Ele é executado mesmo que ocorra um break
// Ele é executado mesmo que ocorra um continue
// Ele é executado mesmo que ocorra um goto

import "fmt"

func main() {
	defer fmt.Println("Hello, dasdas!") // Ele será executado após a função main terminar
	fmt.Println("Hello, World!")
}
