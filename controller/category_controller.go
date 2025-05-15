package controller

import (
	"github.com/aronipurwanto/go-api-gorm/models"
	"github.com/aronipurwanto/go-api-gorm/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type CategoryController struct {
	service  services.CategoryService
	validate *validator.Validate
}

func NewCategoryController(service services.CategoryService) CategoryController {
	return CategoryController{
		service:  service,
		validate: validator.New(),
	}
}

func (c CategoryController) GetAll(ctx *fiber.Ctx) error {
	categories, err := c.service.GetAll()
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Fail to get all categories",
			"error":   err,
			"data":    nil,
		})
	}
	return ctx.JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"data":    categories,
		"message": "Success",
	})
}

func (c CategoryController) GetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Fail to convert id to int",
			"error":   err,
			"data":    nil,
		})
	}

	cat, err := c.service.GetByID(atoi)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Something went wrong",
			"error":   err,
			"data":    nil,
		})
	}
	return ctx.JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Success",
		"data":    cat,
	})
}

func (c CategoryController) Create(ctx *fiber.Ctx) error {
	var cat models.Category
	err := ctx.BodyParser(cat)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Fail to parse body",
			"error":   err,
		})
	}

	err = c.validate.Struct(&cat)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Fail to validate body",
			"error":   err,
		})
	}

	create, err := c.service.Create(cat)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Something went wrong",
			"error":   err,
		})
	}

	return ctx.JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Success",
		"data":    create,
	})
}

func (c CategoryController) Update(ctx *fiber.Ctx) error {
	var cat models.Category
	id := ctx.Params("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Fail to convert id to int",
			"error":   err,
			"data":    nil,
		})
	}

	err = ctx.BodyParser(cat)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Fail to parse body",
			"error":   err,
		})
	}

	err = c.validate.Struct(&cat)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Fail to validate body",
			"error":   err,
		})
	}

	_, err = c.service.GetByID(atoi)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Something went wrong",
			"error":   err,
			"data":    nil,
		})
	}

	update, err := c.service.Update(cat)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Something went wrong",
			"error":   err,
			"data":    nil,
		})
	}

	return ctx.JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Success",
		"data":    update,
	})
}

func (c CategoryController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Fail to convert id to int",
			"error":   err,
			"data":    nil,
		})
	}

	err = c.service.Delete(atoi)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Something went wrong",
			"error":   err,
			"data":    nil,
		})
	}

	return ctx.JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Success",
		"data":    atoi,
	})
}
