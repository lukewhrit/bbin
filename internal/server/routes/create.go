package routes

import (
	"math"
	"net/http"

	"github.com/lukewhrit/sojourner/internal/config"
	"github.com/lukewhrit/sojourner/internal/database"
	"github.com/lukewhrit/sojourner/internal/utils"
)

// CreateHandler is the HTTP handler for the create document endpoint
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	// Parse multipart/form-data body value
	// The maximum body size is 268435456 B (256 MB)
	if err := r.ParseMultipartForm(
		// Multiply MaxPayloadSize by 1024² to get value in bytes
		int64(config.Config.MaxPayloadSize * math.Pow(1024, 2)),
	); err != nil {
		w.WriteHeader(http.StatusRequestEntityTooLarge)
	}

	// Make sure the body has the "content" field and that field is not empty
	if r.FormValue("content") != "" {
		// Generate ID and save the POSTed document
		id := utils.GenerateID()
		err := database.Set([]byte(id), []byte(r.FormValue("content")))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		// Send back ID with a 201 "Created" status code
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(id))
	}
}
