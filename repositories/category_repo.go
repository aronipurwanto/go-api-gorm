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

func (c *categoryRepoImpl) GetAll() ([]models.Category, error) {
	var categories []models.Category
	err := c.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *categoryRepoImpl) GetByID(id int) (models.Category, error) {
	var category models.Category
	err := c.db.First(&category, id).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (c *categoryRepoImpl) Create(category models.Category) (models.Category, error) {
	err := c.db.Create(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (c *categoryRepoImpl) Update(category models.Category) (models.Category, error) {
	err := c.db.Save(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (c *categoryRepoImpl) Delete(id int) error {
	err := c.db.Delete(&models.Category{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
