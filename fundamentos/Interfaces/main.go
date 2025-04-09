package main

import "fmt"

// Interfaces
// Interfaces são conjuntos de métodos que uma struct deve implementar

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return c.raio * c.raio * 3
}

func main() {
	fmt.Println("Hello, World!")

	r := retangulo{
		altura:  10,
		largura: 5,
	}

	escreverArea(r)

	c := circulo{
		raio: 10,
	}

	escreverArea(c)

}
