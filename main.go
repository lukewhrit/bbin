package main

import (
	"fmt"
	"log"

	"github.com/lukewhrit/bbin/config"
	"github.com/lukewhrit/bbin/database"
	"github.com/lukewhrit/bbin/server"
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
	fmt.Println(config.Config)

	if err := server.Start(
		config.Config.HostName,
		config.Config.Port,
	); err != nil {
		log.Fatal(err)
	}

	defer database.Close()
}
