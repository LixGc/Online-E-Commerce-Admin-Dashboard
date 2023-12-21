package repository

import (
	"e-commerce/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepoMock struct {
	mock.Mock
}

func (crm *CategoryRepoMock) FindByID(id uint) (*entity.Category, error) {
	arguments := crm.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	category := arguments.Get(0).(*entity.Category)
	return category, arguments.Error(1)
}

func (crm *CategoryRepoMock) FindByType(categoryType string) (*entity.Category, error) {
	arguments := crm.Mock.Called(categoryType)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	category := arguments.Get(0).(*entity.Category)
	return category, arguments.Error(1)
}
func (crm *CategoryRepoMock) FindAll() []entity.Category {
	arguments := crm.Called()
	if arguments.Get(0) == nil {
		return nil
	}
	categories := arguments.Get(0).([]entity.Category)
	return categories
}
func (crm *CategoryRepoMock) Create(category entity.Category) error {
	arguments := crm.Called(category)
	return arguments.Error(0)
}

func (crm *CategoryRepoMock) Update(category *entity.Category) error {
	arguments := crm.Mock.Called(category)
	return arguments.Error(0)
}
func (crm *CategoryRepoMock) Delete(category *entity.Category) error {
	arguments := crm.Mock.Called(category)
	return arguments.Error(0)
}
