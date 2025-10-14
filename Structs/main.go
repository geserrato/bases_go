package main

import "fmt"

type persona struct {
	name string
	age  int
}

func newPerson(name string, age int) *persona {
	newGuess := persona{name: name}
	newGuess.age = 42

	return &newGuess
}

func main() {
	fmt.Println(persona{"Bob", 20})
	fmt.Println(newPerson("Roberto", 80))
}
