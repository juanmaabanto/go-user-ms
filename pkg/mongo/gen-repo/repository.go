package genrepo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GRepository[D Document] struct {
	Database *mongo.Database
}

func (repo GRepository[D]) FindById(ctx context.Context, id string) (*D, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := repo.GetCollection().FindOne(ctx, bson.D{{Key: "_id", Value: objID}})

	if result.Err() != nil && result.Err() != mongo.ErrNoDocuments {
		return nil, result.Err()
	}

	if result.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}

	var document D
	err = result.Decode(&document)
	if err != nil {
		return nil, err
	}
	return &document, nil
}

func (repo GRepository[D]) InsertOne(ctx context.Context, document D) (string, error) {
	result, err := repo.GetCollection().InsertOne(ctx, document)

	return result.InsertedID.(primitive.ObjectID).Hex(), err
}

func (repo GRepository[D]) RemoveOne(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = repo.GetCollection().DeleteOne(ctx, bson.D{{Key: "_id", Value: objID}})
	return err
}

func (repo GRepository[D]) UpdateOne(ctx context.Context, id string, document D) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = repo.GetCollection().UpdateOne(ctx, bson.D{{Key: "_id", Value: objID}}, bson.M{"$set": document})
	return err
}

func (repo GRepository[D]) GetCollection() *mongo.Collection {
	return repo.Database.Collection((*new(D)).GetCollectionName())
}
