package types

import "database/sql"

type GameTemp struct {
	GameID    string
	OwnerID   string
	OwnerIsOK bool
	EnemyID   sql.NullString
	EnemyIsOK bool
	IsBlock   bool
}
