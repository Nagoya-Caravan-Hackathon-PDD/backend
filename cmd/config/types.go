package config

var Config = &config{}

type config struct {
	Server    Server
	Cockroach Cockroach
}

type Server struct {
	Port string `env:"SERVER_ADDR" envDefault:"localhost:8080"`
}

type Cockroach struct {
	Host     string `env:"PSQL_HOST" envDefault:"localhost"`
	Port     int    `env:"PSQL_PORT" envDefault:"5432"`
	User     string `env:"PSQL_USER" envDefault:"postgres"`
	Password string `env:"PSQL_PASSWORD" envDefault:"postgres"`
	DBName   string `env:"PSQL_DBNAME" envDefault:"pdd-datastore"`
	SSLMode  string `env:"PSQL_SSLMODE" envDefault:"disable"`

	ConnectTimeout  int  `env:"PSQL_CONNECT_TIMEOUT" envDefault:"10"`
	ConnectWaitTime int  `env:"PSQL_CONNECT_WAIT_TIME" envDefault:"10"`
	ConnectAttempts int  `env:"PSQL_CONNECT_ATTEMPTS" envDefault:"3"`
	ConnectBlocks   bool `env:"PSQL_CONNECT_BLOCKS" envDefault:"false"`
	CloseTimeout    int  `env:"PSQL_CLOSE_TIMEOUT" envDefault:"10"`
}
