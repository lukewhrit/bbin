package routes

import "net/http"

// RetrieveHandler is the HTTP handler for the create document endpoint
func RetrieveHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
