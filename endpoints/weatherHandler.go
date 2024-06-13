package endpoints

import (
	"dashboard/database"
	"dashboard/utils"
	"net/http"
)

// DashboardHandler is the handler for the dashboard endpoint
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

// getDashboard is a function to handle GET requests to the dashboard endpoint
func getWeather(db database.Database, w http.ResponseWriter, r *http.Request) {

}
