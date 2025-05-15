package controllers

import (
	"github.com/aronipurwanto/go-api-gorm/models"
	"github.com/aronipurwanto/go-api-gorm/services"
	"github.com/aronipurwanto/go-api-gorm/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type EmployeeController struct {
	service  services.EmployeeService
	validate *validator.Validate
}

func NewEmployeeController(service services.EmployeeService) *EmployeeController {
	return &EmployeeController{
		service:  service,
		validate: validator.New(),
	}
}

func (c *EmployeeController) GetAll(ctx *fiber.Ctx) error {
	page, limit := utils.GetPagination(ctx)

	data, total, err := c.service.GetAllPaginated(page, limit)
	if err != nil {
		return utils.ErrorResponse(ctx, 500, "Failed to retrieve data", []utils.ErrorDetail{{Message: err.Error()}})
	}

	meta := utils.Meta{
		Page:  page,
		Limit: limit,
		Total: int(total),
	}

	return utils.ListResponse(ctx, 200, "List retrieved successfully", data, meta)
}

func (c *EmployeeController) GetByID(ctx *fiber.Ctx) error {
	id, err := utils.ParseID(ctx)
	if err != nil {
		return utils.ErrorResponse(ctx, 400, "Invalid ID", []utils.ErrorDetail{{Message: err.Error()}})
	}

	emp, err := c.service.GetByID(id)
	if err != nil {
		return utils.ErrorResponse(ctx, 404, "Employee not found", []utils.ErrorDetail{{Message: err.Error()}})
	}
	return utils.SuccessResponse(ctx, 200, "Data retrieved successfully", emp)
}

func (c *EmployeeController) Create(ctx *fiber.Ctx) error {
	emp, validationErrs, err := utils.BindAndValidate[models.Employee](ctx, c.validate)
	if err != nil {
		return utils.ErrorResponse(ctx, 400, "Invalid input", []utils.ErrorDetail{{Message: err.Error()}})
	}

	if validationErrs != nil {
		return utils.ErrorResponse(ctx, 400, "Validation failed", validationErrs)
	}

	if err = c.service.Create(emp); err != nil {
		return utils.ErrorResponse(ctx, 500, "Failed to create employee", []utils.ErrorDetail{{Message: err.Error()}})
	}

	return utils.SuccessResponse(ctx, 201, "Employee created successfully", emp)
}

func (c *EmployeeController) Update(ctx *fiber.Ctx) error {
	id, err := utils.ParseID(ctx)
	if err != nil {
		return utils.ErrorResponse(ctx, 400, "Invalid ID", []utils.ErrorDetail{{Message: err.Error()}})
	}

	_, err = c.service.GetByID(id)
	if err != nil {
		return utils.ErrorResponse(ctx, 404, "Employee not found", []utils.ErrorDetail{{Message: err.Error()}})
	}

	input, validationErrs, err := utils.BindAndValidate[models.Employee](ctx, c.validate)
	if err != nil {
		return utils.ErrorResponse(ctx, 400, "Invalid input", []utils.ErrorDetail{{Message: err.Error()}})
	}
	if validationErrs != nil {
		return utils.ErrorResponse(ctx, 400, "Validation failed", validationErrs)
	}

	input.EmployeeID = id

	if err := c.service.Update(input); err != nil {
		return utils.ErrorResponse(ctx, 500, "Update failed", []utils.ErrorDetail{{Message: err.Error()}})
	}

	return utils.SuccessResponse(ctx, 200, "Employee updated successfully", input)
}

func (c *EmployeeController) Delete(ctx *fiber.Ctx) error {
	id, err := utils.ParseID(ctx)
	if err != nil {
		return utils.ErrorResponse(ctx, 400, "Invalid ID", []utils.ErrorDetail{{Message: err.Error()}})
	}

	if err = c.service.Delete(id); err != nil {
		return utils.ErrorResponse(ctx, 500, "Delete failed", []utils.ErrorDetail{{Message: err.Error()}})
	}

	return utils.SuccessResponse(ctx, 200, "Employee deleted successfully", nil)
}
