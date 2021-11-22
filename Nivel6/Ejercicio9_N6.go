package main

import "fmt"

func main() {

	r1 := func(v1 []int) int {
		if len(v1) == 0 {
			return 0
		}
		if len(v1) == 1 {
			return v1[0]
		}
		return v1[0] + v1[len(v1)-1]
	}

	x := foo(r1, []int{1, 2, 3, 4, 5, 6})
	fmt.Println(x)
}

func foo(f func(v1 []int) int, v2 []int) int {
	n := f(v2)
	n++
	return n
}
