package query

import (
	"context"
	"errors"
	"testing"

	"github.com/juanmaabanto/go-user-ms/internal/domain/model"
	"github.com/juanmaabanto/go-user-ms/internal/ports/response"
	"github.com/juanmaabanto/go-user-ms/internal/repository"
	"github.com/juanmaabanto/go-user-ms/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_Handle_GetUserById_Found(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepository[model.User])
	repos := repository.Repositories{
		User: mockRepo,
	}
	ctx := context.Background()
	expectedId := "6209a7fb0ceab9da565c546d"
	userId, _ := primitive.ObjectIDFromHex(expectedId)
	user := model.User{Id: userId, UserName: "test"}
	expected := response.UserResponse{Id: expectedId, UserName: "test"}

	mockRepo.On("FindById", ctx, expected.Id).Return(user, nil)

	// Act
	testQuery := NewGetUserByIdHandler(repos)
	result, err := testQuery.Handle(ctx, GetUserById{Id: expected.Id})

	// Assert
	mockRepo.AssertExpectations(t)

	assert.Equal(t, expected.Id, result.Id)
	assert.Equal(t, expected.UserName, result.UserName)
	assert.Nil(t, err)
}

func Test_Handler_GetUserById_Not_Found(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepository[model.User])
	repos := repository.Repositories{
		User: mockRepo,
	}
	ctx := context.Background()

	mockRepo.On("FindById", ctx, mock.AnythingOfType("string")).Return(nil, nil)

	// Act
	testQuery := NewGetUserByIdHandler(repos)
	userResponse, err := testQuery.Handle(ctx, GetUserById{Id: primitive.NewObjectID().String()})

	// Assert
	mockRepo.AssertExpectations(t)

	assert.NoError(t, err)
	assert.Nil(t, userResponse)
}

func Test_Handler_GetUserById_Error(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepository[model.User])
	repos := repository.Repositories{
		User: mockRepo,
	}
	ctx := context.Background()

	mockRepo.On("FindById", ctx, mock.AnythingOfType("string")).Return(nil, errors.New("An error has occurred"))

	// Act
	testQuery := NewGetUserByIdHandler(repos)
	_, err := testQuery.Handle(ctx, GetUserById{Id: primitive.NewObjectID().String()})

	// Assert
	mockRepo.AssertExpectations(t)

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "An error has occurred")
}
