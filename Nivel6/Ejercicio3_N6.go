package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Hola Mundo")
}

func foo() {
	defer func() {
		fmt.Println("Foo diferida se ejecuta")
	}()
	fmt.Println("Foo se ejecuto")
}
