package main

import "fmt"

func main() {
	s := struct {
		nombre     string
		amigos     map[string]int
		bebidasFav []string
	}{
		nombre: "Carlos",
		amigos: map[string]int{
			"Astrid": 222,
			"David":  444,
			"Pablo":  666,
		},
		bebidasFav: []string{
			"agua",
			"naranjada",
			"cerveza",
		},
	}

	fmt.Println(s.nombre)
	fmt.Println("\tAmigos:")
	for k, v := range s.amigos {
		fmt.Println("\t\t", k, v)
	}
	fmt.Println("\tBebidas favoritas:")
	for i, v := range s.bebidasFav {
		fmt.Println("\t\t", i, v)
	}
}
