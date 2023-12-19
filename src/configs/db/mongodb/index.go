package mongodb

import (
	"context"
	"log"

	"github.com/lehodev/golang-server/src/configs/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func Connect() error {
	uri := env.MyEnv.MONGODB_URI
	dbName := env.MyEnv.MONGODB_DB

	clientOptions := options.Client().ApplyURI(uri)
	db, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting to MongoDB:", err)
		return err
	}

	err = db.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Error pinging MongoDB:", err)
		return err
	}

	log.Println("---> Connected to MongoDB! <----")
	DB = db.Database(dbName)
	return nil
}

func Disconnect() error {
	log.Println("---> Disconnected from MongoDB! <----")
	return DB.Client().Disconnect(context.TODO())
}
