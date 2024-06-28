package endpoints

import (
	"dashboard/database"
	"dashboard/utils"
	"encoding/json"
	"log"
	"net/http"
)

// WeatherHandler is the handler for the dashboard endpoint
func WeatherHandler(db database.Database, w http.ResponseWriter, r *http.Request) {

	utils.EnsureCorrectPath(r)

	switch r.Method {
	case http.MethodGet:
		getWeather(db, w, r)
	default:
		http.Error(w, "REST Method '"+r.Method+" not supported. Currently only '"+
			http.MethodGet+"' is supported.", http.StatusMethodNotAllowed)
	}

}

// getWeather is a function to handle GET requests to the dashboard endpoint
func getWeather(db database.Database, w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	// Extract the username from the request and return if it returns empty
	username := utils.ExtractUsername(w, r)
	if username == "" {
		return
	}

	// Get the user from the database
	statusCode, response, err := db.ReadUser(username)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	// Get the weather for the user
	statusCode, weatherData, err := db.GetWeather(response.Country, response.City)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	// TODO Rework the return value to have the date in the DailyWeather struct and the hours within
	// the hours field.

	// Get the weather forecast for the next 7 days
	weeklyWeather, err := utils.SetWeeklyWeather(weatherData)

	// Encode the response struct to the client
	err = json.NewEncoder(w).Encode(weeklyWeather)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
	log.Println(weeklyWeather.Weather[0].Date)
	log.Println(weeklyWeather.Weather[0].Hours[14].Hour)
}
