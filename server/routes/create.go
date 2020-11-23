package routes

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/lukewhrit/bbin/database"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateID(length int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, length)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

// CreateHandler is the HTTP handler for the create document endpoint
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	id := generateID(8)
	err := database.Set([]byte(id), []byte("xx"))

	if err != nil {
		log.Fatal(err)
	}

	value, err := database.Get([]byte(id))

	if err != nil {
		log.Fatal(err)
	}

	w.Write(value)
}
