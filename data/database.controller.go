// Package data database.controller.go
package data

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func Init() {
    const mongoURI = "mongodb://mongo:27019"
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
    if err != nil {
        panic(err)
    }

    MongoClient = client
}
