package endpoints

import (
	"dashboard/utils"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

// SetCookie is a function to set a cookie for a user
func SetCookie(w http.ResponseWriter, r *http.Request) {

	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		log.Println("Invalid request method")
		return
	}

	// Read the request body
	user := utils.UserRegistration{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error decoding POST request", http.StatusBadRequest)
		log.Println("Error decoding Cookie POST request")
		return
	}

	// Generate a random token
	token, err := utils.GenerateRandomToken(32)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		log.Println("Error generating token")
		return
	}

	log.Println("Token: " + token)

	cookie := &http.Cookie{
		Name:     user.Username,
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		Domain:   "localhost",
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	log.Println("Cookie for user '" + cookie.Name + "' set")

}

// GetCookie is a function to get a cookie for a user
func GetCookie(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		log.Println("Invalid request method")
		return
	}

	// Extract the username from the request path
	utils.EnsureCorrectPath(r)
	username := utils.ExtractUsername(w, r)

	cookie, err := r.Cookie(username)
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			log.Println("No cookie found for user '" + username + "'")

		default:
			log.Println("Error getting cookie for user '" + username + "'.")
		}
		return
	}

	log.Println("Cookie for user '" + cookie.Name + "' found")

}
