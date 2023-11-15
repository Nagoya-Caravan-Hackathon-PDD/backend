package types

import "time"

type CreateGame struct {
	GameID    string    `json:"game_id"`
	OwnerID   string    `json:"owner_id"`
	EnemyID   string    `json:"enemy_id"`
	CreatedAt time.Time `json:"created_at"`
}
