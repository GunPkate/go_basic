package routes

import (
	"fiber-app/handlers"

	"github.com/gofiber/fiber/v2"
)

// Setup registers all application routes on the given Fiber app.
func Setup(app *fiber.App) {
	userHandler := handlers.NewUserHandler()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	api := app.Group("/api/v1")

	users := api.Group("/users")
	users.Get("/", userHandler.GetUsers)
	users.Get("/:id", userHandler.GetUser)
	users.Post("/", userHandler.CreateUser)
	users.Put("/:id", userHandler.UpdateUser)
	users.Delete("/:id", userHandler.DeleteUser)
}
