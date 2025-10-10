package main

import (
	"fmt"
	"os"
)

func main() {
	enVar := os.Getenv("HOME")
	if enVar == "" {
		fmt.Println("La variable de entorno HOME no está definida.")
	} else {
		fmt.Println("El valor de la variable de entorno HOME es:", enVar)
	}

	file, err := os.Create("archivo.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()
}
