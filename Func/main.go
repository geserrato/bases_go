package main

import "fmt"

func suma(numero1 int, numero2 int) int {
	return numero1 + numero2
}

func sumaLarga(numero1 int, numero2 int, numero3 int) int {
	resultado := numero1 + numero2 + numero3
	return resultado
}

func main() {

	var numero1, numero2, numero3 int
	fmt.Println("Ingresa el primer número: ")
	fmt.Scanln(&numero1)
	fmt.Println("Ingresa el segundo número: ")
	fmt.Scanln(&numero2)
	fmt.Println("Ingresa el tercer número: ")
	fmt.Scanln(&numero3)

	resultado := suma(numero1, numero2)
	fmt.Println("El resultado de la suma es:", resultado)
	resultadoLargo := sumaLarga(numero1, numero2, numero3)
	fmt.Println("El resultado de la suma larga es:", resultadoLargo)
}
