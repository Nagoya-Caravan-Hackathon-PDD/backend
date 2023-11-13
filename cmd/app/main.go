package main

import (
	"context"
	"flag"
	"log"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/cmd/config"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/driver/postgres"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/infrastructure/server"
)

func main() {
	var (
		usedotEnv = flag.Bool("usedotenv", false, "use .env file")
		path      = flag.String("path", ".env", "path to .env file")
	)
	flag.Parse()

	if *usedotEnv {
		config.LoadEnv(*path)
	} else {
		config.LoadEnv()
	}
	log.Println(config.Config)
	dbconn := postgres.NewConnection()
	defer dbconn.Close(context.Background())

	db, err := dbconn.Connection()
	if err != nil {
		log.Fatalf("failed to connect to database :%v", err)
	}

	server.NewHTTPserver(db).Run()
}
