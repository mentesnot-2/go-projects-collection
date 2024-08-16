package usecasetest

import (
	"testing"
	"time"
	"github.com/mentesnot-2/adding_testing/Domain"
	"github.com/mentesnot-2/adding_testing/Usecase"
	"github.com/mentesnot-2/adding_testing/testing/mocks/Repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func Test_CreateTask(t *testing.T){
	mockTaskRepo := new(repository.TaskRepository)
	taskUsecase:= usecase.NewTaskUsecase(mockTaskRepo)


	taskId:= primitive.NewObjectID()
	userId:= primitive.NewObjectID()
	task:= &domain.Task{
		ID: taskId,
		UserID: userId,
		Title: "testtask",
		Description: "testdescription",
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockTaskRepo.On("CreateTask", task).Return(nil)
	err := taskUsecase.CreateTask(task)
	assert.NoError(t, err)
	mockTaskRepo.AssertExpectations(t)
}


func Test_GetALLTasks(t *testing.T) {
	mockTaskRepo := new(repository.TaskRepository)
	taskUsecase := usecase.NewTaskUsecase(mockTaskRepo)

	tasks := []*domain.Task{
		{
			ID:primitive.NewObjectID(),
			UserID: primitive.NewObjectID(),
			Title: "Testtask1",
			Description: "Testdescription1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:primitive.NewObjectID(),
			UserID: primitive.NewObjectID(),
			Title: "Testtask2",
			Description: "Testdescription2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	mockTaskRepo.On("GetAllTask").Return(tasks,nil)
	result,err:= taskUsecase.GetAllTask()

	assert.NoError(t,err)
	assert.Equal(t,tasks,result)
	mockTaskRepo.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
	mockRepo := new(repository.TaskRepository)
	taskUsecase := usecase.NewTaskUsecase(mockRepo)
	taskId:= primitive.NewObjectID()
	userId:= primitive.NewObjectID()
	
	task := &domain.Task{
		ID:          taskId,
		UserID:      userId,
		Title:       "Original Title",
		Description: "Original Description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	
	updatedTask := &domain.Task{
		ID:          taskId,
		UserID:      userId,
		Title:       "Updated Title",
		Description: "Updated Description",
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   time.Now(),
	}
	mockRepo.On("CreateTask", task).Return(nil)
	err:= taskUsecase.CreateTask(task)
	assert.NoError(t,err)

	mockRepo.On("UpdateTask", updatedTask).Return(nil)

	err = taskUsecase.UpdateTask(updatedTask)
	assert.NoError(t, err)
	
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	mockRepo := new(repository.TaskRepository)
	taskUsecase := usecase.NewTaskUsecase(mockRepo)

	t.Run("success", func(t *testing.T) {
		taskID := primitive.NewObjectID()
		userID := primitive.NewObjectID()
		task := &domain.Task{
			ID:          taskID,
			UserID:      userID,
			Title:       "Test Task",
			Description: "Test Description",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		mockRepo.On("CreateTask", task).Return(nil)
		err := taskUsecase.CreateTask(task)
		assert.NoError(t, err)


		mockRepo.On("DeleteTask", task.ID.Hex()).Return(nil)

		err = taskUsecase.DeleteTask(task.ID.Hex())
		assert.NoError(t, err)
	
		mockRepo.AssertExpectations(t)
	})
	t.Run("error", func(t *testing.T) {
		taskID := primitive.NewObjectID()
		userID := primitive.NewObjectID()
		task := &domain.Task{
			ID:          taskID,
			UserID:      userID,
			Title:       "Test Task",
			Description: "Test Description",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		mockRepo.On("CreateTask", task).Return(nil)
		err := taskUsecase.CreateTask(task)
		assert.NoError(t, err)

		mockRepo.On("DeleteTask", taskID.Hex()).Return(assert.AnError)
		err = taskUsecase.DeleteTask(taskID.Hex())
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}