package ports

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
)

type InGameInput interface {
	ReadyGame(reqBody input.ReadyGameRequest) (int, *output.ReadyGameResponse)
	ActionGame(reqBody input.ActionGameRequest) (int, *output.ActionGameResponse)
}

type InGameOutput interface {
	ReadyGame(err error) (int, *output.ReadyGameResponse)
	ActionGame(skill *types.Skill, err error) (int, *output.ActionGameResponse)
}
