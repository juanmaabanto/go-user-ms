package repotypes

import (
	"github.com/juanmaabanto/go-user-ms/internal/domain/model"
	genrepo "github.com/juanmaabanto/go-user-ms/pkg/mongo/gen-repo"
)

type UserRepository interface {
	genrepo.Repository[model.User]
}
