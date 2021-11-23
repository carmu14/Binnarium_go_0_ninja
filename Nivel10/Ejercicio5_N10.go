//Resolver ejercicio 4 en https://go.dev/play/p/YHOMV9NYKK
package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 14
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
