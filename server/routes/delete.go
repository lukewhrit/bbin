package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lukewhrit/bbin/database"
)

// DeleteHandler is the HTTP handler for the create document endpoint
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := database.Delete([]byte(id))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("deleted"))
}
