package utils

import "net/http"

// ThrowError handles errors reported by HTTP routes
func ThrowError(w http.ResponseWriter, r *http.Request, err error, status int) {}
