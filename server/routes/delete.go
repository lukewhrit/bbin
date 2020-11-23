package routes

import "net/http"

// DeleteHandler is the HTTP handler for the create document endpoint
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
