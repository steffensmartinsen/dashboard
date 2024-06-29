// File: databaseOperations_test.go
// Contains unit tests for the database operations
//-----------------------------------------------------------------------
// Test Function: TestCreateUser
// Test Function: TestReadUser
// Test Function: TestUpdateUser
// Test Function: TestDeleteUser

package database

import (
	"dashboard/utils"
	"log"
	"net/http"
	"testing"
)

// Declare consts for usage in tests
const (
	USERNAME = "testuser"
	PASSWORD = "123456789"
	EMAIL    = "testuser@example.com"
	COUNTRY  = "France"
	ISOCODE  = "FR"
	CITY     = "Paris"

	USERNAME2  = "testuser2"
	EMAIL2     = "testuser2@example.com"
	INVALID_PW = "password"

	LATITUDE  = 48.85341
	LONGITUDE = 2.3488
)

func TestCreateUser(t *testing.T) {

	db := NewMockDB()

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

	// Test case with user missing required fields
	user = utils.UserRegistration{
		Username: USERNAME2,
		Password: "",
		Email:    EMAIL2,
	}
	status, err := db.CreateUser(user)
	if err == nil {
		t.Errorf("Expected \"missing required fields\" error, got nil")
	}
	if status != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", status)
	}

	// Test case with user containing a password with invalid characters
	user = utils.UserRegistration{
		Username: USERNAME2,
		Password: INVALID_PW,
		Email:    EMAIL2,
	}
	status, err = db.CreateUser(user)
	if err == nil {
		t.Errorf("Expected \"invalid password\" error, got nil")
	}
	if status != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", status)
	}

	// Test case with user containing an existing username
	user = utils.UserRegistration{
		Username: USERNAME,
		Password: PASSWORD,
		Email:    EMAIL2,
	}

	log.Println("------- TestCreateUser passed -------")

}

func TestReadUser(t *testing.T) {
	db := NewMockDB()

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

	// Create dummy case
	_, err := db.CreateUser(user)
	if err != nil {
		t.Fatal("Unable to create user for testing")
	}

	// Test case with user that exists
	status, response, err := db.ReadUser(USERNAME)

	// Test status code
	if status != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", status)
	}

	// Test error value
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	// Test struct values
	if response.Username != user.Username {
		t.Errorf("Expected %s, got %s", user.Username, response.Username)
	}

	if response.Email != user.Email {
		t.Errorf("Expected %s, got %s", user.Email, response.Email)
	}

	if response.Preference.Football != user.Preference.Football {
		t.Errorf("Expected %v, got %v", user.Preference.Football, response.Preference.Football)
	}

	if response.Preference.Movies != user.Preference.Movies {
		t.Errorf("Expected %v, got %v", user.Preference.Movies, response.Preference.Movies)
	}

	// Test case with user that does not exist
	status, _, err = db.ReadUser("testuser2")

	// Test status code
	if status != http.StatusNotFound {
		t.Errorf("Expected status code 404, got %v", status)
	}

	// Test error value
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	log.Println("------- TestReadUser passed -------")
}

func TestUpdateUser(t *testing.T) {
	db := NewMockDB()

	// Create dummy struct to update
	user := utils.UserRegistration{
		Username: USERNAME,
		Password: PASSWORD,
		Email:    EMAIL2,
		Preference: utils.UserPreferences{
			Football: true,
			Movies:   true,
			Weather:  true,
		},
	}

	// Create dummy case
	_, err := db.CreateUser(user)
	if err != nil {
		t.Fatal("Unable to create user for testing")
	}

	// Test case with correct request
	userUpdate := utils.UserRegistration{
		Username: USERNAME,
		Password: PASSWORD,
		Email:    EMAIL,
		Preference: utils.UserPreferences{
			Football: false,
			Movies:   false,
			Weather:  false,
		},
	}
	statusCode, err := db.UpdateUser(user.Username, userUpdate)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}

	// Test case with user that does not exist
	statusCode, err = db.UpdateUser(USERNAME2, userUpdate)
	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	// Test case with attempted username change
	userUpdate = utils.UserRegistration{
		Username: USERNAME2,
	}
	statusCode, err = db.UpdateUser(user.Username, userUpdate)
	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	log.Println("------- TestUpdateUser passed -------")
}

func TestDeleteUser(t *testing.T) {
	db := NewMockDB()

	// Create dummy struct to delete
	user := utils.UserRegistration{
		Username: USERNAME,
		Password: PASSWORD,
		Email:    EMAIL,
	}

	db.CreateUser(user)

	// Test case with user that exists
	statusCode, err := db.DeleteUser(USERNAME)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if statusCode != http.StatusNoContent {
		t.Errorf("Expected status code 204, got %v", statusCode)
	}

	// Test case with user that does not exist
	statusCode, err = db.DeleteUser(USERNAME)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if statusCode != http.StatusNotFound {
		t.Errorf("Expected status code 404, got %v", statusCode)
	}

	log.Println("------- TestDeleteUser passed -------")
}

func TestGetGeoCode(t *testing.T) {
	db := NewMockDB()

	// Test case with user containing all required fields
	user := utils.UserRegistration{
		Username: USERNAME,
		Password: PASSWORD,
		Email:    EMAIL,
		Country:  utils.Country{COUNTRY, ISOCODE},
		City:     CITY,
	}

	db.CreateUser(user)

	statusCode, response, err := db.GetGeoCode(user.Country, user.City)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}
	if response.Latitude != LATITUDE {
		t.Errorf("Expected %f, got %f", LATITUDE, response.Latitude)
	}
	if response.Longitude != LONGITUDE {
		t.Errorf("Expected %f, got %f", LONGITUDE, response.Longitude)
	}

	log.Println("------- TestGetGeoCode passed -------")

}
