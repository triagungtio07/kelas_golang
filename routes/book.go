package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/triagungtio07/kelas_golang/config/database"
	"github.com/triagungtio07/kelas_golang/handlers"
	"github.com/triagungtio07/kelas_golang/repositories"
	"github.com/triagungtio07/kelas_golang/services"
)

func RegisterBookRoutes(router fiber.Router) {
	service := services.Book{
		Repository: repositories.NewBookRepository(database.Db),
	}
	book := handlers.Book{
		Service: service,
	}
	router.Post("/books", book.Create)
	router.Put("/books/:id", book.Update)
	router.Get("/books", book.GetPaginated)
	router.Get("/books/:id", book.Get)
	router.Delete("/books/:id", book.Delete)

}
