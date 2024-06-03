package main

import (
	"context"
	"dashboard/endpoints"
	"dashboard/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"os"
)

var Client *mongo.Client
var Ctx context.Context

func main() {

	// Instantiate the connection to the MongoDB database
	utils.DBConnect()

	// Disconnect from MongoDB when the service is closed
	defer func() {
		if err := Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Check for ENV variable port, if none we set it to default 8080
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set, using default port " + utils.DEFAULT_PORT)
		port = utils.DEFAULT_PORT
	}

	http.HandleFunc("/", endpoints.EmptyHandler)
	http.HandleFunc(utils.DEFAULT_PATH, endpoints.EmptyHandler)
	http.HandleFunc(utils.PATH_REGISTRATIONS, endpoints.RegistrationHandler)

	// Starting server
	log.Println("Starting server on port " + port + "\n")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
