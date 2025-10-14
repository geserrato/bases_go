package main

import "fmt"

func main() {
	// chan - channels
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	message := <-messages
	fmt.Println(message)

	go func() { messages <- "pong" }()

	secondMessage := <-messages
	fmt.Println(secondMessage)
}
