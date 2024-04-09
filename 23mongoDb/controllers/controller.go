package controllers

import (
	"context"
	"fmt"
	"log"

	netflixModel "github.com/abhi-alt/mongodb-connect/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	//client options
	clientOptions := options.Client().ApplyURI(connectionString)

	//connect to mongoDb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongdb connection established")

	//get instance of collection
	collection = client.Database(dbName).Collection(colName)

	fmt.Println("collection instance received")
}

func insertOneMovie(movie netflixModel.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie with id: ", inserted.InsertedID)
}
