package main

import (
	"context"
	"dashboard/endpoints"
	"dashboard/utils"
	"log"
	"net/http"
	"os"
)

func main() {

	// Instantiate the connection to the MongoDB database
	utils.DBConnect()

	// Disconnect from MongoDB when the service is closed
	defer func() {
		if err := utils.Client.Disconnect(context.TODO()); err != nil {
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

	// Starting server
	log.Println("Starting server on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
