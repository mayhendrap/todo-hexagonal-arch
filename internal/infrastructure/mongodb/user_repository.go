package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"todo-hexagonal-arch/internal/core/domain"
	"todo-hexagonal-arch/internal/core/interfaces"
)

type userRepository struct {
	db         *mongo.Database
	collection string
}

func NewUserRepository(db *mongo.Database, collection string) interfaces.UserRepository {
	return &userRepository{db, collection}
}

func (ur *userRepository) FindByEmail(email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := ur.db.Collection(ur.collection)

	var user domain.User

	err := collection.FindOne(ctx, bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		return user, errors.New("couldn't find user with email: " + email)
	}
	return user, err
}

func (ur *userRepository) Create(user domain.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := ur.db.Collection(ur.collection)

	userID := primitive.NewObjectID().Hex()
	user.ID = userID

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (ur *userRepository) Update(user domain.User) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (ur *userRepository) Delete(userID string) error {
	//TODO implement me
	panic("implement me")
}
