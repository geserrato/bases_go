package main

import "fmt"

func suma(numeros ...int) {
	fmt.Println(numeros, "")
	total := 0
	for _, num := range numeros {
		total += num
	}
	fmt.Println("La suma es:", total)
}

func main() {
	suma(1, 2, 3, 4, 5)
	suma(10, 20, 30)

	numeros := []int{100, 200, 300, 400}
	suma(numeros...)
}
