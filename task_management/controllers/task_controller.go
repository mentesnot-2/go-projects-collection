package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mentesnot-2/task_management/data"
	"github.com/mentesnot-2/task_management/models"
)

func GetTasks(c *gin.Context) {
	tasks:=data.GetTasks()
	c.JSON(http.StatusOK,tasks)
}


func GetTask(c *gin.Context) {
	id:=c.Param("id")
	task,err:=data.GetTask(id)
	if err !=nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"Task not found"})
		return 
	}
	c.JSON(http.StatusOK,task)
}


func CreateTask(c *gin.Context) {
	var task models.Task
	if err:=c.ShouldBindJSON(&task);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	fmt.Println(task)
	task = data.CreateTask(task)
	c.JSON(http.StatusCreated,task)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task, err := data.UpdateTask(id, updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := data.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"Task deleted successfully"})
}