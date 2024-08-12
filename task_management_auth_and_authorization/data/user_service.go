package data

import (
	"context"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/mentesnot-2/task_management_auth_and_authorization/data/config"
	"github.com/mentesnot-2/task_management_auth_and_authorization/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(userRegister models.RegisterUser) (*models.User, error) {
	var user models.User
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegister.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error while hashing the password: ", err)
		return nil, err
	}
	userRegister.Password = string(hashedPassword)

	id := primitive.NewObjectID()
	user.ID = id
	user.Username = userRegister.Username
	user.Password = userRegister.Password
	count, err := config.UsersCollection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		fmt.Println("Error while counting the number of users: ", err)	
		return nil, err
	}
	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}
	
	_, err = config.UsersCollection.InsertOne(context.Background(), user)
	if err != nil {
		fmt.Println("Error while inserting the user: ", err)
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := config.UsersCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func LoginUser(username string, password string) (models.UserWithToken, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return models.UserWithToken{}, err
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.UserWithToken{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
	})
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	jwt_token, err := token.SignedString(jwtSecret)

	if err != nil {
		fmt.Println("Error while generating token: ", err)
		return models.UserWithToken{}, err
	}

	return models.UserWithToken{User: user, Token: jwt_token}, nil
}

func PromoteUser(username string) (models.User, error) {
	_, err := config.UsersCollection.UpdateOne(context.Background(), bson.M{"username": username}, bson.M{"$set": bson.M{"role": "admin"}})
	if err != nil {
		return models.User{}, err
	}
	return GetUserByUsername(username)
}
