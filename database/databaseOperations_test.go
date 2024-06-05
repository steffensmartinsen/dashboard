// File: databaseOperations_test.go
// Contains unit tests for the database operations
//-----------------------------------------------------------------------
// Test Function: TestCreateUser
// Test Function: TestReadUser

package database

import (
	"dashboard/utils"
	"net/http"
	"testing"
)

func TestCreateUser(t *testing.T) {

	db := NewMockDB()

	// Test case with user containing all required fields
	user := utils.UserRegistration{
		Username: "testuser",
		Password: "123456789",
		Email:    "testuser@example.com",
		Preference: utils.UserPreferences{
			Football: true,
			Movies:   false,
			Weather:  true,
		},
	}

	// Test case with user missing required fields
	user = utils.UserRegistration{
		Username: "testuser2",
		Password: "",
		Email:    "testuser2@example.com",
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
		Username: "testuser3",
		Password: "password",
		Email:    "testuser3@example.com",
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
		Username: "testuser",
		Password: "123456789",
		Email:    "testuser4@example.com",
	}

}

func TestReadUser(t *testing.T) {
	db := NewMockDB()

	// Test case with user containing all required fields
	user := utils.UserRegistration{
		Username: "testuser",
		Password: "123456789",
		Email:    "testuser@example.com",
		Preference: utils.UserPreferences{
			Football: true,
			Movies:   false,
			Weather:  true,
		},
	}

	// Create dummy case
	db.CreateUser(user)

	// Test case with user that exists
	status, response, err := db.ReadUser("testuser")

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
}
