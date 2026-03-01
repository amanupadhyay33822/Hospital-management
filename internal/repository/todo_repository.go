package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// TodoRepository handles database operations for todos.
type TodoRepository struct {
	collection *mongo.Collection
}

// NewTodoRepository creates a new TodoRepository.
func NewTodoRepository(db *mongo.Database) *TodoRepository {
	return &TodoRepository{
		collection: db.Collection("todos"),
	}
}
