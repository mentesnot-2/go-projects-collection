package testing

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/mentesnot-2/adding_testing/Domain"
	"github.com/mentesnot-2/adding_testing/Repository"
	"github.com/mentesnot-2/adding_testing/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"testing"
	"time"
)

type TaskSuite struct {
	suite.Suite
	db   *testutils.TestDB
	repo repository.TaskRepository
}

func (suite *TaskSuite) SetupSuite() {
	err := godotenv.Load("../../.env")
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.db, err = testutils.SetupTestDB(os.Getenv("MONGO_URI"), os.Getenv("MONGO_DB"))
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.repo = repository.NewTaskRepository(suite.db.DB)
}

func (suite *TaskSuite) TearDownSuite() {
	err := suite.db.TearDown()
	if err != nil {
		suite.T().Fatal(err)
	}
}

func (suit *TaskSuite) TestCreateTask() {
	task := domain.Task{
		Title:       "task",
		Description: "content",
	}
	err := suit.repo.CreateTask(&task)
	assert.NoError(suit.T(), err)
	assert.NotEmpty(suit.T(), task.ID)
	assert.WithinDuration(suit.T(), time.Now(), task.CreatedAt, 5*time.Second)
	assert.WithinDuration(suit.T(), time.Now(), task.UpdatedAt, 5*time.Second)
}

func (suite *TaskSuite) TestFindTaskById() {
	expectedTask := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "task",
		Description: "content",
	}
	_, err := suite.db.DB.Collection("tasks").InsertOne(context.Background(), expectedTask)
	assert.NoError(suite.T(), err)
	task, err := suite.repo.FindTaskById(expectedTask.ID.Hex())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedTask.ID, task.ID)
	assert.Equal(suite.T(), expectedTask.Title, task.Title)
	assert.Equal(suite.T(), expectedTask.Description, task.Description)
}

func (suite *TaskSuite) TestGetAllTask() {
	task1 := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "task1",
		Description: "content1",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	task2 := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "task2",
		Description: "content2",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	_, err := suite.db.DB.Collection("tasks").InsertMany(context.Background(), []interface{}{task1, task2})
	assert.NoError(suite.T(), err)
	tasks, err := suite.repo.GetAllTask()
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), tasks, 4)
}

func (suite *TaskSuite) TestUpdateTask() {
	task := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "task",
		Description: "content",
		Completed:   true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	_, err := suite.db.DB.Collection("tasks").InsertOne(context.Background(), task)
	assert.NoError(suite.T(), err)

	task.Description = "new content"
	err = suite.repo.UpdateTask(&task)
	assert.NoError(suite.T(), err)

	updatedTask, err := suite.repo.FindTaskById(task.ID.Hex())
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "new content", updatedTask.Description)
}

func (suite *TaskSuite) TestDeleteTask() {
	task := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "task",
		Description: "content",
		Completed:   true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	_, err := suite.db.DB.Collection("tasks").InsertOne(context.Background(), task)
	assert.NoError(suite.T(), err)

	err = suite.repo.DeleteTask(task.ID.Hex())
	assert.NoError(suite.T(), err)

	deletedTask, err := suite.repo.FindTaskById(task.ID.Hex())
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), deletedTask)
}

func TestTaskSuite(t *testing.T) {
	suite.Run(t, new(TaskSuite))
}
