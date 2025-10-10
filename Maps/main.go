package main

import (
	"fmt"
	"maps"
)

func main() {
	mapa := make(map[string]int)

	mapa["Gera"] = 4
	mapa["Estrella"] = 8

	fmt.Println("Mapa", mapa)

	version1 := mapa["Estrella"]
	fmt.Println("Version 1", version1)

	version2 := mapa["Gera"]
	fmt.Println("Version 2", version2)

	fmt.Println("Longitud del mapa:", len(mapa))

	_, dato := mapa["Gera"]
	fmt.Println("Dato:", dato)

	delete(mapa, "Gera")
	fmt.Println("Mapa despues de borrar a Gera:", mapa)

	clear(mapa)

	fmt.Println("Mapa despues de limpiar:", mapa)

	map1 := map[string]int{"Gera": 4, "Estrella": 8}
	map2 := map[string]int{"Gera": 4, "Estrella": 8}

	if maps.Equal(map1, map2) {
		fmt.Println("Los mapas son iguales")
	}
}
