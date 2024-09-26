package main

import (
	"github.com/joho/godotenv"
	"github.com/ravanbabayev/nextjs-go-ecommerce/config"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.ConnectDB()
	defer db.Close()

}
