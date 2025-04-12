package main

import (
	"flag"
	"log"
	"real-time-ranking/internal/app"
	"real-time-ranking/internal/config"
	"real-time-ranking/internal/server"
)

var (
	_path = flag.String("config-file-path", "./config.yaml", "config file path")
)

func main() {
	cfg, err := config.Load(*_path)
	if err != nil {
		log.Fatal(err)
	}

	app.Init(cfg)

	srv := server.NewServer(cfg.App)
	log.Println("Starting the server")

	if err := srv.Start(); err != nil {
		log.Fatalln("Failed to gracefully shutdown:", err)
	}

	log.Println("Server was shutdown gracefully")
}
