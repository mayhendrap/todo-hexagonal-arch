package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    context.Context
}

func NewMongoDBAdapter(databaseName, dataSourceName string) *MongoDB {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dataSourceName))
	if err != nil {
		log.Fatalf("[x] cant connect to mongodb: %v", err)
	}

	db := client.Database(databaseName)

	return &MongoDB{
		client: client,
		db:     db,
	}
}

func (a *MongoDB) CloseDBConnection() {
	err := a.client.Disconnect(a.ctx)
	if err != nil {
		log.Fatalf("[x] fail to disconnect mongodb: %v", err)
	}
}

func (a *MongoDB) GetDB() *mongo.Database {
	return a.db
}
