package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model         `swaggerignore:"true"`
	FullName           string               `json:"full_name"`
	Email              string               `json:"email"`
	Password           string               `json:"password"`
	Role               string               `json:"role"`
	Balance            int                  `json:"balance"`
	TransactionHistory []TransactionHistory `gorm:"foreignKey:UserID" json:"transaction_history"`
}

type Product struct {
	gorm.Model         `swaggerignore:"true"`
	Title              string               `json:"title"`
	Price              int                  `json:"price"`
	Stock              int                  `json:"stock"`
	CategoryID         uint                 `json:"category_id"`
	TransactionHistory []TransactionHistory `gorm:"foreignKey:ProductID" json:"transaction_history"`
}

type Category struct {
	gorm.Model        `swaggerignore:"true"`
	Type              string    `json:"type"`
	SoldProductAmount int       `json:"sold_product_amount"`
	Products          []Product `gorm:"foreignKey:CategoryID" json:"products"`
}

type TransactionHistory struct {
	gorm.Model `swaggerignore:"true"`
	ProductID  uint    `json:"product_id"`
	UserID     uint    `json:"user_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice int     `json:"total_price"`
	Product    Product `gorm:"foreignKey:ProductID" json:"product"`
}
