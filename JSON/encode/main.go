package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	person := Person{Name: "Alice", Age: 30, City: "New York"}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println("JSON Data: ", string(jsonData))
}
