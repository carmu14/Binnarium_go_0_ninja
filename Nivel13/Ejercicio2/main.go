package main

import (
	"fmt"

	"github.com/carmu14/go_binnarium/Nivel13/Ejercicio2/cita"
	"github.com/carmu14/go_binnarium/Nivel13/Ejercicio2/palabra"
)

func main() {
	fmt.Println(palabra.Conteo(cita.Cuando))

	for k, v := range palabra.ConteoUso(cita.Cuando) {
		fmt.Println(v, k)
	}
}
