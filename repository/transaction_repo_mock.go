package repository

import (
	"e-commerce/entity"

	"github.com/stretchr/testify/mock"
)

type TransactionRepoMock struct {
	mock.Mock
}

func (trm *TransactionRepoMock) CreateTransactionHistory(transaction entity.TransactionHistory) error {
	args := trm.Called(transaction)
	return args.Error(0)
}
func (trm *TransactionRepoMock) GetTransactionHistoryByUserID(userID uint) ([]entity.TransactionHistory, error) {
	args := trm.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	transactions := args.Get(0).([]entity.TransactionHistory)
	return transactions, args.Error(1)
}
