package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lukewhrit/bbin/database"
)

// RetrieveHandler is the HTTP handler for the create document endpoint
func RetrieveHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	value, err := database.Get([]byte(id))

	if err != nil {
		if err.Error() == "Key not found" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(value)
}
