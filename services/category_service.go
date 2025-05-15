package services

import (
	"github.com/aronipurwanto/go-api-gorm/models"
	"github.com/aronipurwanto/go-api-gorm/repositories"
)

type CategoryService interface {
	GetAll() ([]models.Category, error)
	GetByID(id int) (models.Category, error)
	Create(category models.Category) (models.Category, error)
	Update(category models.Category) (models.Category, error)
	Delete(id int) error
}

type categoryServiceImpl struct {
	repo repositories.CategoryRepo
}

func NewCategoryService(repo repositories.CategoryRepo) CategoryService {
	return &categoryServiceImpl{
		repo: repo,
	}
}
