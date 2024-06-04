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

func TestGetUser(t *testing.T) {

}
