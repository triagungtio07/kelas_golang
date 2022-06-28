package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/triagungtio07/kelas_golang/config/database"
	"github.com/triagungtio07/kelas_golang/config/env"
	"github.com/triagungtio07/kelas_golang/routes"
)

func init() {
	env.Load() // first
	database.Load()

}
func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	routes.RegisterRoutes(app)

	app.Listen(fmt.Sprintf(":%s", env.AppPort))
}
