package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
	"todo-hexagonal-arch/internal/core/domain"
	"todo-hexagonal-arch/internal/core/interfaces"
)

type todoRepository struct {
	db         *mongo.Database
	collection string
}

func NewTodoRepository(db *mongo.Database, collection string) interfaces.TodoRepository {
	return &todoRepository{db, collection}
}

func (tr *todoRepository) FindOne(todoID, userID string) (domain.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := tr.db.Collection(tr.collection)

	var todo domain.Todo

	log.Printf("[repository] TODO ID: %s, USER ID: %s", todoID, userID)

	objectID, err := primitive.ObjectIDFromHex(todoID)
	if err != nil {
		return todo, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": objectID, "user_id": userID}).Decode(&todo)
	if err != nil {
		return todo, errors.New("couldn't find todo with id: " + todoID)
	}
	return todo, err
}

func (tr *todoRepository) FindAll(userID string) ([]domain.Todo, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := tr.db.Collection(tr.collection)

	var todos []domain.Todo

	cursor, err := collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return todos, err
	}

	err = cursor.All(ctx, &todos)
	return todos, err
}

func (tr *todoRepository) Create(todo domain.Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := tr.db.Collection(tr.collection)

	_, err := collection.InsertOne(ctx, todo)
	return err
}

func (tr *todoRepository) Update(todo domain.Todo) (domain.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := tr.db.Collection(tr.collection)

	todoID, err := primitive.ObjectIDFromHex(todo.ID)
	if err != nil {
		return domain.Todo{}, err
	}

	filter := bson.D{
		{"_id", todoID},
		{"user_id", todo.UserID},
	}

	update := bson.D{{"$set", bson.D{
		{"title", todo.Title},
		{"content", todo.Content},
	}}}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return domain.Todo{}, err
	}

	result, err := tr.FindOne(todo.ID, todo.UserID)
	if err != nil {
		return domain.Todo{}, err
	}

	return result, nil
}

func (tr *todoRepository) Delete(todoID, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := tr.db.Collection(tr.collection)

	objectID, err := primitive.ObjectIDFromHex(todoID)
	if err != nil {
		return err
	}

	tood, _ := tr.FindOne(todoID, userID)
	log.Println("TODO: ", tood)

	_, err = collection.DeleteOne(ctx, bson.D{
		{"_id", objectID},
		{"user_id", userID},
	})
	if err != nil {
		return err
	}

	return nil
}
