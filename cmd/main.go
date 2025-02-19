package main

import (
	"fileserver/config"
	"fileserver/routes"
	"fileserver/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	config.LoadConfig()
	services.InitMinio()

	router := gin.Default()

	routes.Fileserver(router)

	if err := router.Run(":9090"); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
