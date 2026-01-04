package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/realtouseef/pulseplus/handlers"
)

func main() {

	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/ping", handlers.Ping)

	app.Listen(":3000")
}
