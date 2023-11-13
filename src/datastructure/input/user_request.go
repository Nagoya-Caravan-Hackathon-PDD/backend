package input

type CreateUser struct {
	GitHub string `json:"github_id"`
}

type GetUser struct {
	UserID string `param:"user_id"`
}

type DeleteUsers struct {
	UserID string `param:"user_id"`
}
