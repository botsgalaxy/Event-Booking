package main

import (
	"log"

	"github.com/botsgalaxy/Event-Booking/database"
	"github.com/botsgalaxy/Event-Booking/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	database.InitDB()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")

}
