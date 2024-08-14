package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/mentesnot-2/refactoring_with_clean_architecture/Domain"
	"github.com/mentesnot-2/refactoring_with_clean_architecture/Usecase"
)


type UserController struct {
	userUsecase usecase.UserUsecase
}
type TaskController struct {
	taskUsecase usecase.TaskUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *UserController {
	return &UserController{userUsecase}
}


func (u *UserController) CreateUser(c *gin.Context) {
	var userInput domain.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}
	err := u.userUsecase.CreateUser(&userInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated,gin.H{"message":"user created successfully"})
}


func (u *UserController) Login(c *gin.Context) {
	var userInput domain.UserLoginInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}
	userLogin,err := u.userUsecase.Login(userInput.Username,userInput.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK,gin.H{"data":userLogin})
}

func NewTaskController(taskUsecase usecase.TaskUsecase) *TaskController {
	return &TaskController{taskUsecase}
}

func (t *TaskController) CreateTask(c *gin.Context) {
	var taskInput domain.Task
	userId := c.MustGet("userId").(string)
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	taskInput.UserID = objectId
	if err := c.ShouldBindJSON(&taskInput); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}
	err = t.taskUsecase.CreateTask(&taskInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated,gin.H{"message":"task created successfully"})
}


func (t *TaskController) GetAllTask(c *gin.Context) {
	tasks,err := t.taskUsecase.GetAllTask()
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK,gin.H{"data":tasks})
}


func (t *TaskController) FindTaskById(c *gin.Context) {
	id := c.Param("id")
	task,err := t.taskUsecase.FindTaskById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK,gin.H{"data":task})
}

func (t *TaskController) UpdateTask(c *gin.Context) {
	var taskInput domain.Task
	if err := c.ShouldBindJSON(&taskInput); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}
	err := t.taskUsecase.UpdateTask(&taskInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK,gin.H{"message":"task updated successfully"})
}

func (t *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := t.taskUsecase.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK,gin.H{"message":"task deleted successfully"})
}