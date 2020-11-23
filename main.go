package main

import (
	"log"
	"os"

	"github.com/lukewhrit/bbin/database"
	"github.com/lukewhrit/bbin/server"
)

var (
	hostname = os.Getenv("BBIN_HOSTNAME")
	port     = os.Getenv("BBIN_PORT")
)

func init() {
	if err := database.Open(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if err := server.Start(hostname, port); err != nil {
		log.Fatal(err)
	}

	defer database.Close()
}
