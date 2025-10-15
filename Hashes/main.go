package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	baseString := "Hello, World!"
	hash := sha256.New()
	hash.Write([]byte(baseString))
	hashedString := hash.Sum(nil)
	fmt.Println(hashedString)
	fmt.Printf("%x\n", hashedString)
}
