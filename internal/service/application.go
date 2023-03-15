package service

import (
	"context"
	"log"

	"github.com/juanmaabanto/go-user-ms/internal/app"
	"github.com/juanmaabanto/go-user-ms/internal/repository"
	"github.com/juanmaabanto/go-user-ms/pkg/mongo/database"
)

func NewApplication(ctx context.Context, dbname, uri string) (app.Application, error) {
	db, err := database.NewDatabase(ctx, dbname, uri)
	if err != nil {
		return app.Application{}, err
	}

	repos := repository.InitRepositories(db)
	log.Print(repos)

	return app.Application{
		Commands: app.Commands{},
		Queries:  app.Queries{},
	}, nil
}
