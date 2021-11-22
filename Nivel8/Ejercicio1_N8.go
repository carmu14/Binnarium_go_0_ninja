package main

import (
	"encoding/json"
	"fmt"
)

type usuario struct {
	Nombre string
	Edad   int
}

func main() {
	u1 := usuario{
		Nombre: "Carlos",
		Edad:   32,
	}

	u2 := usuario{
		Nombre: "Karina",
		Edad:   27,
	}

	u3 := usuario{
		Nombre: "Vicente",
		Edad:   55,
	}

	usuarios := []usuario{u1, u2, u3}

	fmt.Println(usuarios)

	bs, err := json.Marshal(usuarios)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
