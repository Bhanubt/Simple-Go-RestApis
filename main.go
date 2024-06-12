package main

import (
	"pvmsproject/db"
	"pvmsproject/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//Creating Database
	db.Connect()

	//Creating Server
	server := gin.Default()

	//Sending the Server to the Routes
	routes.RegisterRoutes(server)

	//Running the Server on localhost:8080
	server.Run(":8080")
}
