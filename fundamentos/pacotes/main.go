package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {

	erro := checkmail.ValidateFormat("dev@gmail.com")
	fmt.Println(erro)

	auxiliar.M()
	fmt.Println("Hello")
}
