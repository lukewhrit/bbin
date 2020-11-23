package utils

import (
	"math/rand"

	"github.com/lukewhrit/bbin/config"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// GenerateID makes an ID `length` characters long
func GenerateID() string {
	length := config.Config.IDLength

	b := make([]rune, length)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

// IsID validates whether a string is a valid ID
func IsID(id string) bool {
	// Requirements:
	//  1. ID is 8 characters long
	//  2. ID contains only characters in `letters`

	if len(id) == 8 {
		return true
	}

	return false
}
