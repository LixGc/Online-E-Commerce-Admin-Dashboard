package services

import (
	"e-commerce/entity"
	"e-commerce/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var transactionRepo = &repository.TransactionRepoMock{Mock: mock.Mock{}}
var transactionService = TransactionService{
	TransactionRepository: transactionRepo,
}

func TestTransactionServiceCreateTransactionHistory(t *testing.T) {

	transactionRepo := &repository.TransactionRepoMock{}

	dummyTransaction := entity.TransactionHistory{
		UserID:     1,
		ProductID:  2,
		Quantity:   3,
		TotalPrice: 300,
	}

	transactionRepo.On("CreateTransactionHistory", dummyTransaction).Return(nil)

	transactionService := TransactionService{TransactionRepository: transactionRepo}

	err := transactionService.CreateTransactionHistory(TransactionHistoryInput{
		UserID:     dummyTransaction.UserID,
		ProductID:  dummyTransaction.ProductID,
		Quantity:   dummyTransaction.Quantity,
		TotalPrice: dummyTransaction.TotalPrice,
	})

	assert.NoError(t, err)

	transactionRepo.AssertExpectations(t)
	transactionRepo.Mock.AssertCalled(t, "CreateTransactionHistory", dummyTransaction)
}
func TestTransactionServiceGetTransactionHistoryByUserID(t *testing.T) {
	transactionRepo := &repository.TransactionRepoMock{}

	// Dummy user ID
	userID := uint(1)

	dummyTransaction := []entity.TransactionHistory{
		{UserID: userID, ProductID: 1, Quantity: 2, TotalPrice: 200},
		{UserID: userID, ProductID: 2, Quantity: 1, TotalPrice: 150},
	}

	transactionRepo.On("GetTransactionHistoryByUserID", userID).Return(dummyTransaction, nil)

	transactionService := TransactionService{TransactionRepository: transactionRepo}

	resultTransactions, err := transactionService.GetTransactionHistoryByUserID(userID)

	assert.NoError(t, err)
	assert.Equal(t, dummyTransaction, resultTransactions)

	transactionRepo.AssertExpectations(t)
	transactionRepo.Mock.AssertCalled(t, "GetTransactionHistoryByUserID", userID)
}
func TestTransactionServiceGetAllTransactionHistory(t *testing.T) {

	transactionRepo := &repository.TransactionRepoMock{}

	dummyTransactions := []entity.TransactionHistory{
		{UserID: 1, ProductID: 1, Quantity: 2, TotalPrice: 200},
		{UserID: 2, ProductID: 2, Quantity: 1, TotalPrice: 150},
	}

	transactionRepo.On("GetAllTransactionHistory").Return(dummyTransactions, nil)

	transactionService := TransactionService{TransactionRepository: transactionRepo}

	resultTransactions, err := transactionService.GetAllTransactionHistory()

	assert.NoError(t, err)
	assert.Equal(t, dummyTransactions, resultTransactions)

	transactionRepo.AssertExpectations(t)
	transactionRepo.Mock.AssertCalled(t, "GetAllTransactionHistory")
}
