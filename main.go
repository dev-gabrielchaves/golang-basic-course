package main

import (
	"log"

	"github.com/dev-gabrielchaves/golang-learning/database"
	"github.com/dev-gabrielchaves/golang-learning/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connect to MongoDB
	database.ConnectDB()

	// Initialize Fiber
	app := fiber.New()

	// Routes
	app.Get("/users", handlers.GetUsers)
	app.Get("/users/:id", handlers.GetUser)
	app.Post("/users", handlers.CreateUser)
	app.Put("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.DeleteUser)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
