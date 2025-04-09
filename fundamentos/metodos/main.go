package main

import "fmt"

// Metodos
// Metodos são funções que estão associadas a um tipo

type Usuario struct {
	nome  string
	idade int
}

func (u Usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s\n", u.nome)
}

func main() {
	var us Usuario = Usuario{
		nome:  "João",
		idade: 20,
	}

	us.salvar()

	fmt.Println(us)
}
