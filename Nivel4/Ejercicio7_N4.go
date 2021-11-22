package main

import "fmt"

func main() {
	v1 := []string{"Batman", "Jefe", "Vestido de negro"}
	v2 := []string{"Robin", "Ayudante", "Vestido de colores"}

	fmt.Println(v1)
	fmt.Println(v2)

	result := [][]string{v1, v2}
	fmt.Println(result)

	for i, reg := range result {
		fmt.Println("Registro:", i)
		for j, val := range reg {
			fmt.Printf("\tÍndice: %v\tvalor: %v\n", j, val)
		}
	}
}
