package models

type Category struct {
	CategoryId   string `json:"category_id" gorm:"primaryKey;column:CategoryID" validate:"required"`
	CategoryName string `json:"category_name" gorm:"column:CategoryName" validate:"required"`
	Description  string `json:"description" gorm:"column:Description"`
}
