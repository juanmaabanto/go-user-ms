package query

import (
	"context"

	"github.com/juanmaabanto/go-user-ms/internal/ports/response"
	"github.com/juanmaabanto/go-user-ms/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetUserById struct {
	Id string
}

type GetUserByIdHandler struct {
	repos repository.Repositories
}

func NewGetUserByIdHandler(repos repository.Repositories) GetUserByIdHandler {
	return GetUserByIdHandler{repos: repos}
}

func (h GetUserByIdHandler) Handle(ctx context.Context, query GetUserById) (*response.UserResponse, error) {
	user, err := h.repos.User.FindById(ctx, query.Id)
	if err != nil {
		return nil, err
	}

	if user == nil || user.Id == primitive.NilObjectID {
		return nil, nil
	}

	return &response.UserResponse{
		Id:          user.Id.Hex(),
		UserName:    user.UserName,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		BirthDate:   user.BirthDate,
	}, nil
}
