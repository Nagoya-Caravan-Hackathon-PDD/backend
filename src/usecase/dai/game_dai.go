package dai

type FirestoreDai interface {
	CreateGame(string) error
	JoinGame(string, string) error
	UpdateGame() error
}

type GitmonDai interface {
	GetGitmon() error
}
