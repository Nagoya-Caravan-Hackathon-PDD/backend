package main

import (
	"context"
	"log"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/cmd/config"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/driver/postgres"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/infrastructure/server"
)

func main() {
	config.LoadEnv()

	dbconn := postgres.NewConnection()
	defer dbconn.Close(context.Background())

	db, err := dbconn.Connection()
	if err != nil {
		log.Fatalf("failed to connect to database :%v", err)
	}

	server.NewHTTPserver(db).Run()
}
