package services

import (
	"e-commerce/entity"
	"e-commerce/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_Register(t *testing.T) {
	userRepo := &repository.UserRepoMock{}

	// Dummy user data
	dummyUser := &entity.User{
		FullName: "Felix Giancarlo",
		Email:    "felixgiancarlo789@gmail.com",
		Password: "hashed_password",
		Role:     "customer",
		Balance:  0,
	}

	userRepo.On("Create", mock.AnythingOfType("*entity.User")).Return(nil)

	userService := &UserService{UserRepository: userRepo}

	resultUser, err := userService.Register(RegisterInput{
		FullName: dummyUser.FullName,
		Email:    dummyUser.Email,
		Password: "dummy_password",
	})

	assert.NoError(t, err)
	assert.Equal(t, dummyUser.Email, resultUser.Email)
	assert.Equal(t, dummyUser.FullName, resultUser.FullName)
	assert.Equal(t, resultUser.Balance, 0)
	assert.Equal(t, resultUser.Role, "customer")

	userRepo.AssertExpectations(t)
}

func TestUserService_Login(t *testing.T) {
	userRepo := &repository.UserRepoMock{}

	dummyUser := &entity.User{
		FullName: "Felix Giancarlo",
		Email:    "felixgiancarlo789@gmail.com",
		Password: "felix123",
		Role:     "customer",
		Balance:  0,
	}

	userRepo.On("FindByEmail", dummyUser.Email).Return(dummyUser, nil)

	userService := &UserService{UserRepository: userRepo}

	loginInput := LoginInput{
		Email:    dummyUser.Email,
		Password: "felix123",
	}

	resultUser, err := userService.Login(loginInput)

	assert.NoError(t, err)
	assert.NotNil(t, resultUser)
	assert.Equal(t, dummyUser.FullName, resultUser.FullName)
	assert.Equal(t, dummyUser.Email, resultUser.Email)
}
