package app

import (
	"github.com/juanmaabanto/go-user-ms/internal/app/command"
	"github.com/juanmaabanto/go-user-ms/internal/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateUser command.CreateUserHandler
}

type Queries struct {
	GetUserById query.GetUserByIdHandler
}
