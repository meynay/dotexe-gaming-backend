package main

import (
	"log"

	"store/infrastracture"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Println("No .env file found")
	}
	infrastracture.StartServer()
}
