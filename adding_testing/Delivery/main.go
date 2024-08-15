package main

import (
	"github.com/joho/godotenv"
	"github.com/mentesnot-2/adding_testing/Delivery/controllers"
	"github.com/mentesnot-2/adding_testing/Delivery/routers"
	"github.com/mentesnot-2/adding_testing/Infrastructure"
	"github.com/mentesnot-2/adding_testing/Repository"
	"github.com/mentesnot-2/adding_testing/Usecase"
	"log"
	"os"
	"context"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	mongoConfig := infrastructure.NewMongoDBConfig(os.Getenv("MONGO_URI"), os.Getenv("MONGO_DB"))
	mongoClient, err := mongoConfig.Connect()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB")
	}
	defer mongoClient.Disconnect(context.TODO())

	db := mongoClient.Database(os.Getenv("MONGO_DB"))

	userRepo := repository.NewUserRepository(db)
	taskRepo := repository.NewTaskRepository(db)

	passwordSvc := infrastructure.NewPasswordService()
	jwtService := infrastructure.NewJWTService(os.Getenv("JWT_SECRET"))

	userUsecase := usecase.NewUserUseCase(userRepo, passwordSvc, jwtService)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)

	userController := controllers.NewUserController(userUsecase)
	taskController := controllers.NewTaskController(taskUsecase)

	r := routers.SetupRouter(taskController, userController, jwtService)

	r.Run(":8080")

}
