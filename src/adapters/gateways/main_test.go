package gateways

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/cmd/config"
	psqlDriver "github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/driver/postgres"
)

var dbconn *sql.DB

func migrateInstance() (mig *migrate.Migrate, err error) {
	driver, err := postgres.WithInstance(dbconn, &postgres.Config{})
	if err != nil {
		return
	}

	mig, err = migrate.NewWithDatabaseInstance(
		"file://../../../cmd/migration",
		"postgres", driver)
	if err != nil {
		return
	}
	return
}

func migrateUp(t *testing.T) {
	mig, err := migrateInstance()
	if err != nil {
		t.Fatal(err)
	}

	mig.Up()
}

func migrateDown(t *testing.T) {
	t.Cleanup(func() {
		mig, err := migrateInstance()
		if err != nil {
			os.Exit(1)
		}

		mig.Down()
	})
}

func TestMain(m *testing.M) {
	config.LoadEnv()
	var err error
	conn := psqlDriver.NewConnection()
	defer conn.Close(context.Background())

	dbconn, err = conn.Connection()
	if err != nil {
		os.Exit(1)
	}

	os.Exit(m.Run())
}
