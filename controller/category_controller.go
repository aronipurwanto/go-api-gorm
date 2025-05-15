package controller

import (
	"github.com/aronipurwanto/go-api-gorm/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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
