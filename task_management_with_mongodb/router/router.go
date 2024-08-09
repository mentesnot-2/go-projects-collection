package router


import (
	"github.com/gin-gonic/gin"
	"github.com/mentesnot-2/task_management_with_mongodb/controller"


)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/tasks", controller.GetTasks)
	r.GET("/tasks/:id", controller.GetTask)
	r.POST("/tasks", controller.CreateTask)
	r.PUT("/tasks/:id", controller.UpdatedTask)
	r.DELETE("/tasks/:id", controller.DeleteTask)
	return r
}