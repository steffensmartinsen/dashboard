package endpoints

import (
	"dashboard/database"
	"dashboard/utils"
	"net/http"
)

// DashboardHandler is the handler for the dashboard endpoint
func DashboardHandler(db database.Database, w http.ResponseWriter, r *http.Request) {

	utils.EnsureCorrectPath(r)

	switch r.Method {
	case http.MethodGet:
		getDashboard(db, w, r)
	default:
		http.Error(w, "REST Method '"+r.Method+" not supported. Currently only '"+
			http.MethodGet+"' is supported.", http.StatusMethodNotAllowed)
	}

}

// getDashboard is a function to handle GET requests to the dashboard endpoint
func getDashboard(db database.Database, w http.ResponseWriter, r *http.Request) {

	// TODO
	w.Write([]byte("DashboardHandler reached\n"))
}
