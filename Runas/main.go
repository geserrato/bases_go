package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const saludo = "áéíóú"
	const saludo2 = "Hola"

	fmt.Println(saludo)
	fmt.Println(len(saludo))
	fmt.Println("El saludo es: ", saludo2)
	fmt.Println(len(saludo2))

	for i, _ := range saludo {
		fmt.Printf("%x", saludo[i])
	}

	fmt.Println("Conteo de runas: ", utf8.RuneCountInString(saludo))

	for idx, valorRuna := range saludo {
		fmt.Printf("%#U comienza en %d\n", valorRuna, idx)
	}
}
