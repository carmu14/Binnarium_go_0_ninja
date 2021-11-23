package main

import (
	"fmt"
	"github.com/carmu14/go_binnarium/Nivel12/Ejercicio1/perro"
)

type canino struct {
	nombre string
	edad   int
}

func main() {
	c1 := canino{
		nombre: "Dug",
		edad:   perro.Edad(2),
	}
	fmt.Println(c1)
}
