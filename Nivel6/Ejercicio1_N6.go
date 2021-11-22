package main

import "fmt"

func main() {
	n := foo()
	x, s := bar()

	fmt.Println(n, x, s)
}

func foo() int {
	return 14
}

func bar() (int, string) {
	return 1810, "Año independencia del Ecuador"
}
