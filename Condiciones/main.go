package main

import "fmt"

func main() {
	nombre := "Gerardo"
	edad := 34
	if nombre == "Gerardo" && edad >= 18 {
		fmt.Println("Eres mayor de edad y te llamas Gerardo")
	} else {
		fmt.Println("Hola desconocido")
	}

	if 8%2 == 0 {
		fmt.Println("8 es un número par")
	} else {
		fmt.Println("8 es un número impar")
	}

	if numero := 9; numero < 0 {
		fmt.Println("El número es negativo")
	} else if numero < 10 {
		fmt.Println("El número es menor que 10 pero no es negativo")
	} else {
		fmt.Println("El número es mayor o igual que 10")
	}
}
