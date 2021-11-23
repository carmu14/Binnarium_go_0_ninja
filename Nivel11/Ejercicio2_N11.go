//Resolver ejercicio2 que se encuentra en la direccion https://go.dev/play/p/kbeQfbe82Nl
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type persona struct {
	Nombre   string
	Apellido string
	Frases   []string
}

func main() {
	p1 := persona{
		Nombre:   "James",
		Apellido: "Bond",
		Frases:   []string{"Shaken, not stirred", "¿Algún último deseo?", "Nunca digas nunca."},
	}

	bs, err := aJSON(p1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}

// aJSON también necesita retorna un error
func aJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		//return []byte{}, fmt.Errorf("Hubo un error en aJSON: %v", err) //sol1
		return []byte{}, errors.New(fmt.Sprintf("Hubo un error en aJSON: %v", err)) //sol2
	}
	return bs, nil
}
