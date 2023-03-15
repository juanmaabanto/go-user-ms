package user

import (
	"github.com/juanmaabanto/go-user-ms/internal/domain/model"
	"github.com/juanmaabanto/go-user-ms/internal/domain/repotypes"
	genrepo "github.com/juanmaabanto/go-user-ms/pkg/mongo/gen-repo"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	genrepo.GRepository[model.User]
}

func New(db *mongo.Database) repotypes.UserRepository {
	return repository{
		GRepository: genrepo.GRepository[model.User]{
			Database: db,
		},
	}
}
