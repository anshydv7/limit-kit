package main

import (
	"auth-service/base/config"
	"auth-service/base/database"
	"auth-service/routers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()

	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := database.InitiateDatabase(); err != nil {
		log.Panic("error while initiating database", err)
	}
	log.Println("database initiated")

	routers.SetupRouter(engine)

	engine.Run(":8000")
}
