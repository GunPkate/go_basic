package main

import (
	"log"

	"fiber-app/config"
	"fiber-app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg := config.Load()

	app := fiber.New(fiber.Config{
		AppName: "fiber-app",
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	routes.Setup(app)

	log.Fatal(app.Listen(":" + cfg.Port))
}
