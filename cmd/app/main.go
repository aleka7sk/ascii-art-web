package main

import (
	"github.com/aleka7sk/ascii-art-web/config"
	"github.com/aleka7sk/ascii-art-web/server"
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
