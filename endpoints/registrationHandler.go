package endpoints

import (
	"context"
	"dashboard/database"
	"dashboard/utils"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
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

func putRegistration(db database.Database, w http.ResponseWriter, r *http.Request) {

	// Extract the username from the request and return if it returns empty
	username := utils.ExtractUsername(w, r)
	if username == "" {
		return
	}

	//// Check if user exists
	//found, _ := db.CheckUserExistence(username)
	//if !found {
	//	http.Error(w, "User not found", http.StatusNotFound)
	//	log.Println("User not found")
	//	return
	//}
	//
	// Instantiate a new decoder and a new response struct
	decoder := json.NewDecoder(r.Body)
	putRequest := utils.UserRegistration{}
	err := decoder.Decode(&putRequest)
	if err != nil {
		http.Error(w, "Error decoding PUT request", http.StatusBadRequest)
		log.Println("Error decoding PUT request")
		return
	}
	//
	//// Enforce username and email to be lowercase
	//putRequest.Username = strings.ToLower(putRequest.Username)
	//putRequest.Email = strings.ToLower(putRequest.Email)
	//
	//// Disallow any attempted change of username
	//if putRequest.Username != username {
	//	http.Error(w, "Username cannot be changed", http.StatusBadRequest)
	//	log.Println("Attempted change of username")
	//	return
	//}
	//
	//// Fetch the user from the database
	//collection := utils.Client.Database(utils.COLLECTION_USERS).Collection(utils.COLLECTION_USERS)
	//user := utils.UserRegistration{}
	//err = collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	//if err != nil {
	//	http.Error(w, "Error fetching user", http.StatusInternalServerError)
	//	log.Println("Error fetching user in PUT request")
	//	return
	//}
	//
	//// Check if the password is changed
	//if putRequest.Password != "" || putRequest.Password != user.Password {
	//
	//	// Apply constraints and hash the password if it is changed
	//	if !utils.EnforcePassword(putRequest.Password) {
	//		http.Error(w, "Please don't use an actual password for this. The only accepted characters are '1234567890'", http.StatusBadRequest)
	//		log.Println("Password not allowed")
	//		return
	//	}
	//	if len(putRequest.Password) < 8 {
	//		http.Error(w, "Password must be at least 8 characters long", http.StatusBadRequest)
	//		log.Println("Password too short")
	//		return
	//	}
	//	putRequest.Password, err = utils.HashPassword(putRequest.Password)
	//	if err != nil {
	//		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	//		log.Println("Error hashing password")
	//		return
	//	}
	//}
	//
	//// Check if email is changed
	//if putRequest.Email != "" && putRequest.Email != user.Email {
	//	putRequest.Email = strings.ToLower(putRequest.Email)
	//}
	//
	//log.Println(putRequest)
	//
	//// Update the user in the database
	//_, err = collection.UpdateOne(context.TODO(), bson.M{"username": username}, bson.M{"$set": putRequest})
	//if err != nil {
	//	http.Error(w, "Something went wrong", http.StatusInternalServerError)
	//	log.Println("Error updating user")
	//	return
	//}

	statusCode, err := db.UpdateUser(putRequest)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.WriteHeader(statusCode)

}

func deleteRegistration(db database.Database, w http.ResponseWriter, r *http.Request) {

	// Extract the username from the request and return if it returns empty
	username := utils.ExtractUsername(w, r)
	if username == "" {
		return
	}

	// Check if user exists
	found, _ := db.CheckUserExistence(username)
	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		log.Println("User not found")
		return
	}

	// Open the collection and delete the user from the database
	collection := utils.Client.Database(utils.COLLECTION_USERS).Collection(utils.COLLECTION_USERS)
	_, err := collection.DeleteOne(context.TODO(), bson.M{"username": username})
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		log.Println("Error deleting user")
		return
	}
	w.WriteHeader(http.StatusNoContent)
	log.Println("User '" + username + "' deleted.")
}
