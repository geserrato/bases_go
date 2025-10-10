package main

import "fmt"

func main() {
	var arreglos [5]int
	fmt.Println("Arreglo completo:", arreglos)

	arreglos[4] = 100
	fmt.Println("Arreglo en la posicion 4", arreglos[4])
	fmt.Println("Tamaño del arreglo:", len(arreglos))

	lista := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Lista completa:", lista)
	fmt.Println("Tamaño de la lista:", len(lista))

	listTwo := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Lista dos completa:", listTwo)
	fmt.Println("Tamaño de la lista establecido por GO:", len(listTwo))
}
