package endpoints

import (
	"context"
	"dashboard/utils"
	"encoding/json"
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

func getRegistration(w http.ResponseWriter, r *http.Request) {
	// TODO
}

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

	// Enforce that the username is unique
	if utils.UserExists(response.Username) {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		log.Println("Username already exists")
		return
	}

	collection := utils.Client.Database("users").Collection(utils.COLLECTION_USERS)
	insertResult, err := collection.InsertOne(context.TODO(), response)
	if err != nil {
		http.Error(w, "Error inserting user", http.StatusInternalServerError)
		log.Println("Error inserting user")
		return

	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated) // Set the status code to 201 Created
	log.Println("User '" + response.Username + "' registered.")
	log.Println(insertResult.InsertedID)

}

func putRegistration(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func deleteRegistration(w http.ResponseWriter, r *http.Request) {
	// TODO
}
