package models

type Category struct {
	CategoryID   int    `json:"category_id" gorm:"primaryKey;autoIncrement;column:CategoryID"`
	CategoryName string `json:"category_name" gorm:"column:CategoryName" validate:"required"`
	Description  string `json:"description" gorm:"column:Description" validate:"required"`
}
