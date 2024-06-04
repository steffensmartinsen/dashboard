package endpoints

import (
	"dashboard/database"
	"dashboard/utils"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostRegistration(t *testing.T) {

	db := database.NewMockDB()
	client := http.Client{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RegistrationHandler(db, w, r)
	}))
	defer server.Close()

	// Test case with user containing all required fields
	user := utils.UserRegistration{
		Username: "testuser5",
		Password: "123456789",
		Email:    "testuser5@example.com",
		Preference: utils.UserPreferences{
			Football: true,
			Movies:   false,
			Weather:  true,
		},
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", server.URL+utils.PATH_REGISTRATIONS, strings.NewReader(string(jsonUser)))
	if err != nil {
		t.Fatal("POST request failed", err.Error())
	}

	resp, err := client.Do(req)

	if resp != nil && resp.StatusCode != http.StatusCreated {
		t.Errorf("expected %d, but got %d", http.StatusCreated, resp.StatusCode)
	} else if resp == nil {
		t.Error("response is nil")
	}

	// Check the response body
	var response utils.UserRegistration
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatal("Failed to decode response body", err.Error())
	}

	// Test certain values
	if response.Username != user.Username {
		t.Errorf("expected %s, but got %s", user.Username, response.Username)
	}

	if response.Password != user.Password {
		t.Errorf("expected %s, but got %s", user.Password, response.Password)
	}

	if response.Email != user.Email {
		t.Errorf("expected %s, but got %s", user.Email, response.Email)
	}

	if response.Preference.Football != user.Preference.Football {
		t.Errorf("expected %v, but got %v", user.Preference.Football, response.Preference.Football)
	}

}
