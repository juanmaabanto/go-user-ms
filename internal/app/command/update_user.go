package command

import (
	"context"
	"errors"
	"time"

	"github.com/juanmaabanto/go-user-ms/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateUser struct {
	Id          string    `json:"id" validate:"required,max=24,min=24"`
	UserName    string    `json:"userName" validate:"required,max=10"`
	FirstName   string    `json:"firstName" validate:"required,max=20"`
	LastName    string    `json:"lastName" validate:"required,max=20"`
	Email       string    `json:"email" validate:"required,max=50"`
	PhoneNumber string    `json:"phoneNumber" validate:"required,max=50"`
	BirthDate   time.Time `json:"birthDate" validate:"required"`
}

type UpdateUserHandler struct {
	repos repository.Repositories
}

func NewUpdateUserHandler(repos repository.Repositories) UpdateUserHandler {
	return UpdateUserHandler{repos: repos}
}

func (h UpdateUserHandler) Handle(ctx context.Context, command UpdateUser) error {
	existent, err := h.repos.User.FindById(ctx, command.Id)
	if err != nil {
		return err
	}

	if existent == nil || existent.Id == primitive.NilObjectID {
		return errors.New("user not found")
	}

	existent.UserName = command.UserName
	existent.FirstName = command.FirstName
	existent.LastName = command.LastName
	existent.Email = command.Email
	existent.PhoneNumber = command.PhoneNumber
	existent.BirthDate = command.BirthDate

	return h.repos.User.UpdateOne(ctx, command.Id, *existent)
}
