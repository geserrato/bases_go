package main

import (
	"errors"
	"fmt"
)

var ErrorCoffee = fmt.Errorf("Coffee empty")
var ErrorEnergy = errors.New("Energy empty")

func HacerCafe(args int) error {
	if args == 2 {
		return ErrorCoffee
	} else if args == 4 {
		return ErrorEnergy
	}

	return nil
}

func main() {
	for i := range 5 {
		if err := HacerCafe(i); err != nil {
			if errors.Is(err, ErrorCoffee) {
				fmt.Println("Error de cafe")
			} else if errors.Is(err, ErrorEnergy) {
				fmt.Println("Error de energia")
			} else {
				fmt.Printf("Error desconocido: %v\n", err)
			}
			continue
		}

		fmt.Println("Cafe hecho")
	}
}
