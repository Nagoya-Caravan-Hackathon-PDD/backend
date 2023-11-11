package config

var Config = &config{}

type config struct {
	Server struct {
		Port string `envconfig:"SERVER_PORT" default:"8080"`
	}
}
