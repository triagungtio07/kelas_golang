package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	RegisterUserRoutes(router)
	RegisterBookRoutes(router)

}
