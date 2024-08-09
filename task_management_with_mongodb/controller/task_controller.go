package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mentesnot-2/task_management_with_mongodb/data"
	"github.com/mentesnot-2/task_management_with_mongodb/models"
	"net/http"
)

func CreateTask(c *gin.Context) {
	var createTask models.Task
	if err := c.ShouldBindJSON(&createTask); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	task, err := data.CreateTask(createTask)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, task)
}

func GetTasks(c *gin.Context) {
	tasks, err := data.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks, "Number of tasks": len(tasks)})
}

func UpdatedTask(c *gin.Context) {
	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := updatedTask.Validate(); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	task, err := data.UpdateTask(id, updatedTask)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Task updated successfully", "result": task})
}

func GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := data.GetTask(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(200, task)
}

func DeleteTask(c *gin.Context) {
	id:=c.Param("id")
	task,err := data.DeleteTask(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
	}
	c.JSON(200, gin.H{"message": "Task deleted successfully", "result": task})
}
