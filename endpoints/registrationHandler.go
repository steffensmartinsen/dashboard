package endpoints

import (
	"dashboard/database"
	"dashboard/utils"
	"encoding/json"
	"log"
	"net/http"
)

func RegistrationHandler(db database.Database, w http.ResponseWriter, r *http.Request) {

	utils.EnsureCorrectPath(r)

	switch r.Method {
	case http.MethodGet:
		getRegistration(db, w, r)
	case http.MethodPost:
		postRegistration(db, w, r)
	case http.MethodPut:
		putRegistration(db, w, r)
	case http.MethodDelete:
		deleteRegistration(db, w, r)
	}
}

// getRegistration is a function to handle GET requests to the registration endpoint
func getRegistration(db database.Database, w http.ResponseWriter, r *http.Request) {

	// Extract the username from the request and return if it returns empty
	username := utils.ExtractUsername(w, r)
	if username == "" {
		return
	}
	w.Header().Add("Content-Type", "application/json")

	statusCode, response, err := db.ReadUser(username)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
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
func postRegistration(db database.Database, w http.ResponseWriter, r *http.Request) {

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

	// Create user in the database
	statusCode, err := db.CreateUser(response)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	// Set response header to JSON and return status code 201
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error returning output", http.StatusInternalServerError)
		log.Println("Error returning output")
		return
	}
}

// putRegistration is a function to handle PUT requests to the registration endpoint
func putRegistration(db database.Database, w http.ResponseWriter, r *http.Request) {

	// Extract the username from the request and return if it returns empty
	username := utils.ExtractUsername(w, r)
	if username == "" {
		return
	}

	log.Println("Username: ", username)

	// Instantiate a new decoder and a new response struct
	decoder := json.NewDecoder(r.Body)
	putRequest := utils.UserRegistration{}
	err := decoder.Decode(&putRequest)
	if err != nil {
		http.Error(w, "Error decoding PUT request", http.StatusBadRequest)
		log.Println("Error decoding PUT request")
		return
	}

	// Update the user in the database
	statusCode, err := db.UpdateUser(username, putRequest)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.WriteHeader(statusCode)
}

// deleteRegistration is a function to handle DELETE requests to the registration endpoint
func deleteRegistration(db database.Database, w http.ResponseWriter, r *http.Request) {

	// Extract the username from the request and return if it returns empty
	username := utils.ExtractUsername(w, r)
	if username == "" {
		return
	}

	statusCode, err := db.DeleteUser(username)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
	w.WriteHeader(statusCode)

	//// Check if user exists
	//found, _ := db.CheckUserExistence(username)
	//if !found {
	//	http.Error(w, "User not found", http.StatusNotFound)
	//	log.Println("User not found")
	//	return
	//}
	//
	//// Open the collection and delete the user from the database
	//collection := utils.Client.Database(utils.COLLECTION_USERS).Collection(utils.COLLECTION_USERS)
	//_, err := collection.DeleteOne(context.TODO(), bson.M{"username": username})
	//if err != nil {
	//	http.Error(w, "Error deleting user", http.StatusInternalServerError)
	//	log.Println("Error deleting user")
	//	return
	//}
}
