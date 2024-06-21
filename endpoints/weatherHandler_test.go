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
		Username: USERNAME,
		Password: PASSWORD,
		Email:    EMAIL,
		Country:  utils.Country{COUNTRY, ISOCODE},
		City:     CITY,
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
	var response utils.WeeklyWeather
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatal("Failed to decode response body:", err.Error())
	}

	if response.Weather[0].Hours[0].Temperature != 15.7 {
		t.Errorf("expected temperature %f but got %f", 15.7, response.Weather[0].Hours[0].Temperature)
	}
	if response.Weather[0].Hours[0].Precipitation != 0.00 {
		t.Errorf("expected precipitation %f but got %f", 0.00, response.Weather[0].Hours[0].Precipitation)
	}
	if response.Weather[0].Hours[0].CloudCover != 100 {
		t.Errorf("expected cloud cover %f but got %f", 100.0, response.Weather[0].Hours[0].CloudCover)
	}
	if response.Weather[0].Hours[0].WindSpeed != 7.8 {
		t.Errorf("expected wind speed %f but got %f", 7.8, response.Weather[0].Hours[0].WindSpeed)
	}
	if response.Weather[0].Hours[8].Condition != utils.CONDITION_PARTLY_CLOUDY {
		t.Errorf("expected condition %s but got %s", utils.CONDITION_PARTLY_CLOUDY, response.Weather[0].Hours[8].Condition)
	}

	log.Println("------- TestWeatherHandler passed -------")
}
