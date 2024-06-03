package endpoints

import (
	"context"
	"dashboard/utils"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"strings"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {

	utils.EnsureCorrectPath(r)

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

	// Extract the username from the URL path
	path := strings.Split(r.URL.Path, "/")
	username := path[len(path)-2]
	log.Println("Username: " + username)

	// Enforce username to be specified
	if username == "registrations" {
		http.Error(w, "Username must be provided", http.StatusBadRequest)
		log.Println("Username not provided")
		return
	}

	// Open the collection, set content-type to JSON and instantiate a response struct
	collection := utils.Client.Database(utils.COLLECTION_USERS).Collection(utils.COLLECTION_USERS)
	w.Header().Add("Content-Type", "application/json")
	response := utils.UserRegistration{}

	// Attempt to find the user in the database
	err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&response)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		log.Println("User not found")
		return
	}

	// Encode the response struct to the client
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		log.Println("Error encoding response")
		return
	}
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

	// Enforce username and email to be lowercase
	response.Username = strings.ToLower(response.Username)
	response.Email = strings.ToLower(response.Email)

	// Open the collection
	collection := utils.Client.Database(utils.COLLECTION_USERS).Collection(utils.COLLECTION_USERS)

	// Check if the username or email already exists
	existingUser := utils.UserRegistration{}
	err = collection.FindOne(context.TODO(), bson.M{"$or": []bson.M{
		{"username": response.Username},
		{"email": response.Email},
	}}).Decode(&existingUser)

	if err == nil {
		http.Error(w, "Username or Email already exists", http.StatusBadRequest)
		log.Println("Username or Email already exists")
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

	// Hash the password
	response.Password, err = utils.HashPassword(response.Password)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		log.Println("Error hashing password")
		return
	}

	// Insert the user into the database
	_, err = collection.InsertOne(context.TODO(), response)
	if err != nil {
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
