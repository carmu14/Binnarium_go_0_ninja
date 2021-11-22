package main

import "fmt"

func main() {
	a := (50 == 50)
	b := (50 <= 63)
	c := (50 >= 13)
	d := (50 != 43)
	e := (50 < 75)
	f := (50 > 47)
	fmt.Println(a, b, c, d, e, f)
}
