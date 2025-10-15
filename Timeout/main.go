package main

import (
	"fmt"
	"time"
)

func main() {
	contador1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		contador1 <- "Proceso terminado"
	}()

	select {
	case resultado := <-contador1:
		fmt.Println(resultado)
	case <-time.After(1 * time.Second):
		fmt.Println("Tiempo de espera agotado")
	}

	contador2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		contador2 <- "Proceso terminado"
	}()

	select {
	case resultado := <-contador1:
		fmt.Println(resultado)
	case <-time.After(3 * time.Second):
		fmt.Println("Tiempo de espera agotado")
	}
}
