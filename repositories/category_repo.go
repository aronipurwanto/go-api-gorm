package repositories

import (
	"github.com/aronipurwanto/go-api-gorm/models"
	"gorm.io/gorm"
)

type CategoryRepo interface {
	GetAll() ([]models.Category, error)
	GetByID(id int) (models.Category, error)
	Create(category models.Category) (models.Category, error)
	Update(category models.Category) (models.Category, error)
	Delete(id int) error
}

type categoryRepoImpl struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) CategoryRepo {
	return &categoryRepoImpl{
		db: db,
	}
}
