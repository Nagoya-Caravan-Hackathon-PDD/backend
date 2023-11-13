package gateways

import (
	"database/sql"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
)

type userGateway struct {
	db *sql.DB
}

func NewUserGateway(db *sql.DB) *userGateway {
	return &userGateway{db}
}

func (g *userGateway) Create(arg types.CreateUser) (err error) {
	const query = `INSERT INTO users (user_id,github_id) VALUES ($1, $2)`
	row := g.db.QueryRow(query, arg.UserID, arg.GitHubID)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

func (g *userGateway) Read(userID string) (err error) {
	const query = `SELECT * FROM users WHERE user_id = $1`
	row := g.db.QueryRow(query, userID)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

func (g *userGateway) Delete(userID string) (err error) {
	const query = `DELETE FROM users WHERE user_id = $1`
	row := g.db.QueryRow(query, userID)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}
