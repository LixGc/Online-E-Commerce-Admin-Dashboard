package services

import (
	"e-commerce/entity"
	"e-commerce/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepo = &repository.ProductRepoMock{Mock: mock.Mock{}}
var productService = ProductService{
	ProductRepository: productRepo,
}

func TestProductFindAll(t *testing.T) {

	productRepo := &repository.ProductRepoMock{}

	dummyProducts := []entity.Product{
		{ID: "1", Title: "AC", Price: 5000, Stock: 2, CategoryID: 1},
		{ID: "2", Title: "Remote", Price: 30000, Stock: 2, CategoryID: 2},
	}

	productRepo.On("FindAllProduct").Return(dummyProducts)

	productService := ProductService{ProductRepository: productRepo}

	resultProducts, err := productService.GetAllProducts()

	assert.NoError(t, err)
	assert.Equal(t, dummyProducts, resultProducts)

	productRepo.AssertExpectations(t)
}

func TestProductCreate(t *testing.T) {
	// Dummy product data
	dummyProduct := entity.Product{
		ID:         "3",
		Title:      "Smartphone",
		Price:      1000,
		Stock:      5,
		CategoryID: 3,
	}

	productRepo.On("CreateProduct", dummyProduct).Return(nil)

	err := productService.CreateProduct(dummyProduct)

	assert.NoError(t, err)

	productRepo.AssertExpectations(t)
	productRepo.Mock.AssertCalled(t, "CreateProduct", dummyProduct)
}

func TestProductUpdate(t *testing.T) {

	productRepo := &repository.ProductRepoMock{}

	dummyProduct := &entity.Product{
		ID:         "1",
		Title:      "AC",
		Price:      5000,
		Stock:      2,
		CategoryID: 1,
	}

	productRepo.On("FindProductByID", "1").Return(dummyProduct, nil)

	productRepo.On("UpdateProduct", dummyProduct).Return(nil)

	productService := ProductService{ProductRepository: productRepo}

	updateInput := ProductInput{
		ID:         "1",
		Title:      "Air Conditioner",
		Price:      6000,
		Stock:      3,
		CategoryID: 2,
	}

	updatedProduct, err := productService.UpdateProduct(updateInput.ID, updateInput)

	assert.NoError(t, err)

	assert.Equal(t, updateInput.Title, updatedProduct.Title)
	assert.Equal(t, updateInput.Price, updatedProduct.Price)
	assert.Equal(t, updateInput.Stock, updatedProduct.Stock)

	productRepo.AssertExpectations(t)
	productRepo.Mock.AssertCalled(t, "FindProductByID", "1")
	productRepo.Mock.AssertCalled(t, "UpdateProduct", dummyProduct)
}
func TestProductDelete(t *testing.T) {

	productRepo := &repository.ProductRepoMock{}

	dummyProduct := &entity.Product{
		ID:         "1",
		Title:      "AC",
		Price:      5000,
		Stock:      2,
		CategoryID: 1,
	}

	productRepo.On("FindProductByID", "1").Return(dummyProduct, nil)

	productRepo.On("DeleteProduct", dummyProduct).Return(nil)

	productService := ProductService{ProductRepository: productRepo}

	err := productService.DeleteProduct("1")

	assert.NoError(t, err)

	productRepo.AssertExpectations(t)
	productRepo.Mock.AssertCalled(t, "FindProductByID", "1")
	productRepo.Mock.AssertCalled(t, "DeleteProduct", dummyProduct)
}
