package gateways

import "database/sql"

type healthGateway struct {
	db *sql.DB
}

func NewHealthGateway(db *sql.DB) *healthGateway {
	return &healthGateway{db}
}

func (g *healthGateway) Ping() error {
	return g.db.Ping()
}
