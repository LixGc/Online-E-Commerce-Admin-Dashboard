package services

import (
	"e-commerce/entity"
	"e-commerce/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepo = &repository.CategoryRepoMock{Mock: mock.Mock{}}
var categoryService = CategoryService{
	Repository: categoryRepo,
}

func TestCategoryServiceFindAllCategories(t *testing.T) {
	categoryRepo := &repository.CategoryRepoMock{}

	// Dummy category data
	dummyCategories := []entity.Category{
		{ID: 1, Type: "Electronics", SoldProductAmount: 50},
		{ID: 2, Type: "Clothing", SoldProductAmount: 30},
	}

	categoryRepo.On("FindAll").Return(dummyCategories)

	categoryService := CategoryService{Repository: categoryRepo}

	resultCategories, err := categoryService.FindAllCategories()

	assert.NoError(t, err)
	assert.Equal(t, dummyCategories, resultCategories)

	categoryRepo.AssertExpectations(t)
}
func TestCategoryServiceCreateCategory(t *testing.T) {

	categoryRepo := &repository.CategoryRepoMock{}

	// Dummy category data
	dummyCategory := entity.Category{
		ID:                1,
		Type:              "Electronics",
		SoldProductAmount: 0,
	}

	categoryRepo.On("Create", dummyCategory).Return(nil)

	categoryService := CategoryService{Repository: categoryRepo}

	err := categoryService.CreateCategory(dummyCategory)

	assert.NoError(t, err)

	categoryRepo.AssertExpectations(t)
	categoryRepo.Mock.AssertCalled(t, "Create", dummyCategory)
}
func TestCategoryServiceUpdateCategory(t *testing.T) {

	dummyCategory := &entity.Category{
		ID:                1,
		Type:              "Electronic",
		SoldProductAmount: 0,
	}

	categoryRepo := &repository.CategoryRepoMock{}

	categoryRepo.On("FindByType", "Furniture").Return(dummyCategory, nil)

	categoryRepo.On("Update", dummyCategory).Return(nil)

	categoryService := CategoryService{Repository: categoryRepo}

	userInput := CategoryInput{Type: "Furniture"}
	updatedCategory, err := categoryService.UpdateCategory(userInput)

	assert.NoError(t, err)

	assert.Equal(t, userInput.Type, updatedCategory.Type)

	categoryRepo.AssertExpectations(t)
	categoryRepo.Mock.AssertCalled(t, "FindByType", "Furniture")
	categoryRepo.Mock.AssertCalled(t, "Update", dummyCategory)
}
func TestCategoryServiceDeleteCategory(t *testing.T) {

	dummyCategory := &entity.Category{
		ID:                1,
		Type:              "Electronic",
		SoldProductAmount: 0,
	}

	categoryRepo := &repository.CategoryRepoMock{}

	categoryRepo.On("FindByID", uint(1)).Return(dummyCategory, nil)

	categoryRepo.On("Delete", dummyCategory).Return(nil)

	categoryService := CategoryService{Repository: categoryRepo}

	err := categoryService.DeleteCategory(1)

	assert.NoError(t, err)

	categoryRepo.AssertExpectations(t)
	categoryRepo.Mock.AssertCalled(t, "FindByID", uint(1))
	categoryRepo.Mock.AssertCalled(t, "Delete", dummyCategory)
}
