package input

type CreateGameRequest struct {
	OwnerID string `json:"owner_id"`
}

type JoinGameRequest struct {
	GameID string `param:"game_id"`
	UserID string `json:"user_id"`
}

type ListGameRequest struct {
	UserID   string `query:"user_id"`
	PageID   uint   `query:"page_id"`
	PageSize uint   `query:"page_size"`
}

type GetAnyGameRequest struct {
	GameID string `param:"game_id"`
}
