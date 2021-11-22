package main

import "fmt"

func main() {
	an1 := 1988
	for {
		if an1 > 2021 {
			break
		}
		fmt.Println(an1)
		an1++
	}
}
