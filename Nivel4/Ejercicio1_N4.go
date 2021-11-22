package main

import "fmt"

func main() {
	x := [5]int{14, 20, 50, 70, 90}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
