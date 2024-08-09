package models


import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"

)


type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Completed   bool               `bson:"completed"`
}

func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("title cannot be empty")
	}
	if t.Description == "" {
		return errors.New("description cannot be empty")
	}
	return nil
}