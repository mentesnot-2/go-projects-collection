package testutils

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TestDB struct {
	Client *mongo.Client
	DB     *mongo.Database
	Coll   *mongo.Collection
}

func SetupTestDB(uri string, dbName string) (*TestDB, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	return &TestDB{Client: client, DB: db}, nil
}

func (db *TestDB) TearDown() error {
	err := db.DB.Collection("tasks").Drop(context.Background())
	if err != nil {
		return err
	}
	return db.Client.Disconnect(context.Background())
}
