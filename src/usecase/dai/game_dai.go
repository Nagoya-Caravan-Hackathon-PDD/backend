package dai

import "github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"

type FirestoreDai interface {
	CreateGame(types.CreateGame) error
	JoinGame(types.JoinGame) error
}

type GitmonDai interface {
	GetGitmonStatus(ownerID string) (types.GitmonStatus, error)
}
