package main

import (
	"fmt"
)

type user struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	log    string
	numero uint8
}

func main() {
	var us user
	us.nome = "Davi"
	us.idade = 21
	us.endereco.log = "R21"
	us.endereco.numero = 2

	u2 := user{"Davi", 2, endereco{"Rua Boboca", 2}}

	fmt.Println(us)
	fmt.Println(u2)

	u3 := user{idade: 34} // Não foi preenchdio o nome mas foi preenchido a idade
	fmt.Println(u3)

}
