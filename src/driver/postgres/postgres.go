package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/cmd/config"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/pkg/utils"
	_ "github.com/lib/pq"
)

type Connection interface {
	Connection() (*sql.DB, error)
	Close(ctx context.Context) error
}

type pqDatabaseConnection struct {
	dsn string

	connectAttempts        int
	connectWaitTimeSeconds int
	connectBlocks          bool
	connecting             bool
	connectError           error

	status          Status
	conn            *sql.DB
	connectFinished chan bool
}

func NewConnection() Connection {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		config.Config.Cockroach.User,
		config.Config.Cockroach.Password,
		config.Config.Cockroach.Host,
		config.Config.Cockroach.Port,
		config.Config.Cockroach.DBName,
		config.Config.Cockroach.SSLMode,
	)

	connectAttempts := config.Config.Cockroach.ConnectAttempts
	if connectAttempts == 0 {
		connectAttempts = 1
	}

	connectWaitTimeSeconds := config.Config.Cockroach.ConnectWaitTime
	if connectWaitTimeSeconds == 0 {
		connectWaitTimeSeconds = 3
	}

	conn := &pqDatabaseConnection{
		dsn:                    dsn,
		connectAttempts:        connectAttempts,
		connectBlocks:          config.Config.Cockroach.ConnectBlocks,
		connectWaitTimeSeconds: connectWaitTimeSeconds,
		status:                 UNKNOWN,
	}

	conn.updateDBStatus()

	return conn
}

func (c *pqDatabaseConnection) Connection() (*sql.DB, error) {
	c.updateDBStatus()

	if c.status != READY {
		c.tryConnect()
	}

	if c.status == READY {
		return c.conn, nil
	}

	return nil, fmt.Errorf("database connection is not ready: %v, %v", c.status, c.connectError)
}

// tryConnect attempts opening single connection to the database.
func (c *pqDatabaseConnection) tryConnect() {
	if !c.connecting {
		c.connect()
	} else if c.connectBlocks {
		// different goroutine is connecting, wait until finished
		<-c.connectFinished
	}
}

func (c *pqDatabaseConnection) connect() {
	c.connecting = true

	pqConnect := func() (*sql.DB, error) {
		return sql.Open("postgres", c.dsn)
	}
	sleep := func(seconds int) {
		time.Sleep(time.Duration(seconds) * time.Second)
	}

	if c.connectAttempts < 0 {
		for c.status != READY {
			c.conn, c.connectError = pqConnect()
			c.updateDBStatus()

			if c.status != READY {
				log.Printf("unable to connect to database: %v. retrying after %d seconds", c.connectError, c.connectWaitTimeSeconds)
				sleep(c.connectWaitTimeSeconds)
			}
		}

		log.Printf("connected with pq to postgres")
	} else {
		var err error
		for i := 0; i < c.connectAttempts; i++ {
			c.conn, err = pqConnect()
			c.updateDBStatus()

			if c.status != READY {
				log.Printf("unable to connect to database: %v", err)

				if i < c.connectAttempts-1 {
					sleep(c.connectWaitTimeSeconds)
				}
			} else {
				log.Printf("connected with pq to postgres")
				break
			}
		}

		if c.isConnNil() {
			log.Printf("failed to connect to database in %d tries: %v", c.connectAttempts, err)
			c.connectError = err
		}
	}

	c.connecting = false
	go func() {
		c.connectFinished <- true
	}()
}

func (c *pqDatabaseConnection) isConnNil() bool {
	return utils.IsInterfaceNil(c.conn)
}

func (c *pqDatabaseConnection) updateDBStatus() {
	if c.isConnNil() {
		c.status = NOT_READY
		return
	}

	if err := c.conn.Ping(); err != nil {
		log.Printf("failed to ping database: %v", err)
		c.status = ERROR

		if err = c.Close(context.Background()); err == nil {
			c.status = NOT_READY
		}
	} else {
		c.status = READY
	}
}

// Close closes the connection to the database.
func (c *pqDatabaseConnection) Close(ctx context.Context) error {
	log.Printf("closing pq postgres db connection")

	if c.isConnNil() {
		log.Printf("no connection to close")
		return nil
	}

	err := c.conn.Close()
	if err != nil {
		log.Printf("failed to close connection: %v", err)
	}

	c.status = DISCONNECTED
	c.conn = nil
	return err
}
