package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo = InitMongo()

func InitMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "song",
		Password: "242166", // define.MongoPassword
	}).ApplyURI("mongodb://81.69.202.167:27017"))
	if err != nil {
		log.Println("Connection MongoDB Error:", err)
		return nil
	}
	return client.Database("chant")
}
