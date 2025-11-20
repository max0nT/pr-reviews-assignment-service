package main

import (
	"log"

	"github.com/max0nT/pr-assign/config"
	"github.com/max0nT/pr-assign/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
