package main

import "fmt"

func main() {
	x := []int{14, 20, 54, 32, 66, 73, 84, 99, 17, 41}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
