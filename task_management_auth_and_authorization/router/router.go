package router


import (
	"github.com/gin-gonic/gin"
	"github.com/mentesnot-2/task_management_auth_and_authorization/controller"
	"github.com/mentesnot-2/task_management_auth_and_authorization/middleware"


)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // User routes
    r.POST("/register", controller.CreateUser)
    r.POST("/login", controller.LoginUser)

    // Task routes
    taskRoutes := r.Group("/tasks")
    taskRoutes.Use(middleware.AuthMiddleware())
    {
        taskRoutes.GET("/", controller.GetTasks)
        taskRoutes.GET("/:id", controller.GetTask)

        taskRoutes.Use(middleware.AdminMiddleware())
        {
            taskRoutes.POST("/", controller.CreateTask)
            taskRoutes.PUT("/:id", controller.UpdatedTask)
            taskRoutes.DELETE("/:id", controller.DeleteTask)
			taskRoutes.PUT("/promote/:id", controller.PromoteUser)
        }
    }

    return r
}
