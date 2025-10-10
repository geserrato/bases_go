package main

import "fmt"

func main() {
	var numero int = 20
	numero2 := 30
	fmt.Println(numero, numero2)

	var numeroEntero = 20
	var numeroDecimal = 20.5
	resultado := numeroEntero + int(numeroDecimal)
	fmt.Println(resultado)

	var nombre string = "Gerardo"
	var apellido string = "Estrella"

	nombreCompleto := nombre + " " + apellido
	fmt.Println(nombreCompleto)
}
