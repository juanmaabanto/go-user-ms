package genrepo

import "context"

type Document interface {
	GetCollectionName() string
}

type Repository[D Document] interface {
	FindById(context.Context, string) (*D, error)
	InsertOne(context.Context, D) (string, error)
	RemoveOne(context.Context, string) error
	UpdateOne(context.Context, string, D) error
}
