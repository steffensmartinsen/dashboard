package main

import (
	"context"
	"dashboard/database"
	"dashboard/endpoints"
	"dashboard/utils"
	"log"
	"net/http"
	"os"
)

// Needed?
var Ctx context.Context

func main() {

	log.Println("Starting server...")

	// Instantiate the connection to the MongoDB database and MongoDB client
	Client := utils.DBConnect()
	db := database.NewMongoDB(Client, utils.COLLECTION_USERS, utils.COLLECTION_USERS)

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

	// Endpoint handlers
	http.HandleFunc("/", endpoints.EmptyHandler)
	http.HandleFunc(utils.DEFAULT_PATH, endpoints.EmptyHandler)
	http.HandleFunc(utils.PATH_REGISTRATIONS, func(w http.ResponseWriter, r *http.Request) {
		endpoints.RegistrationHandler(db, w, r)
	})
	http.HandleFunc(utils.ENDPOINT_AUTHENTICATION, func(w http.ResponseWriter, r *http.Request) {
		endpoints.AuthenticationHandler(db, w, r)
	})

	// Starting server
	log.Println("Starting server on port " + port + "\n")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
