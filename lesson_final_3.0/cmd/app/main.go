package main

import (
	"log"
	"main/config"
	"main/server"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.NewApp()
	app.Run()
}
