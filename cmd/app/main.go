package main

import (
	"context"
	"flag"
	"log"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/cmd/config"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/driver/postgres"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/infrastructure/server"
)

// @title PDD-GitHub-Go-Backend API
// @version 0.1
// @description This is a PDD-GitHub-Go-Backend API server

// @contact.name murasame29
// @contact.email oogiriminister@gamil.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1
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

	dbconn := postgres.NewConnection()
	defer dbconn.Close(context.Background())

	db, err := dbconn.Connection()
	if err != nil {
		log.Fatalf("failed to connect to database :%v", err)
	}

	server.NewHTTPserver(db).Run()
}
