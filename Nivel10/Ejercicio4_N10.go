//Resolver ejercicio 4 en https://go.dev/play/p/9rDOEvNeBjc
package main

import (
	"fmt"
)

func main() {
	salir := make(chan int)
	c := gen(salir)

	recibir(c, salir)

	fmt.Println("A punto de finalizar.")
}

//sol
func recibir(c, salir <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-salir:
			return
		}
	}
}

func gen(salir chan<- int) <-chan int {
	c := make(chan int)

	go func() { //sol
		for i := 0; i < 10; i++ {
			c <- i
		}
		salir <- 1
		close(c)
	}()

	return c
}
