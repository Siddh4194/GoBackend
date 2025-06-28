package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/siddhant/shetkariBasketGo/config"
	"github.com/siddhant/shetkariBasketGo/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	fmt.Println("🚀 Starting Shetkari Basket...")
	db := config.GetDB()

	r := router.SetupRoutes(db)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}