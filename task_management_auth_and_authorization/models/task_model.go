package models

import (
	"time"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task struct defines the task model
type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	DueDate     time.Time          `bson:"due_date" json:"due_date"`
	Status      string             `bson:"status" json:"status"`
	UserID      primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"` // Foreign key to User
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