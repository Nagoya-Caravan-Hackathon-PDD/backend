package dai

import "github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"

type FirestoreDai interface {
	CreateGame(types.CreateGame) error
	JoinGame(types.JoinGame) error
	AddActionLog(index, gameID, userID string, skill types.Skill) error
	GetActionLog(index, gameID, userID string) (*types.Skill, error)
	GetGitmonStatus(gameID, userID string) (*types.GitmonStatus, error)
}

type GitmonDai interface {
	GetGitmonStatus(ownerID string) (types.GitmonStatus, error)
}
