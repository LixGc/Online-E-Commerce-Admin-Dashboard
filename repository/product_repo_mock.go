package repository

import (
	"e-commerce/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepoMock struct {
	mock.Mock
}

func (prm *ProductRepoMock) FindAllProduct() []entity.Product {
	arguments := prm.Called()
	if arguments.Get(0) == nil {
		return nil
	}
	products := arguments.Get(0).([]entity.Product)
	return products
}

func (prm *ProductRepoMock) CreateProduct(product entity.Product) error {
	arguments := prm.Called(product)
	return arguments.Error(0)
}

func (prm *ProductRepoMock) FindProductByID(productID string) (*entity.Product, error) {
	arguments := prm.Called(productID)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	product := arguments.Get(0).(*entity.Product)
	return product, arguments.Error(1)
}

func (prm *ProductRepoMock) UpdateProduct(product *entity.Product) error {
	arguments := prm.Called(product)
	return arguments.Error(0)
}
func (prm *ProductRepoMock) DeleteProduct(product *entity.Product) error {
	arguments := prm.Called(product)
	return arguments.Error(0)
}
