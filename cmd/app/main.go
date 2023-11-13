package main

import (
	"context"
	"flag"
	"log"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/cmd/config"
	psqlDriver "github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/driver/postgres"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/infrastructure/server"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
	dbconn := psqlDriver.NewConnection()
	defer dbconn.Close(context.Background())

	db, err := dbconn.Connection()
	if err != nil {
		log.Fatalf("failed to connect to database :%v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed to Create Instance:%v", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migration",
		"postgres", driver)
	if err != nil {
		log.Fatalf("failed to Create Instance:%v", err)
	}
	if err := m.Up(); err != nil {
		log.Fatalf("failed to migrate up:%v", err)
	}

	server.NewHTTPserver(db).Run()
}
