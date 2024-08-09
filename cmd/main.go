package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	routes "hastane-takip/app"
	"hastane-takip/app/di"
	"hastane-takip/internal/db"
	"hastane-takip/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	envPath := filepath.Join("..", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db.GetDB()

	utils.InitRedis(
		os.Getenv("REDIS_ADDR"),
		os.Getenv("REDIS_PASSWORD"),
		0,
	)
	fmt.Println("Redis connected")

	app := fiber.New()

	container := di.BuildContainer()

	if err := routes.SetupRoutes(app, container); err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server is running on port %s", port)
	app.Listen(":3000")
}
