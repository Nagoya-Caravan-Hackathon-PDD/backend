package paseto

import "time"

type Maker interface {
	// トークンを作る
	CreateToken(game_id, owner_id string, duration time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}
