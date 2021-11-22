package main

import "fmt"

func main() {
	x := map[string][]string{
		`carlos_munoz`:   []string{`computadoras`, `montañas`, `playa`},
		`maria_palacios`: []string{`leer`, `comprar`, `música`},
		`karina_campos`:  []string{`helado`, `pintar`, `bailar`},
	}

	fmt.Println(x)

	for llave, valor := range x {
		fmt.Println("Registro:", llave)
		for i, val := range valor {
			fmt.Println("\t", i, val)
		}
	}
}
