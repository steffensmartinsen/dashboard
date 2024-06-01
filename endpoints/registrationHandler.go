package endpoints

import "net/http"

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		getRegistration(w, r)
	case http.MethodPost:
		postRegistration(w, r)
	case http.MethodPut:
		putRegistration(w, r)
	case http.MethodDelete:
		deleteRegistration(w, r)
	}
}

func getRegistration(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func postRegistration(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func putRegistration(w http.ResponseWriter, r *http.Request)  {
	// TODO
}

func deleteRegistration(w http.ResponseWriter, r *http.Request) {

}

