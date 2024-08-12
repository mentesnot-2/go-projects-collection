package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mentesnot-2/task_management_auth_and_authorization/data"
	"github.com/mentesnot-2/task_management_auth_and_authorization/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTask(c *gin.Context) {
	var createTask models.Task
	if err := c.ShouldBindJSON(&createTask); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}
	userID, err := primitive.ObjectIDFromHex(user_id.(string))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	createTask.UserID = userID
	task, err := data.CreateTask(createTask)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "Task created successfully", "result": task})
}

func GetTasks(c *gin.Context) {
	tasks, err := data.GetTasks()
	fmt.Println(tasks, err)
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
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}
	userID, err := primitive.ObjectIDFromHex(user_id.(string))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	updatedTask.UserID = userID
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
	c.JSON(200, gin.H{"message": "Task retrieved successfully", "Task": task})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	task, err := data.DeleteTask(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
	}
	c.JSON(200, gin.H{"message": "Task deleted successfully", "result": task})
}

func CreateUser(c *gin.Context) {
	var userRegister models.RegisterUser
	if err := c.ShouldBindJSON(&userRegister); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	createdUser, err := data.CreateUser(userRegister)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "User created successfully", "result": createdUser})
}

func GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := data.GetUserByUsername(username)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, gin.H{"message": "User retrieved successfully", "User": user})
}

func LoginUser(c *gin.Context) {
	var user models.UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	username := user.Username
	password := user.Password
	UserWithToken, err := data.LoginUser(username, password)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Login successful", "User": UserWithToken})
}

func PromoteUser(c *gin.Context) {
	username := c.Param("username")
	user, err := data.PromoteUser(username)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, gin.H{"message": "User promoted successfully", "User": user})
}
