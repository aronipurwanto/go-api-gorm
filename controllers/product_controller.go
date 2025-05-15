package controllers

import (
	"github.com/aronipurwanto/go-api-gorm/models"
	"github.com/aronipurwanto/go-api-gorm/services"
	"github.com/aronipurwanto/go-api-gorm/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	service  services.ProductService
	validate *validator.Validate
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{
		service:  service,
		validate: validator.New(),
	}
}

func (c *ProductController) GetAll(ctx *fiber.Ctx) error {
	page, limit := utils.GetPagination(ctx)
	//page := ctx.Locals("page").(int)
	//limit := ctx.Locals("limit").(int)

	products, total, err := c.service.GetAllPaginated(page, limit)
	if err != nil {
		return utils.ErrorResponse(ctx, 500, "Failed to retrieve products", []utils.ErrorDetail{{Message: err.Error()}})
	}

	meta := utils.Meta{
		Page:  page,
		Limit: limit,
		Total: int(total),
	}

	return utils.ListResponse(ctx, 200, "Products retrieved", products, meta)
}

func (c *ProductController) Search(ctx *fiber.Ctx) error {
	name := ctx.Query("name", "")
	if name == "" {
		return utils.ErrorResponse(ctx, 400, "Missing query parameter: name", nil)
	}

	//page := ctx.QueryInt("page", 1)
	//limit := ctx.QueryInt("limit", 10)
	page, limit := utils.GetPagination(ctx)

	products, total, err := c.service.SearchByName(name, page, limit)
	if err != nil {
		return utils.ErrorResponse(ctx, 500, "Search failed", []utils.ErrorDetail{{Message: err.Error()}})
	}

	meta := utils.Meta{
		Page:  page,
		Limit: limit,
		Total: int(total),
	}

	return utils.ListResponse(ctx, 200, "Search results", products, meta)
}

func (c *ProductController) GetByID(ctx *fiber.Ctx) error {
	id, err := utils.ParseID(ctx)
	if err != nil {
		return utils.ErrorResponse(ctx, 400, "Invalid ID", []utils.ErrorDetail{{Message: err.Error()}})
	}

	data, err := c.service.GetByID(id)
	if err != nil {
		return utils.ErrorResponse(ctx, 404, "Product not found", []utils.ErrorDetail{{Message: err.Error()}})
	}
	return utils.SuccessResponse(ctx, 200, "Product retrieved", data)
}

func (c *ProductController) Create(ctx *fiber.Ctx) error {
	prod, validationErrs, err := utils.BindAndValidate[models.Product](ctx, c.validate)
	if err != nil {
		return utils.ErrorResponse(ctx, 400, "Invalid input", []utils.ErrorDetail{{Message: err.Error()}})
	}

	if validationErrs != nil {
		return utils.ErrorResponse(ctx, 400, "Validation failed", validationErrs)
	}

	created, err := c.service.Create(prod)
	if err != nil {
		return utils.ErrorResponse(ctx, 500, "Create failed", []utils.ErrorDetail{{Message: err.Error()}})
	}
	return utils.SuccessResponse(ctx, 201, "Product created", created)
}

func (c *ProductController) Update(ctx *fiber.Ctx) error {
	id, err := utils.ParseID(ctx)
	if err != nil {
		return utils.ErrorResponse(ctx, 400, "Invalid ID", []utils.ErrorDetail{{Message: err.Error()}})
	}

	_, err = c.service.GetByID(id)
	if err != nil {
		return utils.ErrorResponse(ctx, 404, "Product not found", []utils.ErrorDetail{{Message: err.Error()}})
	}

	prod, validationErrs, err := utils.BindAndValidate[models.Product](ctx, c.validate)
	if err != nil {
		return utils.ErrorResponse(ctx, 400, "Invalid input", []utils.ErrorDetail{{Message: err.Error()}})
	}

	if validationErrs != nil {
		return utils.ErrorResponse(ctx, 400, "Validation failed", validationErrs)
	}

	updated, err := c.service.Update(prod)
	if err != nil {
		return utils.ErrorResponse(ctx, 500, "Update failed", []utils.ErrorDetail{{Message: err.Error()}})
	}
	return utils.SuccessResponse(ctx, 200, "Product updated", updated)
}

func (c *ProductController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Locals("id").(int)

	if err := c.service.Delete(id); err != nil {
		return utils.ErrorResponse(ctx, 500, "Delete failed", []utils.ErrorDetail{{Message: err.Error()}})
	}
	return utils.SuccessResponse(ctx, 200, "Product deleted", nil)
}
