package main

import "fmt"

type persona struct {
	nombre string
}

func (p *persona) hablar() {
	fmt.Println("Hola soy", p.nombre)
}

type humano interface {
	hablar()
}

func diAlgo(h humano) {
	h.hablar()
}

func main() {
	p1 := persona{
		nombre: "Carlos",
	}

	//diAlgo(p1) es valido
	//diAlgo(&p1) es valido
	p1.hablar()
}
