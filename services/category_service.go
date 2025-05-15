package services

import (
	"github.com/aronipurwanto/go-api-gorm/models"
	"github.com/aronipurwanto/go-api-gorm/repositories"
)

type CategoryService interface {
	GetAll() ([]models.Category, error)
	GetByID(id int) (models.Category, error)
	Create(category *models.Category) (models.Category, error)
	Update(category *models.Category) (models.Category, error)
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

func (c *categoryServiceImpl) GetAll() ([]models.Category, error) {
	return c.repo.GetAll()
}

func (c *categoryServiceImpl) GetByID(id int) (models.Category, error) {
	return c.repo.GetByID(id)
}

func (c *categoryServiceImpl) Create(category *models.Category) (models.Category, error) {
	return c.repo.Create(category)
}

func (c *categoryServiceImpl) Update(category *models.Category) (models.Category, error) {
	return c.repo.Update(category)
}

func (c *categoryServiceImpl) Delete(id int) error {
	return c.repo.Delete(id)
}
