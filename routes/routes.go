package routes

import (
	"Go_auth_app/controllers"

	"github.com/gofiber/fiber/v2"
)
// setting up the routes
func SetupRoutes(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
}
