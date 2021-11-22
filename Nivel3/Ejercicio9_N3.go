package main

import "fmt"

func main() {
	deporteFav := "futbol"
	switch deporteFav {
	case "futbol":
		fmt.Println("Ve al estadio")
	case "baloncesto":
		fmt.Println("Ve al coliseo")
	case "natacion":
		fmt.Println("Te quiero ver en las olimpiadas.")
	}
}
