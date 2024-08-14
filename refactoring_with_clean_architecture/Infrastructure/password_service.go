package infrastructure


import (
	"golang.org/x/crypto/bcrypt"
)


type PasswordService interface {
	HashPassword(password string) (string,error)
	CheckPassword(password,hash string) error
}

type passwordServiceImpl struct {}

func NewPasswordService() PasswordService {
	return &passwordServiceImpl{}
}


func (p *passwordServiceImpl) HashPassword(passsword string) (string,error) {
	bytes,err := bcrypt.GenerateFromPassword([]byte(passsword),14)
	return string(bytes),err
}

func (p *passwordServiceImpl) CheckPassword(password,hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
}