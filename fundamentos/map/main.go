package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	// DIferença de map pra struct é que o map parece um "dicionário", visto que pra acessar um valor, tenho que digitar a chave do valor como (usuario['nome'])

	fmt.Println(usuario)

	// Aninhado

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "José",
			"ultimo":   "Pedroca",
		},
		"curso": {
			"nome": "Engenharia",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "curso") //deletando um map

	usuario2["signo"] = map[string]string{
		"nome": "Leão",
	}
}
