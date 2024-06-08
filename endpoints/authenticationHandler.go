package endpoints

import (
	"dashboard/database"
	"dashboard/utils"
	"encoding/json"
	"net/http"
)

func AuthenticationHandler(db database.Database, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch r.Method {
	case http.MethodPost:
		postAuthentication(db, w, r)
	default:
		http.Error(w, "Unsupported request method '"+r.Method+"'. Only "+
			http.MethodPost+" is supported.", http.StatusNotImplemented)
		return
	}

}

func postAuthentication(db database.Database, w http.ResponseWriter, r *http.Request) {

	// Instantiate a new decoder and a new response struct
	decoder := json.NewDecoder(r.Body)
	response := utils.UserAuthentication{}

	err := decoder.Decode(&response)
	if err != nil {
		http.Error(w, "Error decoding request", http.StatusBadRequest)
		return
	}

	statusCode, err := db.AuthenticateUser(db, response)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.WriteHeader(http.StatusOK)

}
