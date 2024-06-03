package utils

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

var Client *mongo.Client

// DBConnect Function to connect to the MongoDB database
func DBConnect() {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://viatheyboy:uvL3rwoDOrt42G8g@dashboards.ylq2uxl.mongodb.net/?retryWrites=true&w=majority&appName=dashboards").SetServerAPIOptions(serverAPI)

	var err error
	// Create a new client and connect to the server
	Client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Get a handle to the users connection
	collection := Client.Database(COLLECTION_USERS).Collection(COLLECTION_USERS)

	// Establish uniqueness constraint on username and email
	usernameIndex := mongo.IndexModel{
		Keys:    bson.M{"username": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err = collection.Indexes().CreateOne(context.TODO(), usernameIndex)
	if err != nil {
		log.Fatal(err)
	}

	emailIndex := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err = collection.Indexes().CreateOne(context.TODO(), emailIndex)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully created indexes on username and email fields.")

	// Send a ping to confirm a successful connection
	if err := Client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("\nService successfully connected to MongoDB.\n")
}

// EnforceShittyPassword Function to ensure users don't use an actual password
func EnforceShittyPassword(password string) bool {

	accepted := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

	for _, char := range password {
		found := false
		for _, element := range accepted {
			if string(char) == element {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// HashPassword Function to hash a password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword Function to check if a password matches a hash
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// EnsureCorrectPath appends a '/' to the path if it is missing
func EnsureCorrectPath(r *http.Request) {
	if r.URL.Path[len(r.URL.Path)-1] != '/' {
		r.URL.Path += "/"
	}
}
