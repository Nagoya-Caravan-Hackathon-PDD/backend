package dai

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
)

type InGameDai interface {
	GetGameInfo(gameid string) (*types.GameTemp, error)
	ReadyGame(arg input.ReadyGameRequest) error
	IsReady(gameid string) (bool, error)
	UpdateBlock(gameID string, status bool) error
	IsBlock(gameid string) (bool, error)
}
