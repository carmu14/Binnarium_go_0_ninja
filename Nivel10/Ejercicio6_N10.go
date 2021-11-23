package main

import "fmt"

func main() {
	c := make(chan int) //declarar canal

	go func() { //gorutina para enviar valores al canal
		for i := 0; i < 100; i++ {
			c <- i //enviamos al canal c el valor de i
		}
		close(c) //cerramos el canal para evitar que se bloquee
	}()

	for v := range c { //imprimir en pantalla los 100 numeros del canal
		fmt.Println(v)
	}

	fmt.Println("A punto de finalizar.")
}
