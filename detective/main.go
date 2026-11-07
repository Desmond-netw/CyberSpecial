package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {

	// load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
