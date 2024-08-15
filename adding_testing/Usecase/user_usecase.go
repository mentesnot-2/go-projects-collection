package usecase

import (
	"errors"
	"github.com/mentesnot-2/adding_testing/Domain"
	"github.com/mentesnot-2/adding_testing/Infrastructure"
	"github.com/mentesnot-2/adding_testing/Repository"
)

type UserUsecase interface {
	CreateUser(user *domain.User) error
	Login(username, password string) (string, error)
}
type UserUsecaseImpl struct {
	userRepo repository.UserRepository
	passwordSvc infrastructure.PasswordService
	jwtSvc infrastructure.JWTService
}


func NewUserUseCase(userRepo repository.UserRepository, passwordSvc infrastructure.PasswordService,jwtSvc infrastructure.JWTService ) UserUsecase {
	return &UserUsecaseImpl{
		userRepo: userRepo,
		passwordSvc: passwordSvc,
		jwtSvc: jwtSvc,
	}
}



func (u *UserUsecaseImpl) CreateUser(user *domain.User) error {
	hashedPassword,err := u.passwordSvc.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return u.userRepo.CreateUser(user)
}

func (u *UserUsecaseImpl) Login(username,password string) (string,error) {
	var token string
	user,err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return "",err
	}
	err = u.passwordSvc.CheckPassword(password,user.Password)
	if err != nil {
		return "",errors.New("invalid password")
	}
	token,err = u.jwtSvc.GenerateToken(user.ID.Hex(),)
	if err != nil {
		return "",err
	}
	
	return token,nil

}

