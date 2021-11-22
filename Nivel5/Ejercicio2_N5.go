package main

import "fmt"

type persona struct {
	nombre     string
	apellido   string
	saboresFav []string
}

func main() {
	p1 := persona{
		nombre:   "Carlos",
		apellido: "Munoz",
		saboresFav: []string{
			"vainilla",
			"chicle",
			"menta",
		},
	}
	p2 := persona{
		nombre:   "Maria",
		apellido: "Palacios",
		saboresFav: []string{
			"fresa",
			"chocolate",
			"limón",
		},
	}

	m1 := map[string]persona{
		p1.apellido: p1,
		p2.apellido: p2,
	}

	//for lla, v := range m1 {
	//fmt.Println(lla)
	//	fmt.Println(v)
	//}

	for _, v := range m1 {
		fmt.Println(v.nombre)
		fmt.Println(v.apellido)
		for i, v := range v.saboresFav {
			fmt.Println(" ", i, v)
		}
		fmt.Println("------------")
	}
}
