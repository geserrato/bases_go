package main

import "fmt"

func main() {
	numeros := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numeros)
	modificarArrays(&numeros)
	fmt.Println(numeros)
}

func modificarArrays(array *[5]int) {
	(*array)[0] = 42
}
