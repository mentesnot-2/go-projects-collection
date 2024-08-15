package usecasetest

import (
	"testing"

	"github.com/mentesnot-2/adding_testing/Domain"
	"github.com/mentesnot-2/adding_testing/Usecase"
	"github.com/mentesnot-2/adding_testing/testing/mocks/infrastructure"
	"github.com/mentesnot-2/adding_testing/testing/mocks/repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserUserCase_CreateUser(t *testing.T){
	mockUserRepo := new(repository.UserRepository)
	mockPasswordSvc := new(infrastructure.PasswordService)
	mockJWTService := new(infrastructure.JWTService)


	// Create a new UserUsecase with the mock repository
	userUsecase := usecase.NewUserUseCase(mockUserRepo, mockPasswordSvc, mockJWTService)


	// Define the user to be creat

	user := &domain.User{
		Username : "testuser",
		Password : "testpassword",
	}

	mockPasswordSvc.On("HashPassword", user.Password).Return("hashedpassword", nil)
	mockUserRepo.On("CreateUser", user).Return(nil)

	// Call the CreateUser method
	err := userUsecase.CreateUser(user)

	// Check if the CreateUser method returned an error
	assert.NoError(t, err)
	assert.Equal(t, "hashedpassword", user.Password)
	mockPasswordSvc.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}


func TestUserUsecase_Login(t *testing.T) {
	mockUserRepo := new(repository.UserRepository)
	mockPasswordSvc := new(infrastructure.PasswordService)
	mockJWTService := new(infrastructure.JWTService)

	userUsecase := usecase.NewUserUseCase(mockUserRepo, mockPasswordSvc, mockJWTService)

	userID := primitive.NewObjectID()
	user := &domain.User{
		Username: "testuser",
		Password: "testpassword",
		ID: userID,
	}
	mockUserRepo.On("GetUserByUsername", user.Username).Return(*user, nil)
	mockPasswordSvc.On("CheckPassword", user.Password, user.Password).Return(nil)
	mockJWTService.On("GenerateToken", user.ID.Hex()).Return("token", nil)

	token,err := userUsecase.Login(user.Username, user.Password)


	assert.NoError(t, err)
	assert.Equal(t, "token", token)
	mockUserRepo.AssertExpectations(t)
	mockPasswordSvc.AssertExpectations(t)
	mockJWTService.AssertExpectations(t)
}