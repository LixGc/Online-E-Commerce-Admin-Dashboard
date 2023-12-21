package repository

import "e-commerce/entity"

type ProductRepo interface {
	FindAllProduct() []entity.Product
	CreateProduct(product entity.Product) error
	FindProductByID(productID string) (*entity.Product, error)
	UpdateProduct(product *entity.Product) error
	DeleteProduct(product *entity.Product) error
}
type CategoryRepo interface {
	FindAll() []entity.Category
	FindByID(id uint) (*entity.Category, error)
	Create(category entity.Category) error
	FindByType(categoryType string) (*entity.Category, error)
	Update(category *entity.Category) error
	Delete(category *entity.Category) error
}
type UserRepo interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	UpdateBalance(user *entity.User) error
}

type TransactionRepo interface {
	CreateTransactionHistory(transaction entity.TransactionHistory) error
	GetTransactionHistoryByUserID(userID uint) ([]entity.TransactionHistory, error)
	GetAllTransactionHistory() ([]entity.TransactionHistory, error)
}
