package controllers

import (
	"github.com/aronipurwanto/go-api-gorm/models"
	"github.com/aronipurwanto/go-api-gorm/services"
	"github.com/aronipurwanto/go-api-gorm/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	service  services.CategoryService
	validate *validator.Validate
}

func NewCategoryController(service services.CategoryService) *CategoryController {
	return &CategoryController{
		service:  service,
		validate: validator.New(),
	}
}

func (c *CategoryController) GetAll(ctx *fiber.Ctx) error {
	data, err := c.service.GetAll()
	if err != nil {
		return utils.ErrorResponse(ctx, 500, "Failed to fetch categories", []utils.ErrorDetail{{Message: err.Error()}})
	}
	return utils.ListResponse(ctx, 200, "Categories fetched successfully", data, utils.Meta{})
}

func (c *CategoryController) GetByID(ctx *fiber.Ctx) error {
	id, err := utils.ParseID(ctx)
	if err != nil {
		return utils.ErrorResponse(ctx, 400, "Invalid ID", []utils.ErrorDetail{{Message: err.Error()}})
	}

	data, err := c.service.GetByID(id)
	if err != nil {
		return utils.ErrorResponse(ctx, 404, "Category not found", []utils.ErrorDetail{{Message: err.Error()}})
	}
	return utils.SuccessResponse(ctx, 200, "Category fetched", data)
}

func (c *CategoryController) Create(ctx *fiber.Ctx) error {
	cat, validationErrs, err := utils.BindAndValidate[models.Category](ctx, c.validate)
	if err != nil {
		return utils.ErrorResponse(ctx, 400, "Invalid input", []utils.ErrorDetail{{Message: err.Error()}})
	}

	if validationErrs != nil {
		return utils.ErrorResponse(ctx, 400, "Validation failed", validationErrs)
	}

	create, err := c.service.Create(cat)
	if err != nil {
		return utils.ErrorResponse(ctx, 500, "Create failed", []utils.ErrorDetail{{Message: err.Error()}})
	}

	return utils.SuccessResponse(ctx, 201, "Category created", create)
}

func (c *CategoryController) Update(ctx *fiber.Ctx) error {
	id, err := utils.ParseID(ctx)
	if err != nil {
		return utils.ErrorResponse(ctx, 400, "Invalid ID", []utils.ErrorDetail{{Message: err.Error()}})
	}

	_, err = c.service.GetByID(id)
	if err != nil {
		return utils.ErrorResponse(ctx, 404, "Category not found", []utils.ErrorDetail{{Message: err.Error()}})
	}

	// Parse + Validate
	input, validationErrs, err := utils.BindAndValidate[models.Category](ctx, c.validate)
	if err != nil {
		return utils.ErrorResponse(ctx, 400, "Invalid input", []utils.ErrorDetail{{Message: err.Error()}})
	}
	if validationErrs != nil {
		return utils.ErrorResponse(ctx, 400, "Validation failed", validationErrs)
	}

	// Force ID from URL to match body
	input.CategoryID = id
	update, err := c.service.Update(input)
	if err != nil {
		return utils.ErrorResponse(ctx, 500, "Update failed", []utils.ErrorDetail{{Message: err.Error()}})
	}

	return utils.SuccessResponse(ctx, 200, "Category updated", update)
}

func (c *CategoryController) Delete(ctx *fiber.Ctx) error {
	id, err := utils.ParseID(ctx)
	if err != nil {
		return utils.ErrorResponse(ctx, 400, "Invalid ID", []utils.ErrorDetail{{Message: err.Error()}})
	}

	if err = c.service.Delete(id); err != nil {
		return utils.ErrorResponse(ctx, 500, "Delete failed", []utils.ErrorDetail{{Message: err.Error()}})
	}

	return utils.SuccessResponse(ctx, 200, "Category deleted", nil)
}
