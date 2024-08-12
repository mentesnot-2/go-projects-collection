package data

import (
	"context"
	"github.com/mentesnot-2/task_management_auth_and_authorization/data/config"
	"github.com/mentesnot-2/task_management_auth_and_authorization/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	cursor, err := config.TasksCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil,err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var task models.Task
		if err:= cursor.Decode(&task);err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err:=cursor.Err();err != nil {
		return nil, err
	}
	return tasks,nil
}

func GetTask(id string) (models.Task, error) {
	var task models.Task
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, err
	}
	err = config.TasksCollection.FindOne(context.TODO(), bson.M{"_id": ObjectID}).Decode(&task)

	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func CreateTask(task models.Task) (*mongo.InsertOneResult,error) {
	if err := task.Validate(); err != nil {
		return nil,err
	}
	result, err := config.TasksCollection.InsertOne(context.TODO(), task)
	if err != nil {
		return nil,err
	}
	return result, nil
}

func UpdateTask(id string, updatedTask models.Task) (*mongo.UpdateResult, error) {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil,err
	}
	var task models.Task
	res := config.TasksCollection.FindOne(context.TODO(), bson.M{"_id": ObjectID}).Decode(&task)
	if res != nil {
		return nil, err
	}

	if updatedTask.UserID != task.UserID {
		return nil, err
	}
	if err := updatedTask.Validate(); err != nil {
		return nil,err
	}

	filter:=bson.M{"_id":ObjectID}
	update:=bson.M{
		"$set" :updatedTask,
	}

	
	result, err := config.TasksCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil,err
	}
	return result, nil
	
}



func DeleteTask(id string) (*mongo.DeleteResult, error) {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil,err
	}
	filter:=bson.M{"_id":ObjectID}
	result, err := config.UsersCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil,err
	}
	return result, nil

}

