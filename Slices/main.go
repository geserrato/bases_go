package main

import "fmt"

func main() {
	var arregloCadenas []string
	fmt.Println("Arreglo de cadenas:", arregloCadenas)

	// Inicializando el slice con make
	arregloCadenas = make([]string, 3)

	arregloCadenas[0] = "Hola"
	arregloCadenas[1] = "Mundo"
	arregloCadenas[2] = "!"

	fmt.Println("Arreglo de cadenas después de asignar valores:", arregloCadenas)

	arregloCadenas = append(arregloCadenas, "Platzi", "Nunca pares de aprender")
	fmt.Println("Longitud", len(arregloCadenas))
}
