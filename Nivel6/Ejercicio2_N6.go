package main

import "fmt"

func main() {
	snum := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := foo(snum...)
	fmt.Println(n)

	snum2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n2 := bar(snum2)
	fmt.Println(n2)
}

func foo(xi ...int) int {
	suma := 0
	for _, v := range xi {
		suma += v
	}
	return suma
}

func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
