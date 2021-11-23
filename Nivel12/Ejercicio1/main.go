//Package main para la ejecucion
package main

import (
	"fmt"

	"github.com/carmu14/go_binnarium/Nivel12/Ejercicio1/perro"
)

//Canino tipo detalle del perro
type canino struct {
	nombre string
	edad   int
}

//main es la funcion principal
func main() {
	c1 := canino{
		nombre: "Dug",
		edad:   perro.Edad(2),
	}
	fmt.Println(c1)
}
