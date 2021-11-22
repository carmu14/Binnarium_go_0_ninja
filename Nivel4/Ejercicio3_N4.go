package main

import "fmt"

func main() {
	x := []int{14, 20, 54, 32, 66, 73, 84, 99, 17, 41}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
	fmt.Println(x)
}
