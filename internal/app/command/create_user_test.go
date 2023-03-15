package command

import (
	"context"
	"errors"
	"testing"

	"github.com/juanmaabanto/go-user-ms/internal/domain/model"
	"github.com/juanmaabanto/go-user-ms/internal/repository"
	"github.com/juanmaabanto/go-user-ms/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_Handle_CreateUser_Insert_Completed(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepository[model.User])
	repos := repository.Repositories{
		User: mockRepo,
	}

	ctx := context.Background()
	item := CreateUser{UserName: "test", FirstName: "first test"}
	expected := primitive.NewObjectID().Hex()

	mockRepo.On("InsertOne", ctx, mock.AnythingOfType("model.User")).Return(expected, nil)

	// Act
	testCommand := NewCreateUserHandler(repos)
	result, err := testCommand.Handle(ctx, item)

	// Assert

	mockRepo.AssertExpectations(t)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func Test_Handle_CreateMeasureType_Insert_Error(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepository[model.User])
	repos := repository.Repositories{
		User: mockRepo,
	}
	ctx := context.Background()
	item := CreateUser{UserName: "test", FirstName: "first test"}

	mockRepo.On("InsertOne", ctx, mock.AnythingOfType("model.User")).Return("", errors.New("An error has occurred"))

	// Act
	testCommand := NewCreateUserHandler(repos)
	_, err := testCommand.Handle(ctx, item)

	// Assert

	mockRepo.AssertExpectations(t)

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "An error has occurred")
}
