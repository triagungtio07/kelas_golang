package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/triagungtio07/kelas_golang/config/database"
	"github.com/triagungtio07/kelas_golang/handlers"

	"github.com/triagungtio07/kelas_golang/repositories"
	"github.com/triagungtio07/kelas_golang/services"
)

func RegisterUserRoutes(router fiber.Router) {
	service := services.User{
		Repository: repositories.NewUserRepository(database.Db),
	}
	user := handlers.User{
		Service: service,
	}
	router.Post("/users", user.Create)
	router.Post("/users/login", user.Login)
	router.Put("/users/:id", user.Update)
	router.Put("/update-password/:id", user.UpdatePassword)
	router.Get("/users", user.GetPaginated)
	router.Get("/users/:id", user.Get)
	router.Delete("/users/:id", user.Delete)

}
