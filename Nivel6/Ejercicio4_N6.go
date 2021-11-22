package main

import "fmt"

type persona struct {
	pnombre string
	snombre string
	edad    int
}

func (p persona) presentar() {
	fmt.Println("Mi nombre es: ", p.pnombre, p.snombre, "y tengo ", p.edad, "años de edad.")
}

func main() {
	p1 := persona{
		pnombre: "Carlos",
		snombre: "Muñoz",
		edad:    32,
	}

	p1.presentar()
}
