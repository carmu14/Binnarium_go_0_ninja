package main

import (
	"fmt"

	"github.com/carmu14/go_binnarium/Nivel13/Ejercicio3/matematicas"
)

func main() {
	xxi := gen()
	for _, v := range xxi {
		fmt.Println(matematicas.PromedioCentrado(v))
	}
}

func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	e := [][]int{a, b, c, d}
	return e
}
