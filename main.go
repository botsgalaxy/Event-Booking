package main

import (
	"github.com/botsgalaxy/Event-Booking/database"
	"github.com/botsgalaxy/Event-Booking/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	database.InitDB()
}

func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")

}
