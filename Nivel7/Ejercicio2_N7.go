package main

import "fmt"

type persona struct {
	nombre string
}

func main() {
	p1 := persona{
		nombre: "Carlos Muñoz",
	}
	fmt.Println(p1)
	cambiame(&p1)
	fmt.Println(p1)
}

func cambiame(p *persona) {
	p.nombre = "Patricio Campos"
	// (*p).nombre = "Patricio Campos" tambien es valido
}
