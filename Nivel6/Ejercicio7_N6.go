package main

import "fmt"

var a int
var b func()

func main() {

	f := func() {
		for i := 0; i <= 2; i++ {
			fmt.Println(i)
		}
	}

	f()
	fmt.Printf("%T\n", f)

	fmt.Println(a)
	fmt.Printf("%T\n", b)

	b := f
	b()
	fmt.Printf("Esta es b %T\n", b)

	fmt.Println("Listo!")
}
