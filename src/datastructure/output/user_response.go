package output

type CreateUserResponse struct {
	// ぎっともんの情報

	// ユーザの情報
	GitHubID string `json:"github_id"`
}

type ReadUserResponse struct {
	// ぎっともんの情報

	// ユーザの情報
	GitHubID string `json:"github_id"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}
