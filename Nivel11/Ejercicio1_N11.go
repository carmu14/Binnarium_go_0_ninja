//Resolver ejercicio1 que se encuentra en la direccion https://go.dev/play/p/Bgads38iFgz
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Nombre   string
	Apellido string
	Frases   []string
}

func main() {
	p1 := person{
		Nombre:   "James",
		Apellido: "Bond",
		Frases:   []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bs))

}
