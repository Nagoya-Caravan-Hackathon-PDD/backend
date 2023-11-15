package output

type CreateUserResponse struct {
	// ユーザの情報
	Message string `json:"message"`
}

type ReadUserResponse struct {
	// ぎっともんの情報

	// ユーザの情報
	UserID   string `json:"user_id"`
	GitHubID string `json:"github_id"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}
