package main

import "fmt"

func main() {

	func() { //funcion anonima, no tiene identificador, ni nombre
		for i := 0; i <= 100; i++ {
			fmt.Println(i)
		}
	}() //el parentesis porq se ejecuta

	fmt.Println("Listo!")
}
