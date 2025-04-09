package main

import (
	"fmt"
	"log"
	"o/app"
	"os"
)

// Aplicação com linha de comando para calcular a área de um retângulo

func main() {
	fmt.Println("Hello, World!")

	ap := app.Gerar()
	if erro := ap.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
