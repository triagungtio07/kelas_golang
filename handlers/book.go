package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/triagungtio07/kelas_golang/models"
	"github.com/triagungtio07/kelas_golang/services"
	"github.com/triagungtio07/kelas_golang/utils"
)

type Book struct {
	Service services.Book
}

func (c Book) Create(ctx *fiber.Ctx) error {
	form := models.CreateBook{}
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

func (c Book) Update(ctx *fiber.Ctx) error {
	form := models.UpdateBook{}
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

func (c Book) Delete(ctx *fiber.Ctx) error {
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

func (c Book) Get(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	Book, err := c.Service.Get(uint64(id))
	if err != nil {
		ctx.JSON(map[string]string{
			"message": err.Message,
		})

		return ctx.SendStatus(err.Code)
	}

	return ctx.JSON(Book)
}

func (c Book) GetPaginated(ctx *fiber.Ctx) error {
	Books, err := c.Service.GetPaginated(*utils.NewPaginator(ctx))
	if err != nil {
		ctx.JSON(map[string]string{
			"message": "unable to get all Books",
		})

		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(Books)
}
