package data

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mentesnot-2/task_management/models"
)

var tasks = make(map[string]models.Task)


func GetTasks() []models.Task {
	values:=make([]models.Task,0,len(tasks))
	for _,task:= range tasks{
		values = append(values,task)
	}
	return values
}


func GetTask(id string) (models.Task,error) {
	task,exists:=tasks[id]
	if !exists {
		return models.Task{},errors.New("task not found")
	}
	return task,nil
}


func CreateTask(task models.Task) models.Task {
	task.ID = uuid.New().String()
	tasks[task.ID] = task
	return task
}

func UpdateTask(id string,updatedTask models.Task) (models.Task,error) {
	task,exists:=tasks[id]

	if !exists {
		return models.Task{},errors.New("Taks not found")
	}
	updatedTask.ID = task.ID
	tasks[id] = updatedTask
	return updatedTask,nil

}


func DeleteTask(id string) error {
	if _,exists :=tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(tasks,id)
	return nil
}