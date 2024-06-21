// Test File: registrationHandler_test.go
// Contains HTTP tests for the registration handler functions.
//-----------------------------------------------------------------------
// Test Function: TestPostRegistration
// Test Function: TestGetRegistration
// Test Function: TestPutRegistration
// Test Function: TestDeleteRegistration

package endpoints

import (
	"dashboard/database"
	"dashboard/utils"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Declare consts for usage in tests
const (
	USERNAME = "testuser"
	PASSWORD = "123456789"
	EMAIL    = "testuser@example.com"

	USERNAME2  = "testuser2"
	INVALID_PW = "password"
	COUNTRY    = "France"
	ISOCODE    = "FR"
	CITY       = "Paris"
)

// TestPostRegistration tests the postRegistration function
func TestPostRegistration(t *testing.T) {

	db := database.NewMockDB()
	client := http.Client{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RegistrationHandler(db, w, r)
	}))
	defer server.Close()

	// Test case with user containing all required fields
	user := utils.UserRegistration{
		Username: USERNAME,
		Password: PASSWORD,
		Email:    EMAIL,
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

	req, err := http.NewRequest(http.MethodPost, server.URL+utils.PATH_REGISTRATIONS, strings.NewReader(string(jsonUser)))
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

	log.Println("------- TestPostRegistration passed -------")
}

// TestGetRegistration tests the getRegistration function
func TestGetRegistration(t *testing.T) {

	db := database.NewMockDB()
	client := http.Client{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RegistrationHandler(db, w, r)
	}))
	defer server.Close()

	// Test case with user containing all required fields
	user := utils.UserRegistration{
		Username: USERNAME,
		Password: PASSWORD,
		Email:    EMAIL,
		Preference: utils.UserPreferences{
			Football: true,
			Movies:   false,
			Weather:  true,
		},
	}

	db.CreateUser(user)

	resp, err := client.Get(server.URL + utils.PATH_REGISTRATIONS + "testuser/")
	if err != nil {
		t.Fatal("Get request to URL failed:", err.Error())
	}

	// Check the status code
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("expected %d but got %d", http.StatusOK, resp.StatusCode)
	} else if resp == nil {
		t.Error("response is nil")
	}

	// Check the response body
	var response utils.UserRegistration
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatal("Failed to decode response body:", err.Error())
	}

	// Test status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected %d but got %d", http.StatusOK, resp.StatusCode)
	}

	// Test certain values
	if response.Username != user.Username {
		t.Errorf("expected %s, but got %s", user.Username, response.Username)
	}

	if response.Email != user.Email {
		t.Errorf("expected %s, but got %s", user.Email, response.Email)
	}

	if response.Preference.Football != user.Preference.Football {
		t.Errorf("expected %v, but got %v", user.Preference.Football, response.Preference.Football)
	}

	if response.Preference.Movies != user.Preference.Movies {
		t.Errorf("expected %v, but got %v", user.Preference.Movies, response.Preference.Movies)
	}

	if response.Preference.Weather != user.Preference.Weather {
		t.Errorf("expected %v, but got %v", user.Preference.Weather, response.Preference.Weather)
	}

	log.Println("------- TestGetRegistration passed -------")
}

// TestPutRegistration tests the putRegistration function
func TestPutRegistration(t *testing.T) {

	db := database.NewMockDB()
	client := http.Client{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RegistrationHandler(db, w, r)
	}))
	defer server.Close()

	// Test case with user containing all required fields
	user := utils.UserRegistration{
		Username: USERNAME,
		Password: PASSWORD,
		Email:    EMAIL,
		Preference: utils.UserPreferences{
			Football: false,
			Movies:   false,
			Weather:  false,
		},
	}
	db.CreateUser(user)

	// Test case with updated value
	userPut := utils.UserRegistration{
		Username: USERNAME,
		Password: PASSWORD,
		Email:    EMAIL,
		Preference: utils.UserPreferences{
			Football: true,
			Movies:   true,
			Weather:  true,
		},
	}

	// JSONify the userPut struct
	userPutJSON, err := json.Marshal(userPut)
	if err != nil {
		t.Fatal("Failed to marshal userPut struct")
	}

	req, err := http.NewRequest(http.MethodPut, server.URL+utils.PATH_REGISTRATIONS+"testuser/", strings.NewReader(string(userPutJSON)))
	if err != nil {
		t.Fatal("Failed to create request")
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Failed to send request")
	}

	// Check the status code
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("expected %d but got %d", http.StatusOK, resp.StatusCode)
	}

	_, response, _ := db.ReadUser("testuser")

	// Test certain values
	if response.Username != userPut.Username {
		t.Errorf("expected %s, but got %s", userPut.Username, response.Username)
	}

	if response.Email != userPut.Email {
		t.Errorf("expected %s, but got %s", userPut.Email, response.Email)
	}

	if response.Preference.Football != userPut.Preference.Football {
		t.Errorf("expected %v, but got %v", userPut.Preference.Football, response.Preference.Football)
	}

	// Test case with changing username
	userPut.Username = USERNAME2
	userPutJSON, err = json.Marshal(userPut)
	if err != nil {
		t.Fatal("Failed to marshal userPut struct")
	}
	req, err = http.NewRequest(http.MethodPut, server.URL+utils.PATH_REGISTRATIONS+"testuser/", strings.NewReader(string(userPutJSON)))
	if err != nil {
		t.Fatal("Failed to create request")
	}
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal("Failed to send request")
	}

	// Check status code
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected %d but got %d", http.StatusBadRequest, resp.StatusCode)
	}

	log.Println("------- TestPutRegistration passed -------")
}

// TestDeleteRegistration tests the deleteRegistration function
func TestDeleteRegistration(t *testing.T) {

	db := database.NewMockDB()
	client := http.Client{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RegistrationHandler(db, w, r)
	}))
	defer server.Close()

	// Test case with user containing all required fields
	user := utils.UserRegistration{
		Username: USERNAME,
		Password: PASSWORD,
		Email:    EMAIL,
		Preference: utils.UserPreferences{
			Football: true,
			Movies:   true,
			Weather:  true,
		},
	}

	db.CreateUser(user)

	req, err := http.NewRequest(http.MethodDelete, server.URL+utils.PATH_REGISTRATIONS+"testuser/", nil)
	if err != nil {
		t.Fatal("Failed to create request")
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Failed to send request")
	}

	// Check the status code
	if resp != nil && resp.StatusCode != http.StatusNoContent {
		t.Errorf("expected %d but got %d", http.StatusNoContent, resp.StatusCode)
	} else if resp == nil {
		t.Error("response is nil")
	}

	// Check if the user is deleted
	statusCode, _, err := db.ReadUser("testuser")
	if err == nil {
		t.Error("User was not deleted")
	}

	if statusCode != http.StatusNotFound {
		t.Errorf("expected %d but got %d", http.StatusNotFound, statusCode)
	}

	// Attempt to delete user that is not in the database
	req, err = http.NewRequest(http.MethodDelete, server.URL+utils.PATH_REGISTRATIONS+"testuser/", nil)
	if err != nil {
		t.Fatal("Failed to create request")
	}
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal("Failed to send request")
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("expected %d but got %d", http.StatusNotFound, resp.StatusCode)
	}

	log.Println("------- TestDeleteRegistration passed -------")
}
