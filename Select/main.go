package main

import "time"

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		channel2 <- "Message from channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-channel1:
			println(msg1)
		case msg2 := <-channel2:
			println(msg2)
		}
	}
}
