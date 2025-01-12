package main

import (
	"os"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"Go_auth_app/database"
	"Go_auth_app/routes"
)

func main() {
	// Loading environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connecting to MongoDB
	database.Connect()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	// Initializing Fiber app
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	// Setting up routes
	routes.SetupRoutes(app)

	// Starting the server
	log.Fatal(app.Listen(":" + PORT))
}
