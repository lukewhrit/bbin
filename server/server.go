package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lukewhrit/bbin/server/routes"
)

// Start begins the server on the configured port
func Start(hostname, port string) error {
	// Create chi router instance
	r := chi.NewRouter()

	// Register middleware
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("multipart/form-data", ""))

	// Register routes
	r.Post("/", routes.CreateHandler)
	r.Get("/{id}", routes.RetrieveHandler)
	r.Delete("/{id}", routes.DeleteHandler)

	// Start server
	fmt.Println(fmt.Sprintf("BBin server listening on %s:%s", hostname, port))

	return http.ListenAndServe(fmt.Sprintf("%s:%s", hostname, port), r)
}
