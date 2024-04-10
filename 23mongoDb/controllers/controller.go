package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	netflixModel "github.com/abhi-alt/mongodb-connect/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// insert one record
func insertOneMovie(movie netflixModel.Netflix) primitive.ObjectID {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	id := inserted.InsertedID.(primitive.ObjectID)
	fmt.Println("Inserted one movie with id: ", inserted.InsertedID)
	return id
}

// update one record
func updateOneRecord(id string) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"movie": "Avengers Endgame"}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Recor is updated and updated count is: ", result.ModifiedCount)
}

// delete one record
func deleteOneRecord(id string) int64 {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": objectId}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("record deleted ", result.DeletedCount)
	return result.DeletedCount
}

// delete all
func deleteAllMovie() int64 {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("all records deleted ", result.DeletedCount)
	return result.DeletedCount
}

func getAllMovies() []bson.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var movies []bson.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies
}

//actual controllers

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movies = getAllMovies()
	json.NewEncoder(w).Encode(movies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie netflixModel.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}
	id := insertOneMovie(movie)
	movie.ObjectId = id
	json.NewEncoder(w).Encode(movie)

}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneRecord(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	count := deleteOneRecord(params["id"])
	json.NewEncoder(w).Encode(count)
}

func DeleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
