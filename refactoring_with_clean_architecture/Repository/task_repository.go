package repository

import (
	"github.com/mentesnot-2/refactoring_with_clean_architecture/Domain"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)
type TaskRepository interface {
	CreateTask(task *domain.Task) error
	FindTaskById(id string) (*domain.Task, error)
	GetAllTask() ([]*domain.Task, error)
	UpdateTask(task *domain.Task) error
	DeleteTask(id string) error
	
}
type taskRepository struct {
	collection *mongo.Collection
}


func NewTaskRepository(db *mongo.Database) TaskRepository {
	return &taskRepository{collection:db.Collection("tasks")}
}


func (t *taskRepository) CreateTask(task *domain.Task) error {
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	_,err := t.collection.InsertOne(context.Background(),task)
	return err
}

func (t *taskRepository) FindTaskById(id string) (*domain.Task, error) {
	var task domain.Task
	objectId,err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil,err
	}
	err = t.collection.FindOne(context.Background(),bson.M{"_id":objectId}).Decode(&task)
	if err != nil {
		return nil,err
	}
	return &task,nil
}

func (t *taskRepository) GetAllTask() ([]*domain.Task, error) {
	var tasks []*domain.Task
	cursor,err := t.collection.Find(context.Background(),bson.M{})
	if err != nil {
		return nil,err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var task domain.Task
		cursor.Decode(&task)
		tasks = append(tasks,&task)
	}
	return tasks,nil
}


func (t *taskRepository) UpdateTask(task *domain.Task) error {
	task.UpdatedAt = time.Now()
	_,err := t.collection.ReplaceOne(context.Background(),bson.M{"_id":task.ID},task)
	return err
}


func (t *taskRepository) DeleteTask(id string) error {
	objectId,err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_,err = t.collection.DeleteOne(context.Background(),bson.M{"_id":objectId})
	return err
}