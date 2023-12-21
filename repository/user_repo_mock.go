package repository

import (
	"e-commerce/entity"

	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (urm *UserRepoMock) Create(user *entity.User) error {
	arguments := urm.Called(user)
	return arguments.Error(0)
}

func (urm *UserRepoMock) FindByEmail(email string) (*entity.User, error) {
	arguments := urm.Called(email)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	user := arguments.Get(0).(*entity.User)
	return user, arguments.Error(1)
}

func (crm *UserRepoMock) UpdateBalance(user *entity.User) error {
	arguments := crm.Mock.Called(user)
	return arguments.Error(0)
}
func (trm *TransactionRepoMock) GetAllTransactionHistory() ([]entity.TransactionHistory, error) {
	args := trm.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	transactions := args.Get(0).([]entity.TransactionHistory)
	return transactions, args.Error(1)
}
