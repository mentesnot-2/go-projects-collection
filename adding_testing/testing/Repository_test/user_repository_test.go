package testing

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/mentesnot-2/adding_testing/Domain"
	"github.com/mentesnot-2/adding_testing/Repository"
	"github.com/mentesnot-2/adding_testing/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"testing"
	"time"
)

type UserSuite struct {
	suite.Suite
	db   *testutils.TestDB
	repo repository.UserRepository
}

func (suite *UserSuite) SetupSuite() {
	err := godotenv.Load("../../.env")
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.db, err = testutils.SetupTestDB(os.Getenv("MONGO_URI"), os.Getenv("MONGO_DB"))
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.repo = repository.NewUserRepository(suite.db.DB)
}

func (suite *UserSuite) TearDownSuite() {
	err := suite.db.TearDown()
	if err != nil {
		suite.T().Fatal(err)
	}
}

func (suite *UserSuite) TestCreateUser() {
	user := domain.User{
		Username: "mentesnot",
		Password: "password",
		Email:    "test@gmail.com",
	}

	err := suite.repo.CreateUser(&user)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), user.ID)
	assert.WithinDuration(suite.T(), time.Now(), user.CreatedAt, 5*time.Second)
	assert.WithinDuration(suite.T(), time.Now(), user.UpdatedAt, 5*time.Second)
}

func (suite *UserSuite) TestGetUserByUsername() {
	expectedUser := domain.User{
		ID:       primitive.NewObjectID(),
		Username: "mentesnot",
		Password: "password",
		Email:    "test@gmail.com",
	}

	_, err := suite.db.DB.Collection("users").InsertOne(context.Background(), bson.M{
		"_id":      expectedUser.ID,
		"username": expectedUser.Username,
		"password": expectedUser.Password,
		"email":    expectedUser.Email,
	})
	assert.NoError(suite.T(), err)

	user, err := suite.repo.GetUserByUsername(expectedUser.Username)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedUser.Username, user.Username)
	assert.Equal(suite.T(), expectedUser.Password, user.Password)
	assert.Equal(suite.T(), expectedUser.Email, user.Email)
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
