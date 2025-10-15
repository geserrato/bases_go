package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	appName := os.Getenv("APP_NAME")
	fmt.Println("App Name:", appName)
	port := os.Getenv("APP_PORT")
	fmt.Println("Port:", port)
	env := os.Getenv("APP_ENV")
	fmt.Println("Env:", env)
}
