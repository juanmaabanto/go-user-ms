package repository

import (
	"github.com/juanmaabanto/go-user-ms/internal/domain/repotypes"
	"github.com/juanmaabanto/go-user-ms/internal/repository/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	User repotypes.UserRepository
}

func InitRepositories(db *mongo.Database) Repositories {
	return Repositories{
		User: user.New(db),
	}
}
