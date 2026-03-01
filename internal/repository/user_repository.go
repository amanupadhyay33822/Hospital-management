package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UserRepository handles database operations for users.
type UserRepository struct {
	collection *mongo.Collection
}

// NewUserRepository creates a new UserRepository.
func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

// EnsureIndex creates a unique index on the email field.
func (r *UserRepository) EnsureIndex(ctx context.Context) error {
	var indexModel mongo.IndexModel
	indexModel = mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	var err error
	_, err = r.collection.Indexes().CreateOne(ctx, indexModel)
	return err
}
