package config

import (
	"log"

	"github.com/caarlos0/env"
)

func init() {
	var Config = &config{}

	if err := env.Parse(Config); err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

}
