package services

import (
	"e-commerce/entity"
	"e-commerce/repository"
	"errors"
)

type ProductService struct {
	ProductRepository repository.ProductRepo
}
type ProductInput struct {
	ID         string
	Title      string
	Price      int
	Stock      int
	CategoryID int
}

func (ps ProductService) GetAllProducts() ([]entity.Product, error) {
	products := ps.ProductRepository.FindAllProduct()
	if len(products) == 0 {
		return nil, errors.New("products not found")
	}
	return products, nil
}
func (ps ProductService) CreateProduct(product entity.Product) error {
	if product.Title == "" || product.Price <= 0 || product.Stock < 0 || product.CategoryID == 0 {
		return errors.New("invalid product input")
	}

	return ps.ProductRepository.CreateProduct(product)
}

func (ps ProductService) UpdateProduct(productID string, userInput ProductInput) (*entity.Product, error) {
	existingProduct, err := ps.ProductRepository.FindProductByID(productID)
	if err != nil {
		return nil, err
	}

	if existingProduct == nil {
		return nil, errors.New("product not found")
	}

	existingProduct.Title = userInput.Title
	existingProduct.Price = userInput.Price
	existingProduct.Stock = userInput.Stock
	existingProduct.CategoryID = userInput.CategoryID

	if err := ps.validateProduct(existingProduct); err != nil {
		return nil, err
	}

	err = ps.ProductRepository.UpdateProduct(existingProduct)
	if err != nil {
		return nil, err
	}

	return existingProduct, nil
}
func (ps ProductService) DeleteProduct(productID string) error {
	existingProduct, err := ps.ProductRepository.FindProductByID(productID)
	if err != nil {
		return err
	}

	if existingProduct == nil {
		return errors.New("product not found")
	}
	err = ps.ProductRepository.DeleteProduct(existingProduct)
	if err != nil {
		return err
	}

	return nil
}

func (ps ProductService) validateProduct(product *entity.Product) error {
	if product.Title == "" || product.Price <= 0 || product.Stock < 0 || product.CategoryID == 0 {
		return errors.New("invalid product input")
	}
	return nil
}
