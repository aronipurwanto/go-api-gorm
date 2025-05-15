package repositories

import (
	"github.com/aronipurwanto/go-api-gorm/models"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	GetAll() ([]models.Employee, error)
	GetPaginated(page int, limit int) ([]models.Employee, int64, error)
	GetByID(id int) (*models.Employee, error)
	Create(employee *models.Employee) error
	Update(employee *models.Employee) error
	Delete(id int) error
}

type employeeRepoImpl struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepoImpl{db}
}

func (r *employeeRepoImpl) GetAll() ([]models.Employee, error) {
	var employees []models.Employee
	err := r.db.Find(&employees).Error
	return employees, err
}

func (r *employeeRepoImpl) GetByID(id int) (*models.Employee, error) {
	var employee models.Employee
	err := r.db.First(&employee, id).Error
	return &employee, err
}

func (r *employeeRepoImpl) GetPaginated(page int, limit int) ([]models.Employee, int64, error) {
	var employees []models.Employee
	var total int64

	offset := (page - 1) * limit
	if err := r.db.Model(&models.Employee{}).Count(&total).Limit(limit).Offset(offset).Find(&employees).Error; err != nil {
		return nil, 0, err
	}

	return employees, total, nil
}

func (r *employeeRepoImpl) Create(employee *models.Employee) error {
	return r.db.Create(employee).Error
}

func (r *employeeRepoImpl) Update(employee *models.Employee) error {
	return r.db.Save(employee).Error
}

func (r *employeeRepoImpl) Delete(id int) error {
	return r.db.Delete(&models.Employee{}, id).Error
}
