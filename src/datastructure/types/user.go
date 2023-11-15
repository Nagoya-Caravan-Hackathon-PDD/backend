package types

type User struct {
	UserID   string `json:"user_id"`
	GitHubID string `json:"github"`
}

type CreateUser struct {
	UserID   string `json:"user_id"`
	GitHubID string `json:"github"`
}
