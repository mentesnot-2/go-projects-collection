package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Task struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	UserID primitive.ObjectID `bson:"user_id,omitempty"`
	Title string `bson:"title,omitempty"`
	Description string `bson:"description,omitempty"`
	Completed bool `bson:"completed"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}


type TaskRepository interface {
	CreateTask(task *Task) error
	FindTaskById(id string) (*Task,error)
	GetAllTask() ([]*Task,error)
	UpdateTask(task *Task) error
	DeleteTask(id string) error
}


type  User struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Username string `bson:"username,omitempty"`
	Email string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`


}
type UserLogin struct {
	User *User  
	Token string  
}
type UserLoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRepository interface {
	CreateUser(user *User) error
	Login(username,password string) (UserLogin,error)
	GetUserByUsername(username string) (User,error)
}