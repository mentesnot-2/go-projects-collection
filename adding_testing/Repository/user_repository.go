package repository

import (
	"context"
	"time"

	"github.com/mentesnot-2/adding_testing/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByUsername(username string) (domain.User, error)
}
type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{collection: db.Collection("users")}
}

func (u *userRepository) CreateUser(user *domain.User) error {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := u.collection.InsertOne(context.Background(), user)
	return err
}

func (u *userRepository) GetUserByUsername(username string) (domain.User, error) {
	var user domain.User
	err := u.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
