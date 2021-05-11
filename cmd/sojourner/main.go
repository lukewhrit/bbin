package main

import (
	"log"

	"github.com/lukewhrit/sojourner/internal/config"
	"github.com/lukewhrit/sojourner/internal/database"
	"github.com/lukewhrit/sojourner/internal/server"
)

func init() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	if err := database.Open(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if err := server.Start(
		config.Config.HostName,
		config.Config.Port,
	); err != nil {
		log.Fatal(err)
	}

	defer database.Close()
}
