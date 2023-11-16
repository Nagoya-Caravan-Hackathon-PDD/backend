package input

type ReadyGameRequest struct {
	GameID string `param:"game_id"`
	UserID string `json:"user_id"`
}

type ActionGameRequest struct {
	GameID    string `param:"game_id"`
	UserID    string `json:"user_id"`
	CommandID int    `json:"command_id"`
}

type FinTurnRequest struct {
	GameID string `param:"game_id"`
	UserID string `json:"user_id"`
}
