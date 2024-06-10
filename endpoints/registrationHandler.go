package endpoints

import (
	"dashboard/database"
	"dashboard/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func RegistrationHandler(db database.Database, w http.ResponseWriter, r *http.Request) {

	utils.EnsureCorrectPath(r)
	// Set the CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	switch r.Method {
	case http.MethodGet:
		getRegistration(db, w, r)
	case http.MethodPost:
		postRegistration(db, w, r)
	case http.MethodPut:
		putRegistration(db, w, r)
	case http.MethodDelete:
		deleteRegistration(db, w, r)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Unsupported request method '"+r.Method+"'. Only "+
			http.MethodGet+", "+http.MethodPost+", "+http.MethodPut+", and "+http.MethodDelete+" are supported.", http.StatusNotImplemented)
		return
	}
}

// getRegistration is a function to handle GET requests to the registration endpoint
func getRegistration(db database.Database, w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

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

	////Instantiate a new decoder and a new response struct
	//decoder := json.NewDecoder(r.Body)
	//registration := utils.UserRegistration{}
	//
	//// Decode the request into the response struct
	//err := decoder.Decode(&registration)
	//if err != nil {
	//	http.Error(w, "Error decoding POST request", http.StatusBadRequest)
	//	log.Println("Error decoding Registration POST request")
	//	return
	//}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading POST request", http.StatusBadRequest)
		log.Println("Error reading Registration POST request")
		return
	}

	var registration utils.UserRegistration
	err = json.Unmarshal(bodyBytes, &registration)
	if err != nil {
		http.Error(w, "Error decoding POST request", http.StatusBadRequest)
		log.Println("Error decoding Registration POST request (p1)")
		return
	}

	log.Println([]byte(registration.Preference))

	var preferences utils.UserPreferences
	err = json.Unmarshal([]byte(registration.Preference), &preferences)
	if err != nil {
		http.Error(w, "Error decoding POST request", http.StatusBadRequest)
		log.Println("Error decoding Registration POST request (p2)")
		return
	}

	// Create user in the database
	statusCode, err := db.CreateUser(registration)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	// Set response header to JSON and return status code 201
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	err = json.NewEncoder(w).Encode(registration)
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

}
