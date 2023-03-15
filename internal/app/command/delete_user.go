package command

import (
	"context"
	"errors"

	"github.com/juanmaabanto/go-user-ms/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteUser struct {
	Id string
}

type DeleteUserHandler struct {
	repos repository.Repositories
}

func NewDeleteUserHandler(repos repository.Repositories) DeleteUserHandler {
	return DeleteUserHandler{repos: repos}
}

func (h DeleteUserHandler) Handle(ctx context.Context, command DeleteUser) error {
	existent, err := h.repos.User.FindById(ctx, command.Id)
	if err != nil {
		return err
	}

	if existent == nil || existent.Id == primitive.NilObjectID {
		return errors.New("user not found")
	}

	return h.repos.User.RemoveOne(ctx, command.Id)
}
