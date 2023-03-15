package mocks

import (
	"context"

	genrepo "github.com/juanmaabanto/go-user-ms/pkg/mongo/gen-repo"
	"github.com/stretchr/testify/mock"
)

type MockRepository[D genrepo.Document] struct {
	mock.Mock
}

func (mock MockRepository[D]) FindById(ctx context.Context, id string) (*D, error) {
	args := mock.Called(ctx, id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}

	model := result.(D)
	return &model, args.Error(1)
}

func (mock MockRepository[D]) InsertOne(ctx context.Context, document D) (string, error) {
	args := mock.Called(ctx, document)
	result := args.Get(0)

	return result.(string), args.Error(1)
}

func (mock MockRepository[D]) RemoveOne(ctx context.Context, id string) error {
	args := mock.Called(ctx, id)

	return args.Error(0)
}

func (mock MockRepository[D]) UpdateOne(ctx context.Context, id string, document D) error {
	args := mock.Called(ctx, id, document)

	return args.Error(0)
}
