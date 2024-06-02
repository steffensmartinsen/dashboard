package endpoints

import (
	"context"
	"dashboard/utils"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		getRegistration(w, r)
	case http.MethodPost:
		postRegistration(w, r)
	case http.MethodPut:
		putRegistration(w, r)
	case http.MethodDelete:
		deleteRegistration(w, r)
	}
}

// getRegistration is a function to handle GET requests to the registration endpoint
func getRegistration(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// postRegistration is a function to handle POST requests to the registration endpoint
func postRegistration(w http.ResponseWriter, r *http.Request) {

	// Instantiate a new decoder and a new response struct
	decoder := json.NewDecoder(r.Body)
	response := utils.UserRegistration{}

	// Decode the request into the response struct
	err := decoder.Decode(&response)
	if err != nil {
		http.Error(w, "Error decoding POST request", http.StatusBadRequest)
		log.Println("Error decoding Registration POST request")
		return
	}

	// Enforce all fields to be populated
	if response.Username == "" || response.Password == "" || response.Email == "" {
		http.Error(w, "Username, Password and Email must be provided", http.StatusBadRequest)
		log.Println("Username, Password or Email is empty")
		return
	}

	// Enforce a password only containing characters '1234567890'
	if !utils.EnforceShittyPassword(response.Password) {
		http.Error(w, "Please don't use an actual password for this. The only accepted characters are '1234567890'", http.StatusBadRequest)
		log.Println("Password not allowed")
		return
	}

	// Seperate check for password length
	if len(response.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters long", http.StatusBadRequest)
		log.Println("Password too short")
		return
	}

	// Insert the user into the database
	collection := utils.Client.Database("users").Collection(utils.COLLECTION_USERS)
	_, err = collection.InsertOne(context.TODO(), response)
	if err != nil {

		// Check if the error du to an already existing user or email
		if writeException, ok := err.(mongo.WriteException); ok {
			for _, writeError := range writeException.WriteErrors {
				if writeError.Code == 11000 {
					http.Error(w, "Username or Email already exists", http.StatusBadRequest)
					log.Println("Username or Email already exists")
					return
				}
			}
		}

		// If the error is not due to an already existing user or email, return a generic error
		http.Error(w, "Error inserting user", http.StatusInternalServerError)
		log.Println("Error inserting user")
		return

	}

	// Set response header to JSON and return status code 201
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	log.Println("User '" + response.Username + "' registered.")
}

func putRegistration(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func deleteRegistration(w http.ResponseWriter, r *http.Request) {
	// TODO
}
