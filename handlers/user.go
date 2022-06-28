package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/triagungtio07/kelas_golang/models"
	"github.com/triagungtio07/kelas_golang/services"
	"github.com/triagungtio07/kelas_golang/utils"
)

type User struct {
	Service services.User
}

func (c User) Create(ctx *fiber.Ctx) error {
	form := models.CreateUser{}
	ctx.BodyParser(&form)

	messages, err := utils.Validate(form)
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"message": messages,
			"error":   err,
		})

		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	model, err := c.Service.Create(form)
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"model": model,
			"error": err,
		})

		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	ctx.JSON(model)

	return ctx.SendStatus(fiber.StatusCreated)
}

func (c User) Update(ctx *fiber.Ctx) error {
	form := models.UpdateUser{}
	ctx.BodyParser(&form)

	messages, err := utils.Validate(form)
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"message": messages,
			"error":   err,
		})

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	idInt, _ := strconv.Atoi(ctx.Params("id"))
	form.ID = uint64(idInt)
	model, err := c.Service.Update(form)
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"model": model,
			"error": err,
		})

		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	ctx.JSON(model)

	return ctx.SendStatus(fiber.StatusCreated)

}

func (c User) UpdatePassword(ctx *fiber.Ctx) error {
	form := models.UpdatePassword{}
	ctx.BodyParser(&form)

	messages, err := utils.Validate(form)
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"message": messages,
			"error":   err,
		})

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	idInt, _ := strconv.Atoi(ctx.Params("id"))
	form.ID = uint64(idInt)
	model, err := c.Service.UpdatePassword(form)
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"model": model,
			"error": err,
		})

		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	ctx.JSON(model)

	return ctx.SendStatus(fiber.StatusCreated)
}

func (c User) Login(ctx *fiber.Ctx) error {
	form := models.Login{}
	ctx.BodyParser(&form)
	messages, err := utils.Validate(form)
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"message": messages,
		})

		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	result, status := c.Service.ValidateLogin(form)
	if status != nil {
		ctx.JSON(map[string]string{
			"message": "err",
		})

		return ctx.SendStatus(200)
	}

	claims := jwt.MapClaims{}
	claims["email"] = result.Email
	ctx.JSON(map[string]string{
		"token": utils.CreateJwtToken(claims),
	})
	return ctx.SendStatus(fiber.StatusOK)
}

func (c User) Delete(ctx *fiber.Ctx) error {
	idInt, _ := strconv.Atoi(ctx.Params("id"))
	err := c.Service.Delete(uint64(idInt))
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Message,
		})

		return ctx.SendStatus(err.Code)
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c User) Get(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	user, err := c.Service.Get(uint64(id))
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Message,
		})

		return ctx.SendStatus(err.Code)
	}

	return ctx.JSON(user)
}

func (c User) GetPaginated(ctx *fiber.Ctx) error {
	users, err := c.Service.GetPaginated(*utils.NewPaginator(ctx))
	if err != nil {
		ctx.JSON(map[string]string{
			"message": "unable to get all users",
		})

		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(users)
}
