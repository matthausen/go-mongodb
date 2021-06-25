package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/matthausen/thirdfort/model"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const (
	connectionString = "mongodb+srv://admin:thirdfort123@thirdfortcluster.4twv3.mongodb.net/notes?retryWrites=true&w=majority"
)

var collection *mongo.Collection

// Initialise the MongoDB client
func init() {
	clientOptions := options.Client().ApplyURI(connectionString)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	// Ping the DB
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatalf("Could not ping db: %v", err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database("Thirdfort").Collection("notes")

	fmt.Println("Collection instance created!")
}

// enableCors -> makes sure we can run the requests from any url
func enableCors(w http.ResponseWriter) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// SaveNote -> Insert a single new note into the DB
func SaveNote(w http.ResponseWriter, r *http.Request) {
	enableCors(w)

	var note model.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Fatalf("Could not decode the body of the request: %v", err)
	}

	res, err := collection.InsertOne(context.Background(), note)
	if err != nil {
		log.Fatalf("Could not save new note: %v", err)
	}
	fmt.Printf("New note saved: %s", res.InsertedID)
	json.NewEncoder(w).Encode(note)
}

// UpdateNote -> update a previously created note. Both its text and it archived status can be updated
func UpdateNote(w http.ResponseWriter, r *http.Request) {
	enableCors(w)

	noteId := r.URL.Query().Get("id")

	var note model.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Fatalf("Could not decode the body of the request: %v\n", err)
	}

	id, _ := primitive.ObjectIDFromHex(noteId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"text": note.Text,
		"archived": note.Archived,
		},
	}
	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Could not update note: %v\n", err)
	}

	json.NewEncoder(w).Encode(res.ModifiedCount)

}

// ListAllSaved -> fetch all the saved notes (non archived)
func ListAllSaved(w http.ResponseWriter, r *http.Request) {
	enableCors(w)

	filter, err := collection.Find(context.Background(), bson.M{"archived": false})
	if err != nil {
		log.Fatalf("Could not fetch archived notes: %v\n", err)
	}
	var savedNotes []bson.M
	if err = filter.All(context.Background(), &savedNotes); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(savedNotes)
}

// DeleteNote -> delete one note from db
func DeleteNote(w http.ResponseWriter, r *http.Request) {

}

// ListAllArchived -> List all archived notes
func ListAllArchived(w http.ResponseWriter, r *http.Request) {

	filter, err := collection.Find(context.Background(), bson.M{"archived": true})
	if err != nil {
		log.Fatalf("Could not fetch archived notes: %v\n", err)
	}
	var archivedNotes []bson.M
	if err = filter.All(context.Background(), &archivedNotes); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(archivedNotes)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("All good")
}
