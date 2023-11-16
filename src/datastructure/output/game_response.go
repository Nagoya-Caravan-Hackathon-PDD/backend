package output

type CreateGameResponse struct {
	GameID          string `json:"game_id"`
	GameServerToken string `json:"game_server_token"`
}

type JoinGameResponse struct {
	GameID          string `json:"game_id"`
	GameServerToken string `json:"game_server_token"`
}

type GetAnyGameResponse struct {
	GameID  string `json:"game_id"`
	OwnerID string `json:"owner_id"`
	EnemyID string `json:"enemy_id"`
	Result  string `json:"result"`
}
