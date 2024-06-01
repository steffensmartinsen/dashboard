package endpoints

import (
	"net/http"
)

func EmptyHandler(w http.ResponseWriter, r *http.Request) {

	// TODO

	http.ServeFile(w, r, "utils/images/under_construction.jpg")
}
