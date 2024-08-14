package infrastructure

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)


type JWTService interface {
	GenerateToken(userId string) (string,error) 
	ValidateToken(token string) (*jwt.Token,error)
}



type JWTServiceImpl struct{
	secretKey string
}



func NewJWTService(secretKey string) JWTService{
	return &JWTServiceImpl{secretKey: secretKey}
}

func (j *JWTServiceImpl) GenerateToken(userID string) (string,error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString([]byte(j.secretKey))
}


func (j *JWTServiceImpl) ValidateToken(token string) (*jwt.Token,error) {
	return jwt.Parse(token,func(token *jwt.Token) (interface{},error) {
		return []byte(j.secretKey),nil
	})
}

