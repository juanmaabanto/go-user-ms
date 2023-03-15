package command

import (
	"context"
	"time"

	"github.com/juanmaabanto/go-user-ms/internal/domain/model"
	"github.com/juanmaabanto/go-user-ms/internal/repository"
)

type CreateUser struct {
	UserName    string    `json:"userName" validate:"required,max=10"`
	FirstName   string    `json:"firstName" validate:"required,max=20"`
	LastName    string    `json:"lastName" validate:"required,max=20"`
	Email       string    `json:"email" validate:"required,max=50"`
	PhoneNumber string    `json:"phoneNumber" validate:"required,max=50"`
	BirthDate   time.Time `json:"birthDate" validate:"required"`
}

type CreateUserHandler struct {
	repos repository.Repositories
}

func NewCreateUserHandler(repos repository.Repositories) CreateUserHandler {
	return CreateUserHandler{repos: repos}
}

func (h CreateUserHandler) Handle(ctx context.Context, command CreateUser) (string, error) {
	user := model.User{}

	user.UserName = command.UserName
	user.FirstName = command.FirstName
	user.LastName = command.LastName
	user.Email = command.Email
	user.PhoneNumber = command.PhoneNumber
	user.BirthDate = command.BirthDate
	user.CreatedAt = time.Now()
	user.CreatedBy = "admin"

	id, err := h.repos.User.InsertOne(ctx, user)
	if err != nil {
		return id, err
	}

	return id, nil
}
