package dai

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
)

type InGameDai interface {
	CreateGame(gameID, ownerID string, speed int) error
	GetGameInfo(gameid string) (*types.GameTemp, error)
	ReadyGame(arg input.ReadyGameRequest) error
	IsReady(gameid string) (bool, error)
	UpdateBlock(gameID string, status bool) (int, error)
	IsBlock(gameid string) (bool, error)

	// 本当は分離したほうがいい
	JoinGame(gameID, userID string, speed int) error
	GetSkill(userID string, skillID int) (*types.Skill, error)
	ReasetReady(gameID string) error
	EndGame(gameID string) error
	IsEnd(gameID string) (bool, error)
}
