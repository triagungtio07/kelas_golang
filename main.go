package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/triagungtio07/golang_fiber/config/database"
	"github.com/triagungtio07/golang_fiber/config/env"
)

func init() {
	env.Load() // first
	database.Load()

}
func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(fmt.Sprintf(":%s", env.AppPort))
}
