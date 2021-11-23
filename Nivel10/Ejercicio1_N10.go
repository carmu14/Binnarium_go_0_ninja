package main

import "fmt"

func main() {
	//c := make(chan int) con gorutina
	c := make(chan int, 1) //con buffer

	c <- 14

	//con gorutina
	//go func() {
	//	c <- 14
	//}()

	fmt.Println(<-c)
}
