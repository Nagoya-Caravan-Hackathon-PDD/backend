package config

import (
	"log"

	"github.com/caarlos0/env"
)

func init() {
	config := &config{}

	if err := env.Parse(&config.Server); err != nil {
		log.Fatalf("env load error: %v", err)
	}
	Config = config
}
