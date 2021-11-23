//Paquete para entender la suma de dos numeros
package main

/*
	Nos permite sumar dos numeros enteros
*/
import "fmt"

func main() {
	result := sumar(5, 2)
	fmt.Printf("El resultado de la suma es: %d", result)

}

func sumar(op1 int, op2 int) int {
	return op1 - op2
}
