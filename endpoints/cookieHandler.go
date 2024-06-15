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

	// Set the headers
	utils.SetHeaders(w, r)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

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
		Domain:   utils.LOCALHOST,
		Path:     utils.SLASH,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	log.Println("Cookie for user '" + cookie.Name + "' set")
}

// GetCookie is a function to get a cookie for a user
func GetCookie(w http.ResponseWriter, r *http.Request) {

	// Set the headers
	utils.SetHeaders(w, r)

	// Ensure the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		log.Println("Invalid request method")
		return
	}

	// Extract the username from the request path
	utils.EnsureCorrectPath(r)
	username := utils.ExtractUsername(w, r)

	// Get the cookie for the user
	cookie, err := r.Cookie(username)
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):

			log.Println("No cookie found for user '" + username + "'")
			w.WriteHeader(http.StatusNotFound)

		default:
			log.Println("Error getting cookie for user '" + username + "'.")
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	log.Println("Cookie for user '" + cookie.Name + "' found")
	w.Write([]byte(cookie.Value))
}

// DeleteCookie is a function to delete a cookie for a user
func DeleteCookie(w http.ResponseWriter, r *http.Request) {

	// Set the headers
	utils.SetHeaders(w, r)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Ensure the request method is DELETE
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		log.Println("Invalid request method")
		return
	}

	// Extract the username from the request path
	utils.EnsureCorrectPath(r)
	username := utils.ExtractUsername(w, r)

	// Get the cookie for the user
	cookie, err := r.Cookie(username)
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):

			log.Println("No cookie found for user '" + username + "'")
			w.WriteHeader(http.StatusNotFound)

		default:
			log.Println("Error getting cookie for user '" + username + "'.")
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// Ensure the same attributes are set for the cookie
	cookie.Expires = time.Unix(0, 0)
	cookie.Domain = utils.LOCALHOST
	cookie.Path = utils.SLASH
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteLaxMode
	cookie.Secure = true

	// Delete the cookie and return '204 No Content'
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusNoContent)
	log.Println("Cookie for user '" + cookie.Name + "' deleted")
}
