package gateways

import (
	"database/sql"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
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

func (g *userGateway) Read(userID string) (*types.User, error) {
	const query = `SELECT * FROM users WHERE user_id = $1`
	var result *types.User
	row := g.db.QueryRow(query, userID)

	if err := row.Scan(&result.UserID, &result.GitHubID); err != nil {
		return nil, err
	}

	return result, nil
}

func (g *userGateway) Delete(userID string) (err error) {
	const query = `DELETE FROM users WHERE user_id = $1`
	result, err := g.db.Exec(query, userID)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return &pq.Error{Code: pgerrcode.NoDataFound}
	}
	return err
}
