package main

import (
	"errors"
	"fmt"
)

func main() {

	// Tipo de variáveis
	/*
		Variáveis podem ser explicitos e implicitos
	*/

	var var1 string = "algo" // Variável explicita
	variavel2 := "Var2"      // Variável implicita (determinado o tipo pelo valor) - inferência de tipos

	var (
		v1 string = "ow"
		v2 string = "ow"
	)

	v3, v4 := "variavel", "varaivelasdad"

	fmt.Println(var1)
	fmt.Println(variavel2)
	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println(v3)
	fmt.Println(v4)

	const vvvv string = "kiokiojkiok"
	fmt.Println(vvvv)

	// Tipos de dados

	var ( // Suportam números negativos
		varint  int8  = 1
		varint2 int16 = 20
		varint3 int32 = 302
		varint4 int64 = 230
	)

	// uint não suporta numeros negativos
	/*
		uint8
		uint16
		uint32
		uint64
	*/

	// alias
	var num3 rune = 123
	fmt.Println(num3)

	var numb byte = 8
	fmt.Println(numb)

	var fl float32
	var fl2 float64

	fmt.Println(fl, fl2)

	fmt.Printf("%d %d %d %d \n", varint, varint2, varint3, varint4)

	ch := 'B'
	fmt.Println(ch)

	var bo bool
	fmt.Println(bo)

	var erro error = errors.New("Erro")
	fmt.Println(erro) // nill (nulo)

}
