package config

var Config = &config{}

type config struct {
	Server struct {
		Port string `env:"SERVER_ADDR" envDefault:"localhost:8080"`
	}
}
