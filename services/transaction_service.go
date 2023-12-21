package services

import (
	"e-commerce/entity"
	"e-commerce/repository"
	"errors"
)

type TransactionService struct {
	TransactionRepository repository.TransactionRepo
}

type TransactionHistoryInput struct {
	UserID     uint
	ProductID  uint
	Quantity   int
	TotalPrice int
}

func (ts TransactionService) CreateTransactionHistory(input TransactionHistoryInput) error {
	if input.UserID == 0 || input.ProductID == 0 || input.Quantity <= 0 || input.TotalPrice <= 0 {
		return errors.New("invalid transaction input")
	}

	transaction := entity.TransactionHistory{
		UserID:     input.UserID,
		ProductID:  input.ProductID,
		Quantity:   input.Quantity,
		TotalPrice: input.TotalPrice,
	}

	return ts.TransactionRepository.CreateTransactionHistory(transaction)
}
func (ts TransactionService) GetTransactionHistoryByUserID(userID uint) ([]entity.TransactionHistory, error) {
	if userID == 0 {
		return nil, errors.New("invalid user ID")
	}

	return ts.TransactionRepository.GetTransactionHistoryByUserID(userID)
}
func (ts TransactionService) GetAllTransactionHistory() ([]entity.TransactionHistory, error) {
	return ts.TransactionRepository.GetAllTransactionHistory()
}
