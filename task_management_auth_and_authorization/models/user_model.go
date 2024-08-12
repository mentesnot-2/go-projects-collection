package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password,omitempty" json:"-"`
	Role     string             `bson:"role" json:"role"`
}

type RegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserWithToken struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}