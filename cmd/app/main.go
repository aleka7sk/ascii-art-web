package main

import (
	"ascii-art-web/config"
	"ascii-art-web/server"
	"log"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Config initialize error: %v", err)
	}
	app := server.NewApp()
	app.Run(config)
}
