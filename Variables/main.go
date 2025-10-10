package main

import "fmt"

func main() {
	var cadena = "Inicial"

	fmt.Println(cadena)
	var entero, entero2 int = 10, 20
	fmt.Println(entero)
	fmt.Println(entero2)

	var booleano bool = true
	fmt.Println(booleano)

	// No necesito inicializar mi variable cuando la declaro
	// Si no la inicializo, toma el valor cero del tipo de dato
	// El valor cero de un int es 0
	var enteroSimple int
	fmt.Println(enteroSimple)
	enteroSimple = 15
	fmt.Println(enteroSimple)

	// := Operador de declaración corta de variables inferencia de tipo
	fruta := "Mango"
	fmt.Println(fruta)
}
