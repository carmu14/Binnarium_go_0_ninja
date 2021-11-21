package main

import "fmt"

type num int
var a num

func main() {
	fmt.Println(a)
	fmt.Printf("El tipo de x es: %T\n", a)

	a = 42
	fmt.Println(a)
}