package main

import (
	"log"

	"quiz-aventure/internal/config"
	"quiz-aventure/internal/server"
)

func main() {
	log.Println("Loading configuration...")
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("could not load config: %v\n", err)
	}

	srv := server.NewServer(cfg)
	log.Println("Starting server on", cfg.Server.Address)
	if err := srv.Run(); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
