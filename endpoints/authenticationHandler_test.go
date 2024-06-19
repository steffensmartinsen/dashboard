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

func TestPostAuthentication(t *testing.T) {

	db := database.NewMockDB()
	client := http.Client{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AuthenticationHandler(db, w, r)
	}))
	defer server.Close()

	// Test case with user containing all required fields
	user := utils.UserRegistration{
		Username: USERNAME,
		Password: PASSWORD,
		Email:    EMAIL,
	}
	db.CreateUser(user)

	// Test case with user containing an invalid password
	request := utils.UserAuthentication{
		Username: USERNAME,
		Password: INVALID_PW,
	}

	jsonRequest, err := json.Marshal(request)
	if err != nil {
		t.Fatal("Failed to marshal request")
	}

	req, err := http.NewRequest(http.MethodPost, server.URL+utils.PATH_AUTHENTICATION,
		strings.NewReader(string(jsonRequest)))
	if err != nil {
		t.Fatal("Failed to create request")
	}

	resp, err := client.Do(req)

	if resp != nil && resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code 400, got %v", resp.StatusCode)
	} else if resp == nil {
		t.Errorf("Response is nil")
	}

	// Test case containing missing password
	request = utils.UserAuthentication{
		Username: USERNAME,
		Password: "",
	}

	jsonRequest, err = json.Marshal(request)
	req, err = http.NewRequest(http.MethodPost, server.URL+utils.PATH_AUTHENTICATION,
		strings.NewReader(string(jsonRequest)))
	if err != nil {
		t.Fatal("Failed to create request")
	}

	resp, err = client.Do(req)
	if resp != nil && resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", resp.StatusCode)
	} else if resp == nil {
		t.Errorf("Response is nil")
	}

	// Test case with user containing a valid request
	request = utils.UserAuthentication{
		Username: USERNAME,
		Password: PASSWORD,
	}

	jsonRequest, err = json.Marshal(request)
	req, err = http.NewRequest(http.MethodPost, server.URL+utils.PATH_AUTHENTICATION,
		strings.NewReader(string(jsonRequest)))
	if err != nil {
		t.Fatal("Failed to create request")
	}

	resp, err = client.Do(req)
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", resp.StatusCode)
	} else if resp == nil {
		t.Errorf("Response is nil")
	}

	log.Println("------- TestPostAuthentication passed -------")

	// TODO Create a Test case for empty username
}
