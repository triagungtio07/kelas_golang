package services

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/copier"
	"github.com/triagungtio07/kelas_golang/models"
	"github.com/triagungtio07/kelas_golang/repositories"
	"github.com/triagungtio07/kelas_golang/utils"
)

type Book struct {
	Repository repositories.BookRepository
}

func (s Book) GetPaginated(paginator utils.Paginator) (utils.Paginator, *models.Error) {
	result, err := s.Repository.FindPaginated(paginator)
	if err != nil {
		return result, &models.Error{
			Code: fiber.StatusInternalServerError,
		}
	}

	return result, nil
}

func (s Book) Get(id uint64) (models.Book, *models.Error) {
	result, err := s.Repository.Find(id)
	if err != nil {
		return result, &models.Error{
			Message: err.Error(),
			Code:    fiber.StatusNotFound,
		}
	}

	return result, nil
}

func (s Book) Create(Book models.CreateBook) (models.Book, error) {
	model := models.Book{}
	copier.Copy(&model, &Book)
	err := s.Repository.Save(&model)
	if err != nil {
		return model, &fiber.Error{
			Message: fiber.ErrInternalServerError.Message,
			Code:    fiber.ErrInternalServerError.Code,
		}
	}

	return model, nil
}

func (s Book) Update(Book models.UpdateBook) (models.Book, error) {
	model, err := s.Repository.Find(Book.ID)
	if err != nil {
		return model, &fiber.Error{
			Code:    404,
			Message: "Not Found",
		}
	}

	copier.Copy(&model, &Book)
	err = s.Repository.Save(&model)
	if err != nil {
		return model, &fiber.Error{
			Code:    500,
			Message: "Internal Server Error",
		}
	}

	return model, nil
}

func (s Book) Delete(id uint64) *models.Error {
	model, err := s.Repository.Find(id)
	if err != nil {
		return &models.Error{
			Message: err.Error(),
			Code:    fiber.StatusNotFound,
		}
	}

	err = s.Repository.Delete(&model)
	if err != nil {
		return &models.Error{
			Code: fiber.StatusInternalServerError,
		}
	}

	return nil
}
