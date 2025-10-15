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
	jsonString := `{"name":"Alice","age":30,"city":"New York"}`
	var person Person
	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	fmt.Println("Person Struct:", person)
}
