package routers

import (
	"github.com/mentesnot-2/adding_testing/Delivery/controllers"
	"github.com/mentesnot-2/adding_testing/Infrastructure"
	"github.com/gin-gonic/gin"
)



func SetupRouter(taskController *controllers.TaskController, userController *controllers.UserController,jwtService infrastructure.JWTService) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	api.POST("/register",userController.CreateUser)
	api.POST("/login",userController.Login)

	api.Use(infrastructure.AuthMiddleware(jwtService))

	api.POST("/tasks",taskController.CreateTask)
	api.GET("/tasks",taskController.GetAllTask)
	api.GET("/tasks/:id",taskController.FindTaskById)
	api.PUT("/tasks/:id",taskController.UpdateTask)
	api.DELETE("/tasks/:id",taskController.DeleteTask)
	

	return r
}