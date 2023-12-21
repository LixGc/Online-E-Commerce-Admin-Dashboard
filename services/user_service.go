package services

import (
	"e-commerce/entity"
	"e-commerce/helpers"
	"e-commerce/repository"
	"errors"
)

type UserService struct {
	UserRepository repository.UserRepo
}

type RegisterInput struct {
	FullName string
	Email    string
	Password string
}
type LoginInput struct {
	Email    string
	Password string
}
type UpdateUserBalanceInput struct {
	Email   string
	Balance int
}

func (us *UserService) Register(input RegisterInput) (*entity.User, error) {

	if input.FullName == "" || input.Email == "" || input.Password == "" {
		return nil, errors.New("full name, email, and password cannot be empty")
	}

	if len(input.Password) < 6 {
		return nil, errors.New("password length must be at least 6 characters")
	}

	hashedPassword, err := helpers.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	newUser := &entity.User{
		FullName: input.FullName,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     "customer",
		Balance:  0,
	}

	err = us.UserRepository.Create(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
func (us *UserService) Login(input LoginInput) (*entity.User, error) {

	user, err := us.UserRepository.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	if err := helpers.ComparePassword(hashedPassword, input.Password); err != nil {
		return nil, errors.New("incorrect password")
	}

	return user, nil
}
func (us *UserService) UpdateUserBalance(input UpdateUserBalanceInput) (*entity.User, error) {

	user, err := us.UserRepository.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	user.Balance = input.Balance

	err = us.UserRepository.UpdateBalance(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
