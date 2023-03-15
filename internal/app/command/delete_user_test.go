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

func Test_Handle_DeleteUser_Find_Error(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepository[model.User])
	repos := repository.Repositories{
		User: mockRepo,
	}
	ctx := context.Background()
	item := DeleteUser{Id: "6209a7fb0ceab9da565c546d"}

	mockRepo.On("FindById", ctx, mock.AnythingOfType("string")).Return(nil, errors.New("An error has occurred"))

	// Act
	testCommand := NewDeleteUserHandler(repos)
	err := testCommand.Handle(ctx, item)

	// Assert
	mockRepo.AssertExpectations(t)

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "An error has occurred")
}

func Test_Handler_DeleteUser_Not_Found(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepository[model.User])
	repos := repository.Repositories{
		User: mockRepo,
	}
	ctx := context.Background()
	item := DeleteUser{Id: "6209a7fb0ceab9da565c546d"}

	mockRepo.On("FindById", ctx, mock.AnythingOfType("string")).Return(nil, nil)

	// Act
	testCommand := NewDeleteUserHandler(repos)
	err := testCommand.Handle(ctx, item)

	// Assert
	mockRepo.AssertExpectations(t)

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "user not found")
}

func Test_Handler_DeleteUser_Error_Completed(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepository[model.User])
	repos := repository.Repositories{
		User: mockRepo,
	}
	ctx := context.Background()
	expectedId := "6209a7fb0ceab9da565c546d"
	userId, _ := primitive.ObjectIDFromHex(expectedId)
	user := model.User{Id: userId, UserName: "test"}
	item := DeleteUser{Id: expectedId}

	mockRepo.On("FindById", ctx, mock.AnythingOfType("string")).Return(user, nil)
	mockRepo.On("RemoveOne", ctx, item.Id).Return(nil)

	// Act
	testCommand := NewDeleteUserHandler(repos)
	err := testCommand.Handle(ctx, item)

	// Assert
	mockRepo.AssertExpectations(t)

	assert.NoError(t, err)
}
