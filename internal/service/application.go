package service

import (
	"context"

	"github.com/juanmaabanto/go-user-ms/internal/app"
	"github.com/juanmaabanto/go-user-ms/internal/app/command"
	"github.com/juanmaabanto/go-user-ms/internal/app/query"
	"github.com/juanmaabanto/go-user-ms/internal/repository"
	"github.com/juanmaabanto/go-user-ms/pkg/mongo/database"
)

func NewApplication(ctx context.Context, dbname, uri string) (app.Application, error) {
	db, err := database.NewDatabase(ctx, dbname, uri)
	if err != nil {
		return app.Application{}, err
	}

	repos := repository.InitRepositories(db)

	return app.Application{
		Commands: app.Commands{
			CreateUser: command.NewCreateUserHandler(repos),
			UpdateUser: command.NewUpdateUserHandler(repos),
			DeleteUser: command.NewDeleteUserHandler(repos),
		},
		Queries: app.Queries{
			GetUserById: query.NewGetUserByIdHandler(repos),
		},
	}, nil
}
