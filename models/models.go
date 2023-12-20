package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName           string               `json:"full_name" validate:"required"`
	Email              string               `json:"email" validate:"required,email,unique"`
	Password           string               `json:"password" validate:"required,min=6"`
	Role               string               `json:"role" validate:"required,oneof=admin customer"`
	Balance            int                  `json:"balance" validate:"required,gte=0,lte=100000000"`
	TransactionHistory []TransactionHistory `gorm:"foreignKey:UserID" json:"transaction_history"`
}

type Product struct {
	gorm.Model
	Title              string               `json:"title" validate:"required"`
	Price              int                  `json:"price" validate:"required,gte=0,lte=50000000"`
	Stock              int                  `json:"stock" validate:"required,gte=5"`
	CategoryID         uint                 `json:"category_id"`
	TransactionHistory []TransactionHistory `gorm:"foreignKey:ProductID" json:"transaction_history"`
}

type Category struct {
	gorm.Model
	Type              string    `json:"type" validate:"required"`
	SoldProductAmount int       `json:"sold_product_amount"`
	Products          []Product `gorm:"foreignKey:CategoryID" json:"products"`
}

type TransactionHistory struct {
	gorm.Model
	ProductID  uint `json:"product_id" validate:"required"`
	UserID     uint `json:"user_id" validate:"required"`
	Quantity   int  `json:"quantity" validate:"required"`
	TotalPrice int  `json:"total_price" validate:"required"`
}
