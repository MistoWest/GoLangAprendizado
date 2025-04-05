// herança é feita de forma abstrata, não quer dizer que seja possivel fazer herança em GO

package main

import "fmt"

// "Herança"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa    // Não preciso especificar o nome da variavel
	curso     string
	faculdade string
}

func main() {

	p1 := pessoa{"João", "Pessoa", 20, 178}
	es := estudante{p1, "Computação", "IF GOIANO"}
	fmt.Println(p1)
	fmt.Println(es)
}
