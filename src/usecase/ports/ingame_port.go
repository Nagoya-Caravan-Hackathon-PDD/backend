package ports

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
)

type InGameInput interface {
	ReadyGame(reqBody input.ReadyGameRequest) (int, *output.ReadyGameResponse)
	ActionGame(reqBody input.ActionGameRequest) (int, *output.ActionGameResponse)
}

type InGameOutput interface {
	ReadyGame(*output.ReadyGameResponse) error
	ActionGame(*output.ActionGameResponse) error
}
