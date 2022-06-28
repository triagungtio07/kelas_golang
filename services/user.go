package services

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/copier"
	"github.com/triagungtio07/kelas_golang/models"
	"github.com/triagungtio07/kelas_golang/repositories"
	"github.com/triagungtio07/kelas_golang/utils"
)

type User struct {
	Repository repositories.UserRepository
}

func (s User) GetPaginated(paginator utils.Paginator) (utils.Paginator, *models.Error) {
	result, err := s.Repository.FindPaginated(paginator)
	if err != nil {
		return result, &models.Error{
			Code: fiber.StatusInternalServerError,
		}
	}

	return result, nil
}

func (s User) Get(id uint64) (models.User, *models.Error) {
	result, err := s.Repository.Find(id)
	if err != nil {
		return result, &models.Error{
			Message: err.Error(),
			Code:    fiber.StatusNotFound,
		}
	}

	return result, nil
}

func (s User) Create(user models.CreateUser) (models.User, error) {
	model := models.User{}
	copier.Copy(&model, &user)
	model.Password = utils.EncodePassword(model.Password)
	err := s.Repository.Save(&model)
	if err != nil {
		return model, &fiber.Error{
			Code:    500,
			Message: "Internal Server Error",
		}
	}

	return model, nil
}

func (s User) Update(user models.UpdateUser) (models.User, error) {
	model, err := s.Repository.Find(user.ID)
	if err != nil {
		return model, &fiber.Error{
			Code:    404,
			Message: "Not Found",
		}
	}

	copier.Copy(&model, &user)
	err = s.Repository.Save(&model)
	if err != nil {
		return model, &fiber.Error{
			Code:    500,
			Message: "Internal Server Error",
		}
	}

	return model, nil
}

func (s User) UpdatePassword(input models.UpdatePassword) (models.User, error) {
	model, err := s.Repository.Find(input.ID)
	if err != nil {
		return model, &fiber.Error{
			Message: fiber.ErrNotFound.Message,
			Code:    fiber.ErrNotFound.Code,
		}
	}

	if !utils.ValidatePassword(model.Password, input.OldPassword) {
		return model, &fiber.Error{
			Message: "old password not match",
			Code:    fiber.StatusBadRequest,
		}
	}

	model.Password = utils.EncodePassword(input.Password)
	err = s.Repository.Save(&model)
	if err != nil {
		return model, &fiber.Error{
			Message: fiber.ErrInternalServerError.Message,
			Code:    fiber.ErrInternalServerError.Code,
		}
	}

	return model, nil
}

func (s User) ValidateLogin(login models.Login) (models.User, error) {
	result, err := s.Repository.FindByEmail(login.Email)
	if err != nil {
		return models.User{}, &fiber.Error{
			Message: "user not found or password not match",
			Code:    fiber.StatusBadRequest,
		}
	}

	if !utils.ValidatePassword(result.Password, login.Password) {
		return models.User{}, &fiber.Error{
			Message: "user not found or password not match",
			Code:    fiber.StatusBadRequest,
		}
	}

	return result, nil
}

func (s User) Delete(id uint64) *models.Error {
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
