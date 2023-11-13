package types

type CreateUser struct {
	UserID   string `json:"user_id"`
	GitHubID string `json:"github"`
}
