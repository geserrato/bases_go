package main

import (
	"fmt"
	"time"
)

func main() {
	function("Modo directo")

	go function("Goroutine")
	go func(message string) {
		fmt.Println(message)
	}("Mensaje desde una funcion anonima")

	time.Sleep(1 * time.Second)
	fmt.Println("Finalizando main")
}

func function(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}
