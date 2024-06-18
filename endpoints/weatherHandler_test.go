package endpoints

import (
	"dashboard/database"
	"dashboard/utils"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWeatherHandler(t *testing.T) {

	db := database.NewMockDB()
	client := http.Client{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		WeatherHandler(db, w, r)
	}))
	defer server.Close()

	// Test case with user containing all required fields
	user := utils.UserRegistration{
		Username: "testuser",
		Password: "1234567890",
		Email:    "testuser@example.com",
		Country:  utils.Country{"France", "FR"},
		City:     "Paris",
	}
	db.CreateUser(user)

	resp, err := client.Get(server.URL + utils.PATH_WEATHER + user.Username + "/")
	if err != nil {
		t.Fatal("Get request to URL failed:", err.Error())
	}

	// Check the status code
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("expected %d but got %d", http.StatusOK, resp.StatusCode)
	}

	// Check the response body
	var response utils.WeatherData
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatal("Failed to decode response body:", err.Error())
	}

	if response.Hourly.Temperature[0] != 15.7 {
		t.Errorf("expected temperature %f but got %f", 15.7, response.Hourly.Temperature[0])
	}
	if response.Hourly.Precipitation[0] != 0.00 {
		t.Errorf("expected precipitation %f but got %f", 0.00, response.Hourly.Precipitation[0])
	}
	if response.Hourly.CloudCover[0] != 100 {
		t.Errorf("expected cloud cover %f but got %f", 100.0, response.Hourly.CloudCover[0])
	}
	if response.Hourly.WindSpeed[0] != 7.8 {
		t.Errorf("expected wind speed %f but got %f", 7.8, response.Hourly.WindSpeed[0])
	}

	log.Println("------- TestWeatherHandler passed -------")
}
