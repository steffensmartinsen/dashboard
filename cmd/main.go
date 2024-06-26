package main

import (
	"context"
	"dashboard/database"
	"dashboard/endpoints"
	"dashboard/utils"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

// Needed?
var Ctx context.Context

func main() {

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
	http.HandleFunc(utils.SLASH, endpoints.EmptyHandler)
	http.HandleFunc(utils.DEFAULT_PATH, endpoints.EmptyHandler)
	http.HandleFunc(utils.PATH_REGISTRATIONS, func(w http.ResponseWriter, r *http.Request) {
		endpoints.RegistrationHandler(db, w, r)
	})
	http.HandleFunc(utils.PATH_AUTHENTICATION, func(w http.ResponseWriter, r *http.Request) {
		endpoints.AuthenticationHandler(db, w, r)
	})
	http.HandleFunc(utils.PATH_WEATHER, func(w http.ResponseWriter, r *http.Request) {
		endpoints.WeatherHandler(db, w, r)
	})
	http.HandleFunc(utils.PATH_SET_COOKIE, endpoints.SetCookie)
	http.HandleFunc(utils.PATH_GET_COOKIE, endpoints.GetCookie)
	http.HandleFunc(utils.PATH_DELETE_COOKIE, endpoints.DeleteCookie)

	// Starting server
	log.Println("Starting server on port " + port + "\n")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
