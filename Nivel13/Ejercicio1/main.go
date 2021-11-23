package main

import (
	"fmt"

	"github.com/carmu14/go_binnarium/Nivel13/Ejercicio1/perrito"
)

type canino struct {
	nombre string
	edad   int
}

func main() {
	detalle := canino{
		nombre: "Dug",
		edad:   perrito.Edad(10),
	}
	fmt.Println(detalle)
	fmt.Println(perrito.EdadDos(20))
}
