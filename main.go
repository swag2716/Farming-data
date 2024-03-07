package main

import (
	"log"
	"os"

	"Farming_data/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Unable to load environment variables")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "9000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.FarmingRoutes(router)

	router.Run(":" + port)
}
