package repositories

import (
	"github.com/aronipurwanto/go-api-gorm/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrderWithDetails(order *models.Order, details []models.OrderDetail) error
	GetAll(page, limit int) ([]models.Order, int64, error)
	GetByID(id int) (*models.Order, []models.OrderDetail, error)
	Update(id int, order *models.Order, details []models.OrderDetail) error
	Delete(id int) error
}

type orderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepositoryImpl{db: db}
}

func (r *orderRepositoryImpl) CreateOrderWithDetails(order *models.Order, details []models.OrderDetail) error {
	tx := r.db.Begin()
	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, detail := range details {
		detail.OrderID = order.OrderID
		if err := tx.Create(&detail).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *orderRepositoryImpl) GetAll(page, limit int) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64

	offset := (page - 1) * limit
	err := r.db.Model(&models.Order{}).Count(&total).
		Limit(limit).Offset(offset).
		Find(&orders).Error
	return orders, total, err
}

func (r *orderRepositoryImpl) GetByID(id int) (*models.Order, []models.OrderDetail, error) {
	var order models.Order
	if err := r.db.First(&order, id).Error; err != nil {
		return nil, nil, err
	}
	var details []models.OrderDetail
	if err := r.db.Where("OrderID = ?", id).Find(&details).Error; err != nil {
		return nil, nil, err
	}
	return &order, details, nil
}

func (r *orderRepositoryImpl) Delete(id int) error {
	tx := r.db.Begin()
	if err := tx.Where("OrderID = ?", id).Delete(&models.OrderDetail{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete(&models.Order{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *orderRepositoryImpl) Update(id int, order *models.Order, details []models.OrderDetail) error {
	tx := r.db.Begin()

	if err := tx.Model(&models.Order{}).Where("OrderID = ?", id).Updates(order).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("OrderID = ?", id).Delete(&models.OrderDetail{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, d := range details {
		d.OrderID = id
		if err := tx.Create(&d).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
