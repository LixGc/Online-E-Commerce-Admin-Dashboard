package services

import (
	"e-commerce/entity"
	"e-commerce/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepo
}
type CategoryInput struct {
	Type string
}

func (cs CategoryService) FindAllCategories() ([]entity.Category, error) {
	categories := cs.Repository.FindAll()
	if len(categories) == 0 {
		return nil, errors.New("categories not found")
	}
	return categories, nil
}

func (cs CategoryService) CreateCategory(category entity.Category) error {

	if category.Type == "" {
		return errors.New("category type cannot be empty")
	}

	return cs.Repository.Create(category)
}
func (cs CategoryService) UpdateCategory(userInput CategoryInput) (*entity.Category, error) {
	existingCategory, err := cs.Repository.FindByType(userInput.Type)
	if err != nil {
		return nil, err
	}
	if existingCategory == nil {
		return nil, errors.New("category not found")
	}

	existingCategory.Type = userInput.Type

	err = cs.Repository.Update(existingCategory)
	if err != nil {
		return nil, err
	}

	return existingCategory, nil
}
func (cs CategoryService) DeleteCategory(categoryID uint) error {
	existingCategory, err := cs.Repository.FindByID(categoryID)
	if err != nil {
		return err
	}

	if existingCategory == nil {
		return errors.New("category not found")
	}

	err = cs.Repository.Delete(existingCategory)
	if err != nil {
		return err
	}

	return nil
}
