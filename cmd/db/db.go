package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var User *mongo.Collection
var Todo *mongo.Collection

func InitDatabase(DB string) error {
	client, err := mongo.Connect(context.TODO(),options.Client().ApplyURI(DB))
	if err != nil {
		return err
	}
	
	db := client.Database("GOLANG")

	User = db.Collection("user")
	Todo = db.Collection("todo")

	log.Println("Database connect successful")
	return nil
}