package main

import "fmt"

func main() {
	m := "Superman"
	if m == "Iroman" {
		fmt.Println(m)
	} else if m == "Hulk" {
		fmt.Println("Hulk destruye")
	} else {
		fmt.Println("Esto no es marvel")
	}
}
